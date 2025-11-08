package cpu

// Documentation:
// * https://meganesu.github.io/generate-gb-opcodes/
// * https://gekkio.fi/files/gb-docs/gbctr.pdf
// * https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#LD_SP,n16

// 0x00: No operation. increment the pc on 1
func NOP(cpu *Cpu) uint8 {

	return 0
}

// 0x41: Load the contents of register C into register B.
func LDBCRegister(cpu *Cpu) uint8 {
	cpu.B = cpu.C

	cpu.MovePC(1)
	return 1
}

// 0x40: Load the contents of register B into register B.
func LDBBRegister(cpu *Cpu) uint8 {
	cpu.B = cpu.B

	cpu.MovePC(1)
	return 1
}

// 0x31: Load the 2 bytes of immediate data into register pair SP.
func LDSPImmediate(cpu *Cpu) uint8 {
	cpu.SP = cpu.ReadWord(cpu.PC + 1)

	cpu.MovePC(3)
	return 3
}

// 0x06: Load the 8-bit immediate operand d8 into register B.
func LDBImmediate(cpu *Cpu) uint8 {
	cpu.B = cpu.Read(cpu.PC + 1)

	cpu.MovePC(2)
	return 2
}

// 0x20: If the Z flag is 0, jump s8 steps from the current address stored in the program counter (PC). If not, the instruction following the current JP instruction is executed (as usual).
func JRNZ(cpu *Cpu) uint8 {
	var cycles uint8

	if !cpu.ZFlag {
		cycles = 3
		cpu.PC = cpu.PC + 2 + uint16(int8(cpu.Read(cpu.PC+1)))
	} else {
		cycles = 2
		cpu.MovePC(2)
	}

	return cycles
}

// 0x21: Load the 2 bytes of immediate data into register pair HL.
func LDHLImmediate(cpu *Cpu) uint8 {
	value := cpu.ReadWord(cpu.PC + 1)

	high, low := splitUInt16ToBytes(value)

	cpu.H = high
	cpu.L = low

	cpu.MovePC(3)
	return 3
}

// 0x26: Load the 8-bit immediate operand d8 into register H.
func LDHImmediate(cpu *Cpu) uint8 {
	value := cpu.Read(cpu.PC + 1)

	cpu.H = value

	cpu.MovePC(2)
	return 2
}

// 0x32: Store the contents of register A into the memory location specified by register pair HL, and simultaneously decrement the contents of HL.
func LDHL_A(cpu *Cpu) uint8 {

	address := jointBytesToUInt16(cpu.H, cpu.L)

	cpu.Write(address, cpu.A)

	address--

	high, low := splitUInt16ToBytes(address)

	cpu.H = high
	cpu.L = low

	cpu.MovePC(1)
	return 2
}

// 0x3E: Load the 8-bit immediate operand d8 into register A.
func LDAImmediate(cpu *Cpu) uint8 {

	cpu.A = cpu.Read(cpu.PC + 1)

	cpu.MovePC(2)
	return 2
}

// 0x0C: Increment the contents of register C by 1.
func INCC(cpu *Cpu) uint8 {
	oldC := cpu.C
	cpu.C++

	cpu.ZFlag = cpu.C == 0
	cpu.NFlag = false
	cpu.HFlag = calculateHalfFlagIncrement(oldC)

	cpu.MovePC(1)
	return 1
}

// 0x0E: Load the 8-bit immediate operand d8 into register C.
func LDCImmediate(cpu *Cpu) uint8 {

	cpu.C = cpu.Read(cpu.PC + 1)

	cpu.MovePC(2)
	return 2
}

// 0xAF: Take the logical exclusive-OR for each bit of the contents of register A and the contents of register A, and store the results in register A.
func XORA(cpu *Cpu) uint8 {
	cpu.A ^= cpu.A // A = A XOR A = 0

	cpu.ZFlag = true
	cpu.NFlag = false
	cpu.HFlag = false
	cpu.CFlag = false

	cpu.MovePC(1)
	return 1
}

// 0xE2: Store the contents of register A in the internal RAM, port register, or mode register at the address in the range 0xFF00-0xFFFF specified by register C.
func LDCA(cpu *Cpu) uint8 {

	address := 0xFF00 + uint16(cpu.C)

	cpu.Write(address, cpu.A)

	cpu.MovePC(1)
	return 2
}
