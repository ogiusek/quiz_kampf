package errs

func Exists(Msg string) CustomError {
	return CustomError{
		Code: 409,
		Msg:  Msg,
	}
}
