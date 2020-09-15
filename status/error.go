package status

type Error struct {
	code int32
	msg  string
	Data interface{}
}

func (e Error) GetCode() int32 {
	return e.code
}

func (e Error) GetMsg() string {
	return e.msg
}

func (e Error) Error() string {
	return e.msg
}

func New(Code int32, msg string, data interface{}) (error, status) {
	err := Error{
		code: Code,
		msg:  msg,
		Data: data,
	}
	return err, err
}

func BadRequest(err error, data interface{}) error {
	return Error{
		code: 400,
		msg:  err.Error(),
		Data: data,
	}
}

// Unauthorized generates a 401 error.
func Unauthorized(err error, data interface{}) error {
	return &Error{
		code: 401,
		msg:  err.Error(),
		Data: data,
	}
}

// Forbidden generates a 403 error.
func Forbidden(err error, data interface{}) error {
	return &Error{
		code: 403,
		msg:  err.Error(),
		Data: data,
	}
}

// NotFound generates a 404 error.
func NotFound(err error, data interface{}) error {
	return &Error{
		code: 404,
		msg:  err.Error(),
		Data: data,
	}
}

// MethodNotAllowed generates a 405 error.
func MethodNotAllowed(err error, data interface{}) error {
	return &Error{
		code: 405,
		msg:  err.Error(),
		Data: data,
	}
}

// Timeout generates a 408 error.
func Timeout(err error, data interface{}) error {
	return &Error{
		code: 408,
		msg:  err.Error(),
		Data: data,
	}
}

// Conflict generates a 409 error.
func Conflict(err error, data interface{}) error {
	return &Error{
		code: 409,
		msg:  err.Error(),
		Data: data,
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(err error, data interface{}) error {
	return &Error{
		code: 500,
		msg:  err.Error(),
		Data: data,
	}
}
