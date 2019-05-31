package ego

import (
	"net/http"
)

type Ego struct {
}

var App *Ego

func getApp() *Ego {
	if App == nil {
		App = &Ego{}
	}
	return App
}

func Run(port ...string) error {
	addr := "8088"
	if len(port) >= 1 {
		addr = port[0]
	}
	return http.ListenAndServe(":"+addr, nil)
}
