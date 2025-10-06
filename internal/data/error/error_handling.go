package repository_error_handling

import "errors"

var DbNotFound = errors.New("record not found")
var DbInternalError = errors.New("db internal error")
