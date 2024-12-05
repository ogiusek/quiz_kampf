package servicepool_test

import (
	"lib/common/servicepool"
	"testing"
	"time"
)

func Test_servicepool_should_return_free_service(t *testing.T) {
	elements := []int{1, 2}
	services := servicepool.PoolFromArray(elements)

	services.Get()
	s2, free2 := services.Get()
	free2()
	s3, _ := services.Get()

	if s2 != s3 {
		t.Errorf("services are not equal %v != %v", s2, s3)
	}
}

func Test_should_get_service_after_service_got_freed(t *testing.T) {
	service := 1
	services := servicepool.PoolFromArray([]int{service})
	executed := false

	s, free := services.Get()

	if s != service {
		t.Errorf("service is not equal to its original value %v != %v", s, service)
		return
	}

	fn := func() {
		time.Sleep(1 * time.Millisecond)
		executed = true
		free()
	}

	go fn()
	s, free = services.Get()
	defer free()
	if !executed {
		t.Errorf("got service when no service was available")
	}

	if s != service {
		t.Errorf("service is not equal to its original value %v != %v", s, service)
	}
}
