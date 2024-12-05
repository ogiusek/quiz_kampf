package errs

func NotFound(Msg string) CustomError {
	return CustomError{
		Code: 404,
		Msg:  Msg,
	}
}
