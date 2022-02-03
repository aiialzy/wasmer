package binary

import "errors"

var (
	errUnexpectedEnd         = errors.New("unexpected end of section or function")
	errIntTooLong            = errors.New("integer representation too long")
	errIntTooLarge           = errors.New("integer too large")
	errMalformedUTF8Encoding = errors.New("malformed UTF-8 encoding")
	//errLenOutOfBounds = errors.New("length out of bounds")
)
