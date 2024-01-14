package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/lei006/gmqtt"
)

func main() {

	var port int
	var multicore bool

	// Example command: go run server.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 1883, "--port 1883")
	flag.BoolVar(&multicore, "multicore", false, "--multicore=true")
	flag.Parse()

	wg := sync.WaitGroup{}
	wg.Add(1)

	cfg := &gmqtt.ServerConfig{
		Addr: fmt.Sprintf(":%d", port),
	}

	gmqtt.ListenAndServer(cfg, &wg)

	wg.Wait()

}
