package array

type list interface {
	Append(value any)
	Remove(value any)
	RemoveByIndex(index int)
	Foreach()
	checkType(value any)
}
