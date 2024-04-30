package array

type list interface {
	Append(value any)
	Remove(value any)
	RemoveByIndex(index int)
	Foreach()
}

type debug interface {
	checkType(value any)
	throwPanic(panicStr string)
}
