package main

import (

	// "github.com/lsh-0/ppp-go/internal/components/lax-proxy/interface"
	// "github.com/lsh-0/ppp-go/internal/http"
	"github.com/lsh-0/ppp-go/internal/api"
	"github.com/lsh-0/ppp-go/internal/log"
	// "github.com/lsh-0/ppp-go/internal/types"
	// "github.com/lsh-0/ppp-go/internal/utils"
	// "fmt"
)

func main() {
	log.Debug("started")
	opts := api.RequestConfig{}
	log.Info("got: ", api.Request("/articles", opts))
	log.Debug("done")
}
