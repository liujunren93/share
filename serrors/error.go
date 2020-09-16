package serrors


type Error struct {
	code int32
	msg  string
	Data interface{}
}

func (e Error) Code() int32 {
	return e.code
}

func (e Error) Error() string {
	return e.msg
}

func New(Code int32, msg string, data interface{}) *Error {
	return &Error{
		code: Code,
		msg:  msg,
		Data: data,
	}

}

func BadRequest(err error, data interface{}) *Error {
	return &Error{
		code: 400,
		msg:  err.Error(),
		Data: data,
	}
}

// Unauthorized generates a 401 Error.
func Unauthorized(err error, data interface{}) *Error {
	return &Error{
		code: 401,
		msg:  err.Error(),
		Data: data,
	}
}

// Forbidden generates a 403 Error.
func Forbidden(err error, data interface{}) *Error {
	return &Error{
		code: 403,
		msg:  err.Error(),
		Data: data,
	}
}

// NotFound generates a 404 Error.
func NotFound(err error, data interface{}) *Error {
	return &Error{
		code: 404,
		msg:  err.Error(),
		Data: data,
	}
}

// MethodNotAllowed generates a 405 Error.
func MethodNotAllowed(err error, data interface{}) *Error {
	return &Error{
		code: 405,
		msg:  err.Error(),
		Data: data,
	}
}

// Timeout generates a 408 Error.
func Timeout(err error, data interface{}) *Error {
	return &Error{
		code: 408,
		msg:  err.Error(),
		Data: data,
	}
}

// Conflict generates a 409 Error.
func Conflict(err error, data interface{}) *Error {
	return &Error{
		code: 409,
		msg:  err.Error(),
		Data: data,
	}
}

// InternalServerError generates a 500 Error.
func InternalServerError(err error, data interface{}) *Error {
	return &Error{
		code: 500,
		msg:  err.Error(),
		Data: data,
	}
}
