package errs

func Unauthorized(Msg string) CustomError {
	return CustomError{
		Code: 401,
		Msg:  Msg,
	}
}
