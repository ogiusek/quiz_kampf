package errs

func InvalidInput(Msg string) CustomError {
	return CustomError{
		Code: 400,
		Msg:  Msg,
	}
}
