package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

const requests = 10

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: client <server-name:port> <difficulty>")
		return
	}

	serverAddress := os.Args[1]

	difficultyStr := os.Args[2]
	difficulty, err := strconv.Atoi(difficultyStr)
	if err != nil {
		fmt.Println("error difficulty number: ", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(requests)
	for i := 0; i < requests; i++ {
		reqNumber := i
		go func() {
			defer wg.Done()

			conn, err := net.Dial("tcp", serverAddress)
			if err != nil {
				fmt.Println("Error connecting to server:", err)
				return
			}
			defer conn.Close()

			// Read challenge from server
			challenge, err := getStringFromConn(conn)
			if err != nil {
				fmt.Println("Error reading challenge:", err)
				return
			}

			// Solve challenge (Proof of Work)
			solution := solveChallenge(challenge, difficulty)

			// Send solution to server
			_, err = conn.Write([]byte(solution + "\n"))
			if err != nil {
				fmt.Println("Error sending solution:", err)
				return
			}

			quote, err := getStringFromConn(conn)
			if err != nil {
				fmt.Println("Error reading quote:", err)
				return
			}

			fmt.Printf("Req: %d Quote from server: %s\n", reqNumber, quote)
		}()
	}
	wg.Wait()
	fmt.Println("exited...")
}

func getStringFromConn(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	challenge, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	challenge = strings.TrimSpace(challenge)
	return challenge, nil
}

func solveChallenge(challenge string, difficulty int) string {
	solution := ""
	for i := 0; ; i++ {
		solution = strconv.Itoa(i)
		hash := sha256.Sum256([]byte(challenge + solution))
		if strings.HasPrefix(fmt.Sprintf("%x", hash), strings.Repeat("0", difficulty)) {
			break
		}
	}
	return solution
}
