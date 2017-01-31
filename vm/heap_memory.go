package vm

type HeapMemory interface{
	Read(index int) interface{}
	Write(index int, v interface{})
}
