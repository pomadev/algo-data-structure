package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	key        int
	next, prev *Node
}

var guard *Node

func init() {
	guard = new(Node)
	guard.next = guard
	guard.prev = guard
}

func listSearch(key int) *Node {
	cur := guard.next
	for cur != guard && cur.key != key {
		cur = cur.next
	}
	return cur
}

func printList() {
	cur := guard.next
	isf := 0
	for {
		if cur == guard {
			break
		}
		if isf > 0 {
			fmt.Print(" ")
		}
		isf++
		fmt.Print(cur.key)
		cur = cur.next
	}
	fmt.Println()
}

func deleteNode(t *Node) {
	// 番兵の場合は処理しない
	if t == guard {
		return
	}
	t.prev.next = t.next
	t.next.prev = t.prev
}

func deleteFirst() {
	deleteNode(guard.next)
}

func deleteLast() {
	deleteNode(guard.prev)
}

func deleteKey(key int) {
	deleteNode(listSearch(key))
}

func insert(key int) {
	x := new(Node)
	x.key = key
	// 番兵の直後に要素を追加
	x.next = guard.next
	guard.next.prev = x
	guard.next = x
	x.prev = guard
}

var scanner = bufio.NewScanner(os.Stdin)

func nextLine() string {
	scanner.Scan()
	return scanner.Text()
}

func nextInt() int {
	n, _ := strconv.Atoi(nextLine())
	return n
}

func main() {
	// TODO 1つテストケース時間切れなので短縮方法考える
	scanner.Split(bufio.ScanWords)

	n := nextInt()
	for i := 0; i < n; i++ {
		c := nextLine()
		if c[0] == 'i' {
			insert(nextInt())
			continue
		}
		if len(c) > 6 {
			if c[6] == 'F' {
				deleteFirst()
				continue
			}
			deleteLast()
			continue
		}
		deleteKey(nextInt())
	}
	printList()
}
