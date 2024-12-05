package errs

func Forbidden(Msg string) CustomError {
	return CustomError{
		Code: 403,
		Msg:  Msg,
	}
}
