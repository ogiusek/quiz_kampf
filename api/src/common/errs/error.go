package errs

import "log"

func Error(format string, v ...any) CustomError {
	log.Printf(format, v...)
	return CustomError{
		Code: 500,
		Msg:  "Server error",
	}
}
