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
	0x05: {Opcode: 0x05, Mnemonic: "DecB", IsIllegal: false, ExecuteFunc: DecB},
	0x06: {Opcode: 0x06, Mnemonic: "LDBImmediate", IsIllegal: false, ExecuteFunc: LDBImmediate},
	0x0C: {Opcode: 0x0C, Mnemonic: "INCC", IsIllegal: false, ExecuteFunc: INCC},
	0x0E: {Opcode: 0x0E, Mnemonic: "LDCImmediate", IsIllegal: false, ExecuteFunc: LDCImmediate},
	0x11: {Opcode: 0x11, Mnemonic: "LDDEd16", IsIllegal: false, ExecuteFunc: LDDEd16},
	0x17: {Opcode: 0x17, Mnemonic: "RLA", IsIllegal: false, ExecuteFunc: RLA},
	0x1A: {Opcode: 0x1A, Mnemonic: "LDADE", IsIllegal: false, ExecuteFunc: LDADE},
	0x20: {Opcode: 0x20, Mnemonic: "JRNZ", IsIllegal: false, ExecuteFunc: JRNZ},
	0x21: {Opcode: 0x21, Mnemonic: "LDHLImmediate", IsIllegal: false, ExecuteFunc: LDHLImmediate},
	0x26: {Opcode: 0x26, Mnemonic: "LDHImmediate", IsIllegal: false, ExecuteFunc: LDHImmediate},
	0x31: {Opcode: 0x31, Mnemonic: "LDImmediate", IsIllegal: false, ExecuteFunc: LDSPImmediate},
	0x32: {Opcode: 0x32, Mnemonic: "LDHL_A", IsIllegal: false, ExecuteFunc: LDHL_A},
	0x3E: {Opcode: 0x3E, Mnemonic: "LDAImmediate", IsIllegal: false, ExecuteFunc: LDAImmediate},
	0x40: {Opcode: 0x40, Mnemonic: "LDBBRegister", IsIllegal: false, ExecuteFunc: LDBBRegister},
	0x41: {Opcode: 0x41, Mnemonic: "LDBCRegister", IsIllegal: false, ExecuteFunc: LDBCRegister},
	0x45: {Opcode: 0x45, Mnemonic: "LDBL", IsIllegal: false, ExecuteFunc: LDBL},
	0x4F: {Opcode: 0x4F, Mnemonic: "LDCA", IsIllegal: false, ExecuteFunc: LDCA},
	0x77: {Opcode: 0x77, Mnemonic: "LDHLA", IsIllegal: false, ExecuteFunc: LDHLA},
	0xAF: {Opcode: 0xAF, Mnemonic: "XORA", IsIllegal: false, ExecuteFunc: XORA},
	0xC1: {Opcode: 0xC1, Mnemonic: "PopBC", IsIllegal: false, ExecuteFunc: PopBC},
	0xC5: {Opcode: 0xC5, Mnemonic: "PUSHBC", IsIllegal: false, ExecuteFunc: PUSHBC},
	0xCD: {Opcode: 0xCD, Mnemonic: "CALLa16", IsIllegal: false, ExecuteFunc: CALLa16},
	0xE0: {Opcode: 0xE0, Mnemonic: "LDa8AImmediate", IsIllegal: false, ExecuteFunc: LDa8AImmediate},
	0xE2: {Opcode: 0xE2, Mnemonic: "LD_C_A", IsIllegal: false, ExecuteFunc: LD_C_A},
	//0xCB: &Instruction{Opcode: 0xAF, Mnemonic: "TwoByteInstruction", IsIllegal: false, ExecuteFunc: TwoByteInstruction},
}

var AdvancedInstructionTable = map[byte]*Instruction{
	//0x20: &Instruction{Opcode: 0x20, Mnemonic: "SLAB", IsIllegal: false, ExecuteFunc: SLAB},
	0x11: {Opcode: 0x1, Mnemonic: "RLC", IsIllegal: false, ExecuteFunc: RLC},
	0x7C: {Opcode: 0x7C, Mnemonic: "Bit7H", IsIllegal: false, ExecuteFunc: Bit7H},
}
