package cpu

// LD r, r’: Load register (register)
func LDRegister(cpu *Cpu) uint8 {
	cpu.B = cpu.C

	return 0
}

// LD r, n: Load register (immediate)
func LDInmediate(cpu *Cpu) uint8 {
	cpu.B = cpu.Read(cpu.PC)

	return 0
}
