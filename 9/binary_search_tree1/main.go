package main

import (
	"bufio"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextStr() string {
	scanner.Scan()
	return scanner.Text()
}

func nextInt() int {
	n, _ := strconv.Atoi(nextStr())
	return n
}

type Node struct {
	key                 int
	parent, left, right *Node
}

var (
	n         int
	root, pre *Node
	T         []Node
)

func insert(k int) {

}

func main() {
	scanner.Split(bufio.ScanWords)

	n = nextInt()
}
