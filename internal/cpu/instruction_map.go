package cpu

type CpuOperation func(*Cpu) uint8

type Instruction struct {
	Opcode      byte         // The instruction's opcode
	Mnemonic    string       // Instruction mnemonic (e.g., "LDA", "STA")
	IsIllegal   bool         // Whether it's an illegal/unofficial opcode
	ExecuteFunc CpuOperation // Function to execute the instruction
}

var InstructionTable = map[byte]*Instruction{
	0x00: {Opcode: 0x00, Mnemonic: "NOP", IsIllegal: false, ExecuteFunc: NOP},
	0x06: {Opcode: 0x06, Mnemonic: "LDBImmediate", IsIllegal: false, ExecuteFunc: LDBImmediate},
	0x0C: {Opcode: 0x0C, Mnemonic: "INCC", IsIllegal: false, ExecuteFunc: INCC},
	0x20: {Opcode: 0x20, Mnemonic: "JRNZ", IsIllegal: false, ExecuteFunc: JRNZ},
	0x21: {Opcode: 0x21, Mnemonic: "LDHLImmediate", IsIllegal: false, ExecuteFunc: LDHLImmediate},
	0x26: {Opcode: 0x26, Mnemonic: "LDHImmediate", IsIllegal: false, ExecuteFunc: LDHImmediate},
	0x31: {Opcode: 0x31, Mnemonic: "LDImmediate", IsIllegal: false, ExecuteFunc: LDSPImmediate},
	0x32: {Opcode: 0x32, Mnemonic: "LDHL_A", IsIllegal: false, ExecuteFunc: LDHL_A},
	0x3E: {Opcode: 0x3E, Mnemonic: "LDAImmediate", IsIllegal: false, ExecuteFunc: LDAImmediate},
	0x40: {Opcode: 0x40, Mnemonic: "LDBBRegister", IsIllegal: false, ExecuteFunc: LDBBRegister},
	0x41: {Opcode: 0x41, Mnemonic: "LDBCRegister", IsIllegal: false, ExecuteFunc: LDBCRegister},
	0x0E: {Opcode: 0x0E, Mnemonic: "LDCImmediate", IsIllegal: false, ExecuteFunc: LDCImmediate},
	0xAF: {Opcode: 0xAF, Mnemonic: "XORA", IsIllegal: false, ExecuteFunc: XORA},
	0xE2: {Opcode: 0xE2, Mnemonic: "LDCA", IsIllegal: false, ExecuteFunc: LDCA},
	//0xCB: &Instruction{Opcode: 0xAF, Mnemonic: "TwoByteInstruction", IsIllegal: false, ExecuteFunc: TwoByteInstruction},
}

var AdvancedInstructionTable = map[byte]*Instruction{
	//0x20: &Instruction{Opcode: 0x20, Mnemonic: "SLAB", IsIllegal: false, ExecuteFunc: SLAB},
	0x7C: {Opcode: 0x7C, Mnemonic: "Bit7H", IsIllegal: false, ExecuteFunc: Bit7H},
}
