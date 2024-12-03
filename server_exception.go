package server

import (
	"net/http"
)

const (
	// AuthExceptionCode auth error code
	AuthExceptionCode = 10000

	// RequestExceptionCode request error code
	RequestExceptionCode = 10001

	// InternalExceptionCode internal error code
	// like store error
	InternalExceptionCode = 10002

	// NotFoundExceptionCode request path not found
	NotFoundExceptionCode = 10003

	// UnknownExceptionCode unknown error
	UnknownExceptionCode = 10100

	// DataExistedExceptionCode data is existed
	DataExistedExceptionCode = 20000

	// DataNotFoundExceptionCode data not found
	DataNotFoundExceptionCode = 20001

	// DataExpiredExceptionCode data not found
	DataExpiredExceptionCode = 20002

	// DataUnsupportedExceptionCode data unsupported found
	DataUnsupportedExceptionCode = 20003

	// UserNamePasswordNotMatchExceptionCode user's username and password is not match
	UserNamePasswordNotMatchExceptionCode = 30000

	// UserDeactivatedExceptionCode user is deactivated
	UserDeactivatedExceptionCode = 30001

	// UserVerifiedCodeErrorExceptionCode user is deactivated
	UserVerifiedCodeErrorExceptionCode = 30002
)

func AuthException() *Exception {
	return newException(
		AuthExceptionCode,
		http.StatusText(http.StatusUnauthorized),
	)
}

func NotFoundException() *Exception {
	return newException(
		NotFoundExceptionCode,
		http.StatusText(http.StatusNotFound),
	)
}

func NotFoundExceptionWithMsg(msg string) *Exception {
	return newException(
		NotFoundExceptionCode,
		msg,
	)
}

func RequestException(msg string) *Exception {
	return newException(
		RequestExceptionCode,
		msg,
	)
}

func UnknownException(msg string) *Exception {
	return newException(
		UnknownExceptionCode,
		msg,
	)
}

func InternalException() *Exception {
	return newException(
		InternalExceptionCode,
		http.StatusText(http.StatusInternalServerError),
	)
}

func DataExistedException() *Exception {
	return newException(
		DataExistedExceptionCode,
		"data is existed",
	)
}

func DataNotFoundException(msg string) *Exception {
	return newException(
		DataNotFoundExceptionCode,
		msg,
	)
}

func DataExpiredException() *Exception {
	return newException(
		DataExpiredExceptionCode,
		"data is expired",
	)
}

func UserDeactivatedException() *Exception {
	return newException(
		UserDeactivatedExceptionCode,
		"user is deactivated",
	)
}

func UserNamePasswordNotMatchException() *Exception {
	return newException(
		UserNamePasswordNotMatchExceptionCode,
		"username and password is not match",
	)
}

func UserVerifiedCodeErrorException() *Exception {
	return newException(
		UserVerifiedCodeErrorExceptionCode,
		"verified code is not correct",
	)
}
