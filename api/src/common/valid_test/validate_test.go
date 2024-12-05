package valid_test

import (
	valid1 "lib/common/valid"
	"testing"
)

func Test_validate_should_return_error(t *testing.T) {
	err := valid1.ValidMany([]valid1.Validable{invalid})

	if err == nil || err.Error() != invalidError.Error() {
		t.Errorf("error: %v; expected: %v", err, invalidError)
	}
}

func Test_validate_should_return_nil(t *testing.T) {
	err := valid1.ValidMany([]valid1.Validable{valid})

	if err != nil {
		t.Errorf("error: %v; expected: nil", err)
	}
}
