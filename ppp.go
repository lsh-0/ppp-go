package main

import (

	// "github.com/lsh-0/ppp-go/internal/components/lax-proxy/interface"
	// "github.com/lsh-0/ppp-go/internal/http"
	"github.com/lsh-0/ppp-go/internal/log"
	// "github.com/lsh-0/ppp-go/internal/types"
	// "github.com/lsh-0/ppp-go/internal/utils"
	"fmt"
)

func main() {
	fmt.Println("started...")
	log.Debug("debug?")
	log.Info("info.")
	log.Warn("*warn*")
	log.Error("error!")
	fmt.Println("...done.")
}
