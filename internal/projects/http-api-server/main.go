package main

import (
	"net/http"

	"github.com/lsh-0/ppp-go/internal/api"
)

func main() {
	http.HandleFunc("/", api.ProxyHttpRequest)
	http.ListenAndServe(":8090", nil)
}
