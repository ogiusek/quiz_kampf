package usecases

import "errors"

var ErrAlreadyInMatch error = errors.New("you're currently in other match you have to leave it")
