package cpu

import (
	"fmt"
	"gb-emulator/internal/memory"
)

// Documentation
// * https://gbdev.io/pandocs/CPU_Registers_and_Flags.html

type Cpu struct {
	PC    uint16 // Program Counter/Pointer
	SP    uint16 // Stack Pointer
	A     byte   // high part of the AF register
	B     byte
	C     byte
	D     byte
	E     byte
	H     byte
	L     byte
	ZFlag bool // bit 7 of AF, Zero flag
	NFlag bool // bit 6 of AF, Subtraction flag (BCD)
	HFlag bool // bit 5 of AF, Half Carry flag (BCD)
	CFlag bool // bit 4 of AF, also CY, also carry flag

	memory.Memory
}

// AF
//7	z	Zero flag
//6	n	Subtraction flag (BCD)
//5	h	Half Carry flag (BCD)
//4	c	Carry flag

// NewCPU creates a new CPU instance
func NewCPU() *Cpu {
	memoryInstance := memory.Memory{}
	cpu := &Cpu{Memory: memoryInstance}

	cpu.initParams()

	return cpu
}

func (cpu *Cpu) initParams() {
	cpu.PC = 0
	cpu.SP = 0xFFFE
	cpu.Boot = true
	// TODO
}

// Step executes a single CPU instruction
func (c *Cpu) Step() (uint8, error) {

	// Read opcode
	var cycles uint8
	opcode := c.Memory.Read(c.PC)

	// Get the execution function for the instruction
	executeFunc := c.GetInstructionFunc(opcode)
	if executeFunc != nil {
		cycles = executeFunc.ExecuteFunc(c)
	} else {
		return 0, fmt.Errorf("missing method for instruction opcode: %02X", opcode)
	}

	//c.PC++

	return cycles, nil // Return cycles used and no error
}

func (c *Cpu) MovePC(offset uint16) {
	c.PC = c.PC + offset
}

// Arithmetic Logic Unit
type Alu struct {
}

// Increment/Decrement Unit
type Idu struct {
}
