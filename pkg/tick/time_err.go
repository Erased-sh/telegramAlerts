package tick

import (
	"errors"
)

var (
	errInvalidTag      = errors.New("invalid time_tag provided")
	errMalformedAction = errors.New("function run out with error")
)
