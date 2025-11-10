package memory

import "fmt"

// TODO: implement MBC1

//https://bgb.bircd.org/pandocs.htm

const (
	InitialAddress = 0x0000

	RomBank0Size             = 0x4000
	RomBank0SizeStartAddress = InitialAddress

	SwitchableRomBankSize         = 0x4000
	SwitchableRomBankStartAddress = RomBank0SizeStartAddress + RomBank0Size

	VideoRamSize         = 0x2000
	VideoRamStartAddress = SwitchableRomBankStartAddress + SwitchableRomBankSize

	SwitchableRamBankSize         = 0x2000 // external ram
	SwitchableRamBankStartAddress = VideoRamStartAddress + VideoRamSize

	InternalRamSize         = 0x1000
	InternalRamStartAddress = SwitchableRamBankStartAddress + SwitchableRamBankSize

	SwitchableRamSize         = 0x1000 // WRAM Bank 1 (or n)
	SwitchableRamStartAddress = InternalRamStartAddress + InternalRamSize

	EchoInternalRamSize         = 0x1E00
	EchoInternalRamStartAddress = SwitchableRamStartAddress + SwitchableRamSize

	OAMSize         = 0xA0
	OAMStartAddress = EchoInternalRamStartAddress + EchoInternalRamSize

	EmptyIO1Size         = 0x60
	EmptyIO1StartAddress = OAMStartAddress + OAMSize

	IOPortSize         = 0x7F
	IOPortStartAddress = EmptyIO1StartAddress + EmptyIO1Size

	HighRamSize         = 0x7F // high speed ram
	HighRamStartAddress = IOPortStartAddress + IOPortSize

	IESize         = 0x0001
	IEStartAddress = HighRamStartAddress + HighRamSize
)

// Memory represents the memory system of the GB
type Memory struct {
	BootRomBank0      [RomBank0Size]byte
	RomBank0          [RomBank0Size]byte // reemplace BootRomBank0
	SwitchableRomBank [SwitchableRomBankSize]byte
	VideoRam          [VideoRamSize]byte
	SwitchableRamBank [SwitchableRamBankSize]byte
	InternalRam       [InternalRamSize]byte
	SwitchableRam     [SwitchableRamSize]byte
	EchoInternalRam   [EchoInternalRamSize]byte // not used
	OAM               [OAMSize]byte
	EmptyIO1          [EmptyIO1Size]byte // undetermined
	IOPort            [IOPortSize]byte
	HighRam           [HighRamSize]byte
	IE                [IESize]byte // Interrupt Enable Register (IE)
	Boot              bool         // if is true, the console is booting
}

// New creates a new Memory instance
func New() *Memory {
	// Create a new memory instance
	m := &Memory{Boot: true}

	return m
}

// Return the memory address of the requested byte
func (m *Memory) getMemoryAddress(address uint16) *byte {

	switch {
	case address < SwitchableRomBankStartAddress:
		if m.Boot {
			return &m.BootRomBank0[address]
		} else {
			return &m.RomBank0[address]
		}
	case address < VideoRamStartAddress:
		return &m.SwitchableRomBank[address-SwitchableRomBankStartAddress]
	case address < SwitchableRamBankStartAddress:
		return &m.VideoRam[address-VideoRamStartAddress]
	case address < InternalRamStartAddress:
		return &m.SwitchableRamBank[address-SwitchableRamBankStartAddress]
	case address < SwitchableRamStartAddress:
		return &m.InternalRam[address-InternalRamStartAddress]
	case address < EchoInternalRamStartAddress:
		return &m.SwitchableRam[address-SwitchableRamStartAddress]
	case address < OAMStartAddress:
		offset := address - EchoInternalRamStartAddress
		if offset < InternalRamSize {
			return &m.InternalRam[offset]
		} else {
			return &m.SwitchableRam[offset-InternalRamSize]
		}
	case address < EmptyIO1StartAddress:
		return &m.OAM[address-OAMStartAddress]
	case address < IOPortStartAddress:
		return &m.EmptyIO1[address-EmptyIO1StartAddress]
	case address < HighRamStartAddress:
		return &m.IOPort[address-IOPortStartAddress]
	case address < IEStartAddress:
		return &m.HighRam[address-HighRamStartAddress]
	case address < 0xFFFF:
		return &m.IE[0]
	default:
		fmt.Printf("Invalid memory address: %x \n", address)
		panic("Invalid memory address")
	}

}

// Read returns a byte from the specified memory address
func (m *Memory) Read(address uint16) byte {
	return *m.getMemoryAddress(address)

}

// Write writes a byte to the specified memory address
func (m *Memory) Write(address uint16, value byte) {
	memoryValue := m.getMemoryAddress(address)

	*memoryValue = value
}

// ReadWord reads a 16-bit word from the specified memory address
// gb is little-endian, so the first byte is the low byte
func (m *Memory) ReadWord(address uint16) uint16 {
	low := uint16(m.Read(address))
	high := uint16(m.Read(address + 1))
	return (high << 8) | low
}

// WriteWord writes a 16-bit word to the specified memory address
func (m *Memory) WriteWord(address uint16, value uint16) {
	m.Write(address, byte(value&0xFF))
	m.Write(address+1, byte(value>>8))
}
