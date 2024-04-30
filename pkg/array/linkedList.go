package array

import (
	"fmt"
	"reflect"
)

type LinkedList struct {
	cap    int
	length int
	Type   reflect.Type // the type of value in LinkedList
	head   *node
	tail   *node
}

type node struct {
	next  *node
	prev  *node
	value any
}

func (l LinkedList) checkType(value any) {
	if l.length == 0 {
		l.Type = reflect.TypeOf(value)
	} else {
		valueType := reflect.TypeOf(value)
		if valueType != l.Type {
			panic(fmt.Sprintf("%v IS NOT SUPPORT IN [%v]\n", valueType, l.Type))
		}
	}
}

func (l LinkedList) Append(value any) {
	l.checkType(value)
	if l.length == 0 {
		l.head.next = &node{
			next:  l.tail,
			prev:  l.head,
			value: value,
		}
	} else {
		// TODO 尾插法
		//l.tail.prev = &node{
		//	next:  l.tail,
		//	prev:  l.tail.prev,
		//	value: value,
		//}
		p := l.head
		for p.next != nil {

		}
		p.next = &node{
			next:  l.tail,
			prev:  p,
			value: value,
		}
	}
	l.length++
}

func (l LinkedList) Remove(value any) {
	l.checkType(value)

}

func (l LinkedList) RemoveByIndex(index int) {

}

func (l LinkedList) Foreach() {
	p := l.head
	for p.next != nil {
		fmt.Println(p.value)
		p = p.next
	}
}
