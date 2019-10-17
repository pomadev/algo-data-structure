package main

import (
	"bufio"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextWord() string {
	scanner.Scan()
	return scanner.Text()
}

func nextInt() int {
	s := nextWord()
	n, _ := strconv.Atoi(s)
	return n
}

type Order struct {
	order   string
	content string
}

func main() {
	scanner.Split(bufio.ScanWords)

	n := nextInt()

	s := make([]Order, n)
	for i := range s {
		s[i].order = nextWord()
		s[i].content = nextWord()
	}
}
