package controller_error_handling

import "errors"

var ControllerInputInvalid = errors.New("Input is invalid")
var ControllerInternal = errors.New("Internal error")
