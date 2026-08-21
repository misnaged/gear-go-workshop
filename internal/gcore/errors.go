package gcore

import "errors"

var (
	ErrBufferTooSmall     = errors.New("buffer is too small")
	ErrReadFailed         = errors.New("gr_read failed")
	ErrReplyTo            = errors.New("gr_reply_to failed")
	ErrSendFailed         = errors.New("gr_send failed")
	ErrReplyDepositFailed = errors.New("gr_reply_deposit failed")
	ErrWakeFailed         = errors.New("gr_wake failed")
)
