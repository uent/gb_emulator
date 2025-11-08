package cpu

// 0xBC7C: Copy the complement of the contents of bit 7 in register H to the Z flag of the program status word (PSW).
func Bit7H(cpu *Cpu) uint8 {
	bit := (cpu.H >> 7) & 1

	if bit == 0 {
		cpu.ZFlag = true
	} else {
		cpu.ZFlag = false
	}

	cpu.NFlag = false
	cpu.HFlag = true

	cpu.MovePC(2)
	return 2
}

// 0xCB20: Shift the contents of register B to the left. That is, the contents of bit 0 are copied to bit 1, and the previous contents of bit 1 (before the copy operation) are copied to bit 2. The same operation is repeated in sequence for the rest of the register. The contents of bit 7 are copied to the CY flag, and bit 0 of register B is reset to 0.
//func SLAB(cpu *Cpu) uint8 {
//	cpu.CFlag = ((cpu.B >> 7) & 1) == 1 // set true or false
//
//	cpu.B = cpu.B << 1
//
//		cpu.ZFlag = cpu.B == 0
//		cpu.NFlag = false
//		cpu.HFlag = false
//
//	cpu.MovePC(2)
//	return 2
//}
