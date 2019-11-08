package main

import (
	"bufio"
	"os"
	"strconv"
)

var n int

var scanner = bufio.NewScanner(os.Stdin)

func nextStr() string {
	scanner.Scan()
	return scanner.Text()
}

func nextInt() int {
	n, _ := strconv.Atoi(nextStr())
	return n
}

func main() {
	scanner.Split(bufio.ScanWords)

	n = nextInt()
}
