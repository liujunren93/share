package serrors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func New(code int32, msg string) error {
	return status.New(codes.Code(code), msg).Err()
}

func BadRequest(err error) error {
	msg := http.StatusText(400)
	if err == nil {
		msg = err.Error()
	}
	return status.New(codes.Code(400), msg).Err()

}

// Unauthorized generates a 401 Error.
func Unauthorized(err error) error {
	msg := http.StatusText(401)
	if err == nil {
		msg = err.Error()
	}
	return status.New(codes.Code(401), msg).Err()
}

// Forbidden generates a 403 Error.
func Forbidden(err error) error {
	msg := http.StatusText(403)
	if err == nil {
		msg = err.Error()
	}
	return status.New(codes.Code(403),msg).Err()
}

// NotFound generates a 404 Error.
func NotFound(err error) error {
	msg := http.StatusText(404)
	if err == nil {
		msg = err.Error()
	}
	return status.New(codes.Code(404),msg).Err()
}

// MethodNotAllowed generates a 405 Error.
func MethodNotAllowed(err error) error {
	msg := http.StatusText(405)
	if err == nil {
		msg = err.Error()
	}
	return status.New(codes.Code(405), msg).Err()
}

// Timeout generates a 408 Error.
func Timeout(err error) error {
	msg := http.StatusText(408)
	if err == nil {
		msg = err.Error()
	}
	return status.New(codes.Code(408), msg).Err()
}

// Conflict generates a 409 Error.
func Conflict(err error) error {
	msg := http.StatusText(405)
	if err == nil {
		msg = err.Error()
	}
	return status.New(codes.Code(409), msg).Err()
}

// InternalServerError generates a 500 Error.
func InternalServerError(err error) error {
	msg := http.StatusText(500)
	if err == nil {
		msg = err.Error()
	}
	return status.New(codes.Code(500), msg).Err()
}
