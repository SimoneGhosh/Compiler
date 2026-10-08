package code

import (
	"encoding/binary"
	"fmt"
)

// Instructions is a slice of bytes
// no Instruction because easier to use []byte
type Instructions []byte

// Opcode is a byte
type Opcode byte

// Opcode defn
//  Each definition will have an Op prefix and the
//  value it refers to will be determined by iota
const (
	OpConstant Opcode = iota
)

// Bytecode defined in compiler package

// OpConstant has one operand
type Definition struct {
	Name		  string // make readable
	OperandWidths []int  // number of bytes for each operand
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}

// create a single bytecode instruction
// that’s made up of an Opcode and an optional number of operands
func Make(op Opcode, operands ...int) []byte{
	def, ok := definitions[op]
	if !ok{
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}

	return instruction
}