package main

import (
	"flag"
	"fmt"

	"github.com/haonancool/simple-tools/server"
)

var (
	port = flag.Int("p", 0, "port")
	help = flag.Bool("h", false, "help")
)

func main() {
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		return
	}

	if *port <= 0 {
		panic(fmt.Errorf("port[%d] <= 0", *port))
	}

	addr := fmt.Sprintf(":%d", *port)
	s := server.NewServer(addr)
	if err := s.Start(); err != nil {
		panic(err)
	}
}
