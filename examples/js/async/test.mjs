import { readFileSync } from 'node:fs';

import {
    GearApi,
    generateCodeHash,
} from '@gear-js/api';

import { Keyring } from '@polkadot/keyring';
import { waitReady } from '@polkadot/wasm-crypto';


const NODE = process.env.GEAR_NODE ?? 'ws://127.0.0.1:9944';
const ACCOUNT_URI = process.env.GEAR_ACCOUNT ?? '//Alice';

const [echoPath, asyncPath] = process.argv.slice(2);

if (!echoPath || !asyncPath) {
    console.error(
        'Usage: node test.mjs /path/to/echo.wasm /path/to/async.wasm'
    );
    process.exit(1);
}


function timeout(promise, ms, message) {
    return Promise.race([
        promise,

        new Promise((_, reject) => {
            setTimeout(
                () => reject(new Error(message)),
                ms,
            );
        }),
    ]);
}


async function getAccount(uri) {
    await waitReady();

    const keyring = new Keyring({
        type: 'sr25519',
        ss58Format: 137,
    });

    return keyring.addFromUri(uri);
}


async function submitTransaction(
    api,
    tx,
    account,
    wantedEvent = null,
) {
    return new Promise((resolve, reject) => {
        let unsubPromise;

        unsubPromise = tx.signAndSend(
            account,
            ({ events, status }) => {
                const failed = events.find(
                    ({ event }) =>
                        event.method === 'ExtrinsicFailed'
                );

                if (failed) {
                    Promise.resolve(unsubPromise)
                        .then((unsub) => unsub())
                        .catch(() => {});

                    try {
                        reject(
                            api.getExtrinsicFailedError(
                                failed.event,
                            ),
                        );
                    } catch {
                        reject(
                            new Error(
                                `ExtrinsicFailed: ${failed.event.data.toString()}`
                            ),
                        );
                    }

                    return;
                }

                if (!status.isInBlock) {
                    return;
                }

                let foundEvent = null;

                if (wantedEvent !== null) {
                    const record = events.find(
                        ({ event }) =>
                            event.method === wantedEvent
                    );

                    if (record) {
                        foundEvent = record.event;
                    }
                }

                const blockHash =
                    status.asInBlock.toHex();

                Promise.resolve(unsubPromise)
                    .then((unsub) => unsub())
                    .catch(() => {});

                resolve({
                    event: foundEvent,
                    blockHash,
                });
            },
        );

        Promise.resolve(unsubPromise).catch(reject);
    });
}


function waitProgramActive(api, programId) {
    let unsubPromise;

    const promise = new Promise((resolve, reject) => {
        unsubPromise =
            api.gearEvents.subscribeToGearEvent(
                'ProgramChanged',
                ({ data }) => {
                    if (!data.id.eq(programId)) {
                        return;
                    }

                    if (data.change.isProgramSet) {
                        console.log(
                            `    ProgramSet: ${programId}`
                        );

                        return;
                    }

                    if (data.change.isActive) {
                        console.log(
                            `    Active: ${programId}`
                        );

                        Promise.resolve(unsubPromise)
                            .then((unsub) => unsub())
                            .catch(() => {});

                        resolve();
                        return;
                    }

                    if (data.change.isTerminated) {
                        Promise.resolve(unsubPromise)
                            .then((unsub) => unsub())
                            .catch(() => {});

                        reject(
                            new Error(
                                `Program terminated: ${programId}`
                            ),
                        );
                    }
                },
            );

        Promise.resolve(unsubPromise).catch(reject);
    });

    return timeout(
        promise,
        60_000,
        `Timeout waiting for program Active: ${programId}`,
    );
}


async function ensureCode(
    api,
    account,
    wasm,
    name,
) {
    const codeId = generateCodeHash(wasm);

    console.log(`${name} CodeId: ${codeId}`);

    if (await api.code.exists(codeId)) {
        console.log(
            `${name}: code already uploaded`
        );

        return codeId;
    }

    console.log(`${name}: uploading code...`);

    const {
        extrinsic,
    } = await api.code.upload(wasm);

    await submitTransaction(
        api,
        extrinsic,
        account,
    );

    console.log(`${name}: code uploaded`);

    return codeId;
}


