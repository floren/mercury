package main

import (
	"os"
	"io"
	"net"
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Printf("Mercury host (q to quit)> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("failed to read input: %v", err)
		}
		input = strings.TrimSpace(input)
		if input == "q" || input == "quit" {
			return
		}
		conn, err := net.Dial("tcp", input)
		if err != nil {
			log.Printf("failed to connect to mercury server: %v", err)
			continue
		}
		io.Copy(os.Stdout, conn)
		conn.Close()
	}
}
