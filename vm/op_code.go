package vm

const (
	NULL int8 = iota
	ICONST
	BCONST			// load int8
	SCONST			// load string
	MOVE
	GMOVE			// move global variable
	IFZERO
	GOTO
	CALL
	RETURN
	SAVE			// save all registers
	RESTORE			// restore all rigisters
	NEG
	ADD
	SUB
	MUL
	DIV
	REM
	EQUAL
	MORE
	LESS
)

const (
	MAX_INT8 = 1 << 7 - 1
	MAX_INT16 = 1 << 15 - 1
)

func encodeRegister(reg int) int8 {
	if reg > NUM_OF_REG {
		panic("to many registers required")
	} else {
		return int8(-(reg + 1))
	}
}

func decodeRegister(operand int8) int {
	return int(-1 - operand)
}

func encodeOffset(offset int) int8 {
	if offset > MAX_INT8 {
		panic("too big int8 offset")
	} else {
		return int8(offset)
	}
}

func decodeOffset(operand int8) int {
	return int(operand)
}

func encodeShortOffset(offset int) int16 {
	if offset > MAX_INT16 {
		panic("too big int8 offset")
	} else {
		return int16(offset)
	}
}

func decodeShortOffset(operand int16) int {
	return int(operand)
}

func isRegister(operand int8) bool {
	return operand < 0
}

func isOffset(operand int8) bool {
	return operand >= 0
}
