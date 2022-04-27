package gosign

import "errors"

var (
	ErrSignFailed = errors.New("sign failed")
	ErrFileName   = errors.New("error file name")
	ErrSigned     = errors.New("file has been signed")
)
