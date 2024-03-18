package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type hashCash struct {
	difficulty int // Number of leading zeros required for PoW
}

func NewHashCash(difficulty int) POW {
	return &hashCash{
		difficulty: difficulty,
	}
}

func (h *hashCash) generateChallenge() string {
	// Generate random challenge
	challenge := strconv.FormatInt(time.Now().UnixNano(), 10)

	i := 0
	// Calculate Proof of Work
	for {
		hash := sha256.Sum256([]byte(challenge))
		if strings.HasPrefix(fmt.Sprintf("%x", hash), strings.Repeat("0", h.difficulty)) {
			fmt.Printf("%d__%s %x\n", i, challenge, hash)
			break
		} else {
			i++
		}
		challenge = strconv.FormatInt(time.Now().UnixNano(), 10)
	}

	return challenge
}

func (h *hashCash) validateResponse(challenge, response string) bool {
	hash := sha256.Sum256([]byte(challenge + response))
	return strings.HasPrefix(fmt.Sprintf("%x", hash), strings.Repeat("0", h.difficulty))
}
