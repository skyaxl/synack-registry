package apierrors

import (
	"fmt"
	"net/http"
)

//ApiError
type ApiError string

var (
	//ErrNotFound err not found
	ErrNotFound = ApiError("not_found")
	//ErrNotAuthorized api not authorized
	ErrNotAuthorized = ApiError("not_authorized")
	//ErrInternalServerError api not authorized
	ErrInternalServerError = ApiError("internal_server")
	//ErrBadRequest api not authorized
	ErrBadRequest = ApiError("internal_server")

	codeMap = map[ApiError]int{
		ErrNotFound:            http.StatusNotFound,
		ErrNotAuthorized:       http.StatusForbidden,
		ErrInternalServerError: http.StatusInternalServerError,
		ErrBadRequest:          http.StatusBadRequest,
	}
)

func (ae ApiError) Error() string {
	return string(ae)
}

//HTTPCode Get http response code
func (ae ApiError) HTTPCode() int {
	return codeMap[ae]
}

//MarshalJSON custom serializer
func (ae ApiError) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf(`{"message": "%s","status": %d }`, string(ae), ae.HTTPCode())
	return []byte(str), nil
}
