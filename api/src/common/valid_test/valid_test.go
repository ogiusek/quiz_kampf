package valid_test

import (
	valid1 "lib/common/valid"
	"testing"
)

func Test_valid_should_return_error(t *testing.T) {
	model := TestModel{
		Validable: invalid,
		Val:       "",
	}
	expected := model.Validable.Valid()

	err := valid1.Valid(model)

	if err == nil {
		t.Errorf("error: nil; expected: %v", expected)
		return
	}

	if err.Error() != expected.Error() {
		t.Errorf("error: %v; expected: %v", err, expected)
		return
	}
}

func Test_valid_should_return_nil(t *testing.T) {
	model := TestModel{
		Validable: valid,
		Val:       "",
	}

	err := valid1.Valid(model)

	if err != nil {
		t.Errorf("error: %v; expected: nil", err)
		return
	}
}
