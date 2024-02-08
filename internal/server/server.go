package server

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewServer(port string) *http.Server {
	r := NewRouter()

    h2s := &http2.Server{}
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		Handler: h2c.NewHandler(r, h2s),
	}

	return server
}
