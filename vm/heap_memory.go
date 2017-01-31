package vm

type HeapMemory interface{
	read(index int) interface{}
	write(index int, v interface{})
}
