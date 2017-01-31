package vm

import (
	"unsafe"
)

/*
	register address [-1 ~ -7]
	heap address [0 ~ 4G]
 */

const (
	NUM_OF_REG     = 6
	SAVE_AREA_SIZE = NUM_OF_REG + 2
	TRUE           = 1
	FLASE          = 0
)

type VM struct {
	code    []int8
	stack   []interface{}
	strings []string
	heap    HeapMemory

	pc, fp, sp, ret int // fp frame pointer; sp stack pointer
	registers       []interface{}
}

func NewVM(codeSize, stackSize, stringsSize int, hm HeapMemory) *VM {
	return &VM{code:   make([]int8, codeSize),
		stack:     make([]interface{}, stringsSize),
		strings:   make([]string, stringsSize),
		heap:      hm,
		registers: make([]interface{}, NUM_OF_REG)}
}

func (self *VM) GetReg(i int) interface{}        { return self.registers[i] }
func (self *VM) SetReg(i int, value interface{}) { self.registers[i] = value }
func (self *VM) Strings() []string               { return self.strings }
func (self *VM) Code() []int8                    { return self.code }
func (self *VM) Stack() []interface{}            { return self.stack }
func (self *VM) Heap() HeapMemory                { return self.heap }

func (self *VM) Run(entry int) {
	self.pc = entry
	self.fp = 0
	self.sp = 0

	for self.pc >= 0 {
		self.mainLoop()
	}
}

func (self *VM) mainLoop() {
	switch self.code[self.pc] {
	case ICONST:
	case BCONST:
	case SCONST:
	case MOVE:
	case GMOVE:
	case IFZERO:
	case GOTO:
	case CALL:
	case RETURN:
	case SAVE:
	case RESTORE:
	case NEG:
	default:

	}
}

func (self *VM) moveValue() {
	src, dest := self.code[self.pc + 1], self.code[self.pc + 2]
	var value interface{}
	if isRegister(src) {
		value = self.registers[decodeRegister(src)]
	} else {
		value = self.stack[self.fp + decodeOffset(src)]
	}

	if isRegister(dest) {
		self.registers[decodeRegister(src)] = value
	} else {
		self.stack[self.fp + decodeOffset(src)] = value
	}

	self.pc += 3
}

func (self *VM) moveHeapValue() {
	rand := self.code[self.pc + 1]
	if (isRegister(rand)) {
		dest := readShort(self.code, self.pc + 1)
		self.heap.Write(dest, self.registers[decodeRegister(rand)])
	} else {
		src := readShort(self.code, self.pc + 1)
		self.registers[decodeRegister(self.code[self.pc + 3])] = self.heap.Read(src)
	}
	self.pc += 4
}



func readInt(b []int8, i int) int {
	x := uint32(b[i + 3]) | uint32(b[i + 2]) << 8 | uint32(b[i + 1]) << 16 | uint32(b[i]) << 24
	return int(*(*int32)(unsafe.Pointer(&x)))
}

func readShort(b []int8, i int) int {
	x := uint32(b[i + 1]) | uint32(b[i]) << 8
	return int(*(*int16)(unsafe.Pointer(&x)))
}
