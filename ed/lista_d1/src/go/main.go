package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	Value int
	prev *Node
	next *Node
}

type LList struct {
	root *Node
}

func NewLList() *LList {
	root := &Node{Value:0}
	root.prev = root
	root.next = root
	return &LList{root: root}
}

func (ll *LList) Insert(node *Node) {
	A := node.prev
	C := node.next
	A.next = node
	C.prev = node
}

func (ll *LList) Remove(node *Node) {
	A := node.prev
	C := node.next
	A.next = C
	C.prev = A
	node.prev = nil
	node.next = nil
}

func (ll *LList) PushFront(value int) {
	node := &Node{Value: value}
	node.prev = ll.root
	node.next = ll.root.next
	ll.Insert(node)
}

func (ll *LList) PushBack(value int) {
	node := &Node{Value: value}
	node.prev = ll.root.prev
	node.next = ll.root
	ll.Insert(node)
}

func (ll *LList) PopFront() {
	if ll.root.next != ll.root {
		ll.Remove(ll.root.next)
	}
}

func (ll *LList) PopBack() {
	if ll.root.prev != ll.root {
		ll.Remove(ll.root.prev)
	}
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
}

func (ll *LList) Size() int {
	contador := 0

	atual := ll.root.next

	for atual != ll.root {
		contador++
		atual = atual.next
	}

	return contador
}

func (ll *LList) String() string {
	if ll.Size() == 0 {
		return "[]"
	}

	saida := "["
	atual := ll.root.next

	for atual != ll.root {
		saida += fmt.Sprintf("%d", atual.Value)

		if atual.next != ll.root {
			saida += ", "
		}
		atual = atual.next
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
