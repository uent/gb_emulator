package cpu

func (cpu *Cpu) pushStack(value uint8) {
	cpu.SP--

	cpu.Write(cpu.SP, value)
}

func (cpu *Cpu) pushWordStack(value uint16) {

	cpu.pushStack(uint8(value >> 8))   // High byte
	cpu.pushStack(uint8(value & 0xFF)) // Low byte
}

func (cpu *Cpu) popStack() uint8 {

	value := cpu.Read(cpu.SP)

	cpu.SP++

	return value
}

func (cpu *Cpu) popWordStack() uint16 {
	low := cpu.popStack()
	high := cpu.popStack()

	return uint16(high<<8) | uint16(low)
}
