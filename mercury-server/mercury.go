package main

import (
	"flag"
	"log"
	"net"
	"os"
)

var (
	fFile = flag.String("file", "", "File to serve")
	fBind = flag.String("bind", ":1958", "Bind string to listen on")

	contents []byte
)

func main() {
	var err error

	flag.Parse()
	if *fFile == "" {
		log.Fatalf("must specify file to serve")
	}

	if contents, err = os.ReadFile(*fFile); err != nil {
		log.Fatalf("could not read input file: %v", err)
	}

	l, err := net.Listen("tcp", *fBind)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("failed to accept: %v", err)
			continue
		}
		log.Printf("incoming connection from %v", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	c.Write(contents)
	c.Close()
}
