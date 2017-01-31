package vm

const (
	NUM_OF_REG = 6
	SAVE_AREA_SIZE = NUM_OF_REG + 2
	TRUE = 1
	FLASE = 0
)

type VM struct {
	code []byte
	stack []interface{}
	strings []string
	heap HeapMemory

	pc, fp, sp, ret int
	registers []interface{}
}


func NewVM(codeSize, stackSize, stringsSize int, hm HeapMemory) *VM {
	return &VM{code : make([]byte, codeSize),
		stack : make([]interface{}, stringsSize),
		strings : make([]string, stringsSize),
		heap : hm,
		registers : make([]interface{}, NUM_OF_REG)}
}

func (self *VM) Run(entry int) {

}

func (self *VM) mainLoop() {

}

