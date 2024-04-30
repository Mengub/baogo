package main

import "github.com/mengub/baogo/pkg/array"

func main() {
	var list array.LinkedList
	list.Append(1)
	list.Append("S")
	list.Foreach()
}
