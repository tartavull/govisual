package main

import (
	"flag"
	"fmt"
	"github.com/tartavull/govisual/server"
)

const message = `Graphs the dependencies  between Go packages.

Usage

     $GOPATH/bin/govisual 

The program presents a web interface on port 8080(default).
The various visualisations are registred as top level handlers. Here are some examples

    http://localhost:8080/tree/io
    http://localhost:8080/radial/math/rand
    http://localhost:8080/forcegraph/fmt
    http://localhost:8080/chord/cmd/go

govisual works sort of like godoc.

  flags:
  
`

func main() {

	port := flag.Int("port", 8080, "Port number")

	run := true
	flag.Usage = func() {
		fmt.Printf(message)
		flag.PrintDefaults()
		run = false
		return
	}

	flag.Parse()

	if run {
		server.Serve(*port)
	}
}
