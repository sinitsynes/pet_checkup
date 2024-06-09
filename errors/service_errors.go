package service_errors

import "errors"

var ErrDecoder = errors.New("decoder error")
var ErrDatabaseConnection = errors.New("DB connection error")
var ErrMarshalError = errors.New("marshalling error")
