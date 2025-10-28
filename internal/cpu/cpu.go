package cpu

import (
	"fmt"
	"gb-emulator/internal/memory"
)

type Cpu struct {
	PC uint16 // Program Counter/Pointer
	SP uint16 // Stack Pointer
	A  byte   // Accumulator
	R  byte   // Register
	IR byte   // instruction Register
	IE byte   // Interrupt Enable
	B  byte   // BC
	C  byte
	D  byte // DE
	E  byte
	H  byte // HL
	L  byte
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
	// TODO
}

// Step executes a single CPU instruction
func (c *Cpu) Step() (uint8, error) {
	// Check for interrupts first
	//if c.nmiPending || c.irqPending {
	//	return c.handleInterrupts(), nil
	//}

	// Read opcode
	var cycles uint8
	opcode := c.Memory.Read(c.PC)

	// Get the execution function for the instruction
	executeFunc := GetInstructionFunc(opcode)
	if executeFunc != nil {
		cycles = executeFunc.ExecuteFunc(c)
	} else {
		return 0, fmt.Errorf("missing method for instruction opcode: %02X", opcode)
	}

	return cycles, nil // Return cycles used and no error
}

// Arithmetic Logic Unit
type Alu struct {
}

// Increment/Decrement Unit
type Idu struct {
}
