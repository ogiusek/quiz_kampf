package id_test

import (
	"lib/common/id"
	"testing"
)

func Test_new_id_should_not_be_null(t *testing.T) {
	id := id.New()

	if string(id) == "" {
		t.Errorf("new error id is empty")
	}
}
