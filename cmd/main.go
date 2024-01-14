package main

import (
	"flag"
	"fmt"

	"github.com/lei006/gmqtt"
	"github.com/sohaha/zlsgo/zlog"

	"github.com/panjf2000/gnet/v2"
)

func main() {

	var port int
	var multicore bool

	// Example command: go run server.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 1883, "--port 1883")
	flag.BoolVar(&multicore, "multicore", false, "--multicore=true")
	flag.Parse()
	ss := &gmqtt.SimpleServer{
		Network:   "tcp",
		Addr:      fmt.Sprintf(":%d", port),
		Multicore: multicore,
	}
	err := gnet.Run(ss, ss.Network+"://"+ss.Addr, gnet.WithMulticore(multicore))
	zlog.Infof("server exits with error: %v", err)

	/*
		var port int
		var multicore bool

		// Example command: go run echo.go --port 9000 --multicore=true
		flag.IntVar(&port, "port", 1883, "--port 1883")
		flag.BoolVar(&multicore, "multicore", false, "--multicore true")
		flag.Parse()
		echo := &gmqtt.EchoServer{Addr: fmt.Sprintf("tcp://:%d", port), Multicore: multicore}
		log.Fatal(gnet.Run(echo, echo.Addr, gnet.WithMulticore(multicore)))
	*/
}
