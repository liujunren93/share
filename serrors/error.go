package serrors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(code int32, msg string) error {
	return status.New(codes.Code(code), msg).Err()

}

func BadRequest(err error) error {
	if err == nil {
		return nil
	}
	return status.New(codes.Code(400), err.Error()).Err()

}

// Unauthorized generates a 401 Error.
func Unauthorized(err error) error {
	if err == nil {
		return nil
	}
	return status.New(codes.Code(401), err.Error()).Err()
}

// Forbidden generates a 403 Error.
func Forbidden(err error) error {
	if err == nil {
		return nil
	}
	return status.New(codes.Code(403), err.Error()).Err()
}

// NotFound generates a 404 Error.
func NotFound(err error) error {
	if err == nil {
		return nil
	}
	return status.New(codes.Code(404), err.Error()).Err()
}

// MethodNotAllowed generates a 405 Error.
func MethodNotAllowed(err error) error {
	if err == nil {
		return nil
	}
	return status.New(codes.Code(405), err.Error()).Err()
}

// Timeout generates a 408 Error.
func Timeout(err error) error {
	if err == nil {
		return nil
	}
	return status.New(codes.Code(408), err.Error()).Err()
}

// Conflict generates a 409 Error.
func Conflict(err error) error {
	if err == nil {
		return nil
	}
	return status.New(codes.Code(409), err.Error()).Err()
}

// InternalServerError generates a 500 Error.
func InternalServerError(err error) error {
	if err == nil {
		return nil
	}
	return status.New(codes.Code(500), err.Error()).Err()
}
