BIN:=./build/tinygo
TINYGO_CFG:= ./config/tinygo/gear.json
#OUTPUT:= ./program.wasm
OUTPUT:= ../gear-go/assets/wasm/test/program.wasm

wasm-build:
	$(BIN) build -target $(TINYGO_CFG) -o  $(OUTPUT)   .


tidy:
	go mod tidy


update:
	go get -u ./...