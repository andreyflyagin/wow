package main

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

type quotes struct {
	words []string
	rand  *rand.Rand
}

func (p *quotes) next() string {
	return p.words[rand.Intn(len(p.words))]
}

func NewQuotes(filename string) (Payload, error) {
	source := rand.NewSource(time.Now().UnixNano())

	words, err := loadTextFileIntoSlice(filename)
	if err != nil {
		return nil, err
	}

	return &quotes{
		words: words,
		rand:  rand.New(source),
	}, nil
}

func loadTextFileIntoSlice(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
