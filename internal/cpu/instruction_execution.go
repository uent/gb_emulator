package cpu

import "fmt"

// GetInstructionFunc returns the execution function for the given opcode
func (cpu *Cpu) GetInstructionFunc(opcode byte) *Instruction {
	var instruction *Instruction
	var exists bool
	var realOpcode byte

	if opcode != 0xCB { // check if is 16 bit opt code
		instruction, exists = InstructionTable[opcode]
	} else {
		realOpcode = cpu.Read(cpu.PC + 1)
		instruction, exists = AdvancedInstructionTable[realOpcode]
	}

	if !exists {
		//time.Sleep(1000 * time.Second)
		panic(fmt.Sprintf("Invalid opcode: %02X, advanced opcode %02X", opcode, realOpcode))
		// If the opcode is not found, return nil
		//return nil
	}
	return instruction

}
