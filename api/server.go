package api

import "net/http"

func NewServer(port string) (*http.Server, error) {
	return &http.Server{
		Addr:    ":" + port,
		Handler: newRouter(),
	}, nil
}
