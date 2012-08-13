package main

import (
	"fmt"
)

type Node struct {
	Parent   *Node
	Name     string
	Children []*Node
}

func WhosNext(c string) string {
	if c == "alo" {
		return "alo2"
	}
	return ""
}

func FindNamesakeChild(n *Node, name string, ch chan *Node) {
	for _, v := range n.Children {
		if v.Name == name {
			n := WhosNext(v.Name)
			if n != "" {
				FindNamesakeChild(v, n, ch)
			}
			ch <- v
		}
	}
}

func FindRec(n *Node, name string) chan *Node {
	ch := make(chan *Node)
	go func() {
		FindNamesakeChild(n, name, ch)
		close(ch)
	}()
	return ch
}

func PrintHierarchy(n *Node) {
	for {
		if n == nil {
			break
		}
		fmt.Printf("%s<---", n.Name)
		n = n.Parent
	}
	fmt.Println()
	fmt.Println("---------------------------------------")
}

func main() {
	root := &Node{Name: "root"}

	one := &Node{Name: "alo", Parent: root}
	two := &Node{Name: "alo2", Parent: one}

	three := &Node{Name: "alo", Parent: root}
	four := &Node{Name: "alo2", Parent: three}

	root.Children = append(root.Children, one)
	root.Children = append(root.Children, three)

	one.Children = append(one.Children, two)
	three.Children = append(one.Children, four)

	ch := FindRec(root, "alo")
	for {
		n, ok := <-ch
		if !ok {
			break
		}
		if n.Name == "alo2" {
			PrintHierarchy(n)
		}

	}
}
