package ping_test

import (
	testconfig "lib/test_config"
	"net/http"
	"strconv"
	"sync"
	"testing"
)

func fetch(url string, t *testing.B, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	_, err := http.Get(url)

	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}
}

func BenchmarkTest(t *testing.B) {
	url := testconfig.Config.Api + "/ping"
	wg := &sync.WaitGroup{}

	requests := 1000

	t.ResetTimer()
	for i := 0; i <= requests; i++ {
		go fetch(url, t, wg)
	}
	wg.Wait()
	elapsed := t.Elapsed()
	sec := strconv.FormatFloat(elapsed.Seconds(), 'f', 2, 64) + "s"
	req_per_sec := strconv.FormatFloat(float64(requests)/elapsed.Seconds(), 'f', 0, 64)
	t.Log("requests per second (" + strconv.Itoa(requests) + "/" + sec + "): " + req_per_sec)
}
