package main

import (
	"fmt"
	"math/rand"
	"time"
)

type lNode struct {
	data int
	next *lNode
}

func (l *lNode) newNextNode(data int) {
	l.next = &lNode{data: data}
}

func (l *lNode) forEachData(f func(data int) bool) {
	ctn := f(l.data)
	if !ctn {
		return
	}
	if l.next == nil {
		return
	}
	l.next.forEachData(f)
}

func (l *lNode) print() {
	fmt.Printf("List len = %d\n", l.data)
	p := l.next
	for p != nil {
		fmt.Printf("%d ", p.data)
		p = p.next
	}
	fmt.Println()
}

func newList(ltFigure, length int) *lNode {
	l := &lNode{}
	p := l
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		num := rand.Int() % ltFigure
		p.newNextNode(num)
		p = p.next
		l.data += 1
	}
	return l
}
