package domain

import "errors"

var (
	ErrDeviceNotFound     = errors.New("device not found")
	ErrInvalidCertificate = errors.New("invalid certificate")
)
