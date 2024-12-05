package errs

import "fmt"

type CustomError struct {
	Code int
	Msg  string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Msg)
}
