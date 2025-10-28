package cpu

import "fmt"

type CpuOperation func(*Cpu) uint8

type Instruction struct {
	Opcode      byte         // The instruction's opcode
	Mnemonic    string       // Instruction mnemonic (e.g., "LDA", "STA")
	IsIllegal   bool         // Whether it's an illegal/unofficial opcode
	ExecuteFunc CpuOperation // Function to execute the instruction
}

var InstructionTable = map[byte]*Instruction{
	0x41: &Instruction{Opcode: 0x41, Mnemonic: "LDRegister", IsIllegal: false, ExecuteFunc: LDRegister},
	0x06: &Instruction{Opcode: 0x06, Mnemonic: "LDInmediate", IsIllegal: false, ExecuteFunc: LDInmediate},
}

// GetInstructionFunc returns the execution function for the given opcode
func GetInstructionFunc(opcode byte) *Instruction {
	instruction, exists := InstructionTable[opcode]
	if !exists {
		//time.Sleep(1000 * time.Second)
		panic(fmt.Sprintf("Invalid opcode: %02X", opcode))
		// If the opcode is not found, return nil
		//return nil
	}
	return instruction
}
