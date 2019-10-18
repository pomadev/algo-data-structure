package main

import (
	"bufio"
	"fmt"
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

func main() {
	scanner.Split(bufio.ScanWords)

	n := nextInt()

	m := make(map[string]bool)
	for i := 0; i < n; i++ {
		order, content := nextWord(), nextWord()
		if order == "insert" {
			m[content] = true
		} else {
			_, ok := m[content]
			if ok {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		}
	}
}

func h1(key, m int) int {
	return key % m
}

func h2(key, m int) int {
	return 1 + (key % (m - 1))
}

func h(key, m, i int) int {
	return (h1(key, m) + i*h2(key, m)) % m
}
