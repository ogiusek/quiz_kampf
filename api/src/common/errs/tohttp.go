package errs

import (
	"net/http"
)

func (err CustomError) ToHttp(w http.ResponseWriter) {
	w.WriteHeader(err.Code)
	w.Write([]byte(err.Msg))
}

func ToHttp(w http.ResponseWriter, err error) {
	if customErr, ok := err.(CustomError); ok {
		customErr.ToHttp(w)
		return
	}
	Error("wrong server response %v", err).ToHttp(w)
}