async function createProgram(
    api,
    account,
    codeId,
    initPayload,
    name,
) {
    console.log(`${name}: creating program...`);

    const {
        programId,
        extrinsic,
    } = api.program.create({
        codeId,
        initPayload,

        gasLimit: api.blockGasLimit,

        value: 0,
        keepAlive: true,
    });

    console.log(
        `${name} ProgramId: ${programId}`
    );


    const activePromise =
        waitProgramActive(api, programId);

    await submitTransaction(
        api,
        extrinsic,
        account,
    );

    await activePromise;

    console.log(`${name}: ready`);

    return programId;
}


async function sendPing(
    api,
    account,
    asyncProgramId,
) {
    console.log('');
    console.log('Sending PING...');


    const payload = Buffer.from('PING', 'utf8');

    const tx = api.message.send({
        destination: asyncProgramId,
        payload,

        gasLimit: 5_000_000_000,

        value: 0,
        keepAlive: true,
    });

    const {
        event,
        blockHash,
    } = await submitTransaction(
        api,
        tx,
        account,
        'MessageQueued',
    );

    if (!event) {
        throw new Error(
            'MessageQueued event not found'
        );
    }

    const messageId =
        event.data.id?.toHex?.()
        ?? event.data[0].toHex();

    console.log(
        `PING MessageId: ${messageId}`
    );

    console.log('Waiting for reply...');

    const reply = await timeout(
        api.message.getReplyEvent(
            asyncProgramId,
            messageId,
            blockHash,
        ),
        60_000,
        'Timeout waiting for PONG',
    );

    const replyHex =
        reply.data.message.payload.toHex();

    const replyText =
        Buffer.from(
            replyHex.slice(2),
            'hex',
        ).toString('utf8');

    console.log('');
    console.log('==========================');
    console.log(`Reply hex:  ${replyHex}`);
    console.log(`Reply text: ${replyText}`);
    console.log('==========================');

    if (replyText !== 'PONG') {
        throw new Error(
            `Expected PONG, received: ${replyText}`
        );
    }

    console.log('');
    console.log('SUCCESS: PING -> PONG');
}


async function main() {
    console.log(`Connecting to ${NODE}...`);

    const api = await GearApi.create({
        providerAddress: NODE,
    });

    try {
        await api.isReadyOrError;

        const account =
            await getAccount(ACCOUNT_URI);

        console.log(
            `Account: ${account.address}`
        );

        const echoWasm =
            new Uint8Array(
                readFileSync(echoPath),
            );

        const asyncWasm =
            new Uint8Array(
                readFileSync(asyncPath),
            );


        console.log('');
        console.log('=== ECHO ===');

        const echoCodeId =
            await ensureCode(
                api,
                account,
                echoWasm,
                'echo',
            );

        const echoProgramId =
            await createProgram(
                api,
                account,
                echoCodeId,

                new Uint8Array(),

                'echo',
            );


        console.log('');
        console.log('=== ASYNC ===');

        const asyncCodeId =
            await ensureCode(
                api,
                account,
                asyncWasm,
                'async',
            );


        const echoProgramIdBytes =
            Buffer.from(
                echoProgramId.slice(2),
                'hex',
            );

        if (echoProgramIdBytes.length !== 32) {
            throw new Error(
                `Invalid echo ProgramId length: ${echoProgramIdBytes.length}`
            );
        }

        const asyncProgramId =
            await createProgram(
                api,
                account,
                asyncCodeId,


                echoProgramIdBytes,

                'async',
            );


        await sendPing(
            api,
            account,
            asyncProgramId,
        );
    } finally {
        await api.disconnect();
    }
}


main().catch((err) => {
    console.error('');
    console.error('FAILED');
    console.error(err);

    process.exitCode = 1;
});