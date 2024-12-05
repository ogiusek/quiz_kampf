package ping

import (
	"io"
	"lib/common/httpapi"
	"lib/common/wraps"

	"net/http"
	"strconv"
	"sync"

	"github.com/ogiusek/hw/src/hw"
)

func fetchData() string {
	// time.Sleep(1 * time.Second)
	return "apecadło chuj i imadłó"
}

var (
	i  int = 0
	mu sync.Mutex
)

func endpointPing(w http.ResponseWriter, r *http.Request) {
	data := fetchData()
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, data)
	io.WriteString(w, "\n")
	mu.Lock()
	i++
	i_as_str := strconv.Itoa(i)
	mu.Unlock()
	io.WriteString(w, i_as_str)
}

type PingCommand struct{}

func ping(PingCommand) any {
	mu.Lock()
	i++
	res := fetchData() + "\n" + strconv.Itoa(i)
	mu.Unlock()
	return res
}

func UsePing(api httpapi.IApi) {
	api.Map("v1/ping", http.MethodGet, endpointPing)
	api.Map("v1/pingg", http.MethodGet, wraps.Wrap(hw.Run(ping)))
}
