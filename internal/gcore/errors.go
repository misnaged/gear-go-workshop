package gcore

import "errors"

var (
	ErrBufferTooSmall             = errors.New("buffer is too small")
	ErrReadFailed                 = errors.New("gr_read failed")
	ErrReplyTo                    = errors.New("gr_reply_to failed")
	ErrSendFailed                 = errors.New("gr_send failed")
	ErrReplyDepositFailed         = errors.New("gr_reply_deposit failed")
	ErrWakeFailed                 = errors.New("gr_wake failed")
	ErrReserveGasFailed           = errors.New("gr_reserve_gas failed")
	ErrUnreserveGasFailed         = errors.New("gr_unreserve_gas failed")
	ErrSystemReserveGasFailed     = errors.New("gr_system_reserve_gas failed")
	ErrReplyCodeFailed            = errors.New("gr_reply_code failed")
	ErrSignalCodeFailed           = errors.New("gr_signal_code failed")
	ErrUnsupportedSignalCode      = errors.New("unsupported signal code")
	ErrReplyFromReservationFailed = errors.New("gr_reservation_reply failed")
)
