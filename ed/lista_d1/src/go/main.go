package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	info int
	prev *Node
	next *Node
}

type List struct {
	node *Node
}

func (l *List) NewLList() *List{
	n := &Node{info:0}
	n.next = n
	n.prev = n
	return &List{node: n}
}

func Insert(B *Node) {
	A := B.prev
	C := B.next
	A.next = B
	C.prev = B
}

func Remove(B *Node) {
	A := B.prev
	C := B.next
	A.next = C
	C.prev = A
	B.prev = nil
	B.next = nil 
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			// fmt.Println(ll.Size())
		case "push_back":
			// for _, v := range args[1:] {
			// 	num, _ := strconv.Atoi(v)
			// 	ll.PushBack(num)
			// }
		case "push_front":
			// for _, v := range args[1:] {
			// 	num, _ := strconv.Atoi(v)
			// 	ll.PushFront(num)
			// }
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			// ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
