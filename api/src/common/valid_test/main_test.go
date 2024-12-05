package valid_test

import "errors"

type TestValidable string

func (v TestValidable) Value() any {
	return string(v)
}

func (v TestValidable) Valid() error {
	if v.Value() == invalid.Value() {
		return invalidError
	}
	return nil
}

type TestModel struct {
	Validable TestValidable
	Val       string
}

var invalid TestValidable = TestValidable("")
var valid TestValidable = TestValidable("v")

//lint:ignore ST1012 This is an intentional naming choice
var invalidError error = errors.New("value cannot be empty")
