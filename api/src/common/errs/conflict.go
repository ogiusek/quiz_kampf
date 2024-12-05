package errs

func Conflict(Msg string) CustomError {
	return CustomError{
		Code: 409,
		Msg:  Msg,
	}
}
