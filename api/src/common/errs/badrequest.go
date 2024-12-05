package errs

func BadRequest(Msg string) CustomError {
	return CustomError{
		Code: 400,
		Msg:  Msg,
	}
}
