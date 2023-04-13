package util

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func RunPprof(enable bool, port int) {
	if !enable {
		return
	}
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()
}
