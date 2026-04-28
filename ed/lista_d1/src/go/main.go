package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	info int
	prev *Node
	next *Node
}

type List struct {
	node *Node
}

func NewLList() *List{
	n := &Node{info:0}
	n.next = n
	n.prev = n
	return &List{node: n}
}

func (l *List) Insert(B *Node) {
	A := B.prev
	C := B.next
	A.next = B
	C.prev = B
}

func (l *List) Remove(B *Node) {
	A := B.prev
	C := B.next
	A.next = C
	C.prev = A
	B.prev = nil
	B.next = nil
}

func (l *List) PushFront(value int) {
	node := &Node{info: value}

	node.prev = l.node
	node.next = l.node.next

	l.Insert(node)
}

func (l *List) PushBack(value int) {
	node := &Node{info: value}

	node.prev = l.node.prev
	node.next = l.node

	l.Insert(node)
}

func (l *List) PopBack() {
	if l.Size() > 0 {
		l.Remove(l.node.prev)
	}
}

func (l *List) PopFront() {
	if l.Size() > 0 {
		l.Remove(l.node.next)
	}
}

func (l *List) Size() int {
	contaElementos := 0

	nodeAtual := l.node.next

	for nodeAtual != l.node {
		contaElementos++
		nodeAtual = nodeAtual.next
	}

	return contaElementos
}

func (l *List) Clear() {
	l.node.next = l.node
	l.node.prev = l.node
}

func (l *List) String() string {
	if l.Size() == 0 {
		return "[]"
	}

	saida := "["

	it := l.node.next

	for it != l.node {
		saida += fmt.Sprintf("%d", it.info)

		if it.next != l.node {
			saida += ", "
		}
		it = it.next
	}

	return saida + "]"
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
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
			num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
			num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
