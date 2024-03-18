package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type POW interface {
	generateChallenge() string
	validateResponse(challenge, response string) bool
}

type Payload interface {
	next() string
}

func handleConnection(payload Payload, pow POW, conn net.Conn) {
	defer conn.Close()

	// Generate challenge
	challenge := pow.generateChallenge()

	// Send challenge to client
	_, err := conn.Write([]byte(challenge + "\n"))
	if err != nil {
		fmt.Println("Error writing challenge:", err)
		return
	}

	// Read response from client
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	response = strings.TrimSpace(response)

	// Validate response
	if pow.validateResponse(challenge, response) {

		// Handle client requests...
		n := payload.next()
		fmt.Println("Response verified. Client authorized:", conn.RemoteAddr(), n)

		_, err := conn.Write([]byte(n + "\n"))
		if err != nil {
			fmt.Println("Error writing quote:", err)
			return
		}
	} else {
		fmt.Println("Response verification failed. Closing connection:", conn.RemoteAddr())
	}
}
