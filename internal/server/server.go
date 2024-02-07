package server

import (
	"fmt"
	"net/http"
)

func NewServer(port string) *http.Server {
	r := NewRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	return server
}
