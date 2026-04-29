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
	next *Node
	prev *Node
	root *Node
}

type LList struct {
	root *Node
	size int
}

func (n *Node) Next() *Node {
	if n.root.next == n.root {
		return nil
	}

	return n.next
}

func (n *Node) Prev() *Node {
	if n.root.prev == n.root {
		return nil
	}

	return n.prev
}

func NewLList() *LList {
	n := &Node{Value: 0}

	n.next = n
	n.prev = n
	n.root = n

	return &LList{root: n, size: 0}
}

func (l *LList) Front() *Node {
	if l.root.next == nil || l.root.next == l.root{
		return nil
	}

	return l.root.next
}

func (l *LList) Back() *Node {
	if l.root.prev == nil || l.root.prev == l.root {
		return nil
	}

	return l.root.prev
}

func (l *LList) Search(value int) *Node {
	if l.root.next == l.root || l.root.next == nil {
		return nil
	}

	atual := l.root.next

	for atual != l.root {
		if atual.Value == value {
			return atual
		}
		atual = atual.next
	}

	return nil
}

func (l *LList) Insert(node *Node, value int) {
	novoNode := &Node{Value: value, root: l.root}

	A := node
	C := node.next

	novoNode.prev = A
	novoNode.next = C
	A.next = novoNode
	C.prev = novoNode

	l.size++
}

func (l *LList) PushBack(value int) {
	l.Insert(l.root.prev, value)
}

func (l *LList) PushFront(value int) {
	l.Insert(l.root.prev.next, value)
}

func (l *LList) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
}

func (l *LList) Size() int {
	return l.size
}

func (l *LList) String() string {
	if l.size == 0 {
		return "[]"
	}

	saida := "["

	atual := l.root.next

	for atual != l.root {
		saida += fmt.Sprintf("%d", atual.Value)

		if atual.next != l.root {
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
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
			 	fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
			 	fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			// oldvalue, _ := strconv.Atoi(args[1])
			// newvalue, _ := strconv.Atoi(args[2])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	node.Value = newvalue
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
			 	ll.Insert(node, newvalue)
			} else {
			 	fmt.Println("fail: not found")
			}
		case "remove":
			// oldvalue, _ := strconv.Atoi(args[1])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Remove(node)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
