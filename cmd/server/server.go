package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: server <filename>")
		return
	}

	filename := os.Args[1]

	difficultyStr := os.Args[2]
	difficulty, err := strconv.Atoi(difficultyStr)
	if err != nil {
		fmt.Println("error difficulty number: ", err)
		return
	}

	fmt.Println("Starting TCP Server...")

	// Start TCP server
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	var algo POW
	algo = NewHashCash(difficulty)

	var payload Payload
	payload, err = NewQuotes(filename)
	if err != nil {
		fmt.Println("Error creating payload:", err)
		return
	}

	// Listen for connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Client connected:", conn.RemoteAddr())
		go handleConnection(payload, algo, conn)
	}
}
