package main

import (
	"flag"
	"fmt"
)

func main() {
	const defaultPort = 3000
	var port int
	flag.IntVar(&port, "port", defaultPort, "port to use")
	flag.IntVar(&port, "p", defaultPort, "port to use (short)")
	flag.Parse()
	fmt.Println(port)
}
