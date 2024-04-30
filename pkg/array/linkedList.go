package array

import (
	"fmt"
	"github.com/mengub/baogo/pkg/baoFmt"
	"reflect"
)

type LinkedList struct {
	cap    int
	length int
	Type   reflect.Type // the type of value in LinkedList
	head   *node
	tail   *node
	index  map[int]*node // INDEX to RemoveBy Index
}

type node struct {
	next  *node
	prev  *node
	value any
}

func (l *LinkedList) checkType(value any) {
	if l.length == 0 {
		l.Type = reflect.TypeOf(value)
	} else {
		valueType := reflect.TypeOf(value)
		if valueType != l.Type {
			l.throwPanic(fmt.Sprintf("%v IS NOT SUPPORT", valueType))
		}
	}
}

func (l *LinkedList) throwPanic(panicStr string) {
	baseStr := fmt.Sprintf("linkedList[]%v --- %s", l.Type, panicStr)
	panic(baseStr)
}

func (l *LinkedList) Append(value any) {
	l.checkType(value)
	if l.length == 0 {
		l.head = &node{
			next:  nil,
			prev:  nil,
			value: nil,
		}
		l.tail = &node{
			next:  nil,
			prev:  nil,
			value: nil,
		}
		node := &node{
			next:  l.tail,
			prev:  l.head,
			value: value,
		}
		l.head.next = node
		l.tail.prev = node
	} else {
		node := &node{
			next:  l.tail,
			prev:  l.tail.prev,
			value: value,
		}
		node.prev.next = node
		node.next.prev = node
	}
	l.length++
}

func (l *LinkedList) Remove(value any) {
	l.checkType(value)
	p := l.head
	for p.next != nil && p.next.value != nil {
		if p.value == value {
			p.prev.next = p.next
			p.next.prev = p.prev
			return
		}
		p = p.next
	}
}

func (l *LinkedList) RemoveByIndex(index int) {
	if index > l.length {
		l.throwPanic("OUT OF INDEX")
	}
}

func (l *LinkedList) Foreach() {
	p := l.head
	for p.next != nil && p.next.value != nil {
		p = p.next
		baoFmt.PrintLn(p.value)
	}
}
