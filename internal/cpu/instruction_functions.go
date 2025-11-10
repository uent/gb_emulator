package cpu

// Documentation:
// * https://meganesu.github.io/generate-gb-opcodes/
// * https://gekkio.fi/files/gb-docs/gbctr.pdf
// * https://rgbds.gbdev.io/docs/v0.9.4/gbz80.7#LD_SP,n16

// d8	"data 8-bit"	8 bits (1 byte)	Un valor inmediato de 8 bits que se usa como dato (por ejemplo, una constante que se carga en un registro).
// d16	"data 16-bit"	16 bits (2 bytes)	Un valor inmediato de 16 bits, usado como constante de 2 bytes (por ejemplo para direcciones o registros de 16 bits).
// a8	"address 8-bit"	8 bits (1 byte)	Un valor de dirección de 8 bits, que se usa junto con el prefijo 0xFF00. Es decir, la dirección final será 0xFF00 + a8.
// a16	"address 16-bit"	16 bits (2 bytes)	Una dirección absoluta de 16 bits (sin prefijo), usada directamente.

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

// 0x45: Load the contents of register L into register B.
func LDBL(cpu *Cpu) uint8 {
	cpu.B = cpu.L

	cpu.MovePC(1)
	return 1
}

// 0x40: Load the contents of register B into register B.
func LDBBRegister(cpu *Cpu) uint8 {
	cpu.B = cpu.B

	cpu.MovePC(1)
	return 1
}

// 0x4F: Load the contents of register A into register C.
func LDCA(cpu *Cpu) uint8 {
	cpu.C = cpu.A

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

// 0x11: Load the 2 bytes of immediate data into register pair DE.
func LDDEd16(cpu *Cpu) uint8 {

	high, low := splitUInt16ToBytes(cpu.ReadWord(cpu.PC + 1))

	cpu.D = high
	cpu.E = low

	cpu.MovePC(3)
	return 3
}

// 0x17: Rotate the contents of register A to the left, through the carry (CY) flag. That is, the contents of bit 0 are copied to bit 1, and the previous contents of bit 1 (before the copy operation) are copied to bit 2. The same operation is repeated in sequence for the rest of the register. The previous contents of the carry flag are copied to bit 0.
func RLA(cpu *Cpu) uint8 {

	lastBit := (cpu.A >> 7) & 1

	cpu.A = cpu.A << 1

	cpu.A = cpu.A | bool2u8(cpu.CFlag)

	cpu.ZFlag = false
	cpu.NFlag = false
	cpu.HFlag = false
	cpu.CFlag = lastBit == 1

	cpu.MovePC(1)
	return 1
}

// 0x1A: Load the 8-bit contents of memory specified by register pair DE into register A.
func LDADE(cpu *Cpu) uint8 {

	address := jointBytesToUInt16(cpu.D, cpu.E)

	cpu.A = cpu.Read(address)

	cpu.MovePC(1)
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

// 0x77: Store the contents of register A in the memory location specified by register pair HL.
func LDHLA(cpu *Cpu) uint8 {
	address := jointBytesToUInt16(cpu.H, cpu.L)

	cpu.Write(address, cpu.A)

	cpu.MovePC(1)
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

// 0xC5: Push the contents of register pair BC onto the memory stack by doing the following:
// * Subtract 1 from the stack pointer SP, and put the contents of the higher portion of register pair BC on the stack.
// *  Subtract 2 from SP, and put the lower portion of register pair BC on the stack.
// * Decrement SP by 2.
// TODO: check if the order is correct
func PUSHBC(cpu *Cpu) uint8 {

	cpu.pushStack(cpu.B)
	cpu.pushStack(cpu.C)

	cpu.MovePC(1)
	return 4
}

// 0xC1: Pop the contents from the memory stack into register pair into register pair BC by doing the following
// * Load the contents of memory specified by stack pointer SP into the lower portion of BC.
// * Add 1 to SP and load the contents from the new memory location into the upper portion of BC.
// * By the end, SP should be 2 more than its initial value.
func PopBC(cpu *Cpu) uint8 {

	value := cpu.popWordStack()

	high, low := splitUInt16ToBytes(value)

	cpu.B = high
	cpu.C = low

	cpu.MovePC(1)
	return 3
}

// 0xCD: In memory, push the program counter PC value corresponding to the address following the CALL instruction to the 2 bytes following the byte specified by the current stack pointer SP. Then load the 16-bit immediate operand a16 into PC.
func CALLa16(cpu *Cpu) uint8 {

	cpu.pushWordStack(cpu.PC + 3)

	cpu.PC = cpu.ReadWord(cpu.PC + 1)

	return 6
}

// 0xE0: Store the contents of register A in the internal RAM, port register, or mode register at the address in the range 0xFF00-0xFFFF specified by the 8-bit immediate operand a8.
func LDa8AImmediate(cpu *Cpu) uint8 {

	address := 0xFF00 + uint16(cpu.Read(cpu.PC+1))

	cpu.Write(address, cpu.A)

	cpu.MovePC(2)
	return 3
}

// 0xE2: Store the contents of register A in the internal RAM, port register, or mode register at the address in the range 0xFF00-0xFFFF specified by register C.
func LD_C_A(cpu *Cpu) uint8 {

	address := 0xFF00 + uint16(cpu.C)

	cpu.Write(address, cpu.A)

	cpu.MovePC(1)
	return 2
}
