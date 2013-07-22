package set

type set interface {
	add(elem interface{})
	remove(elem interface{}) bool
	size() int
	contains(elem interface{}) bool
}
