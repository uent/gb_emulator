package memory

import "fmt"

//https://bgb.bircd.org/pandocs.htm

// -- 0000                       |
// 16kB ROM bank #0 |						 |
// -- 4000											 |
// 16kB switchable ROM bank			 |= 32kB Cartrigbe
// -- 8000											 |
// 8kB Video RAM
// -- A000
// 8kB switchable RAM bank
// -- C000
// 8kB Internal RAM
// -- E000
// Echo of 8kB Internal RAM
// -- FE00
// Sprite Attrib Memory (OAM)
// -- FEA0
// Empty but unusable for I/O
// -- FF00
// I/O ports
// -- FF4C
// Empty but unusable for I/O
// -- FF80
// High RAM
// -- FFFF
// Interrupt Enable Register

const (
	InitialAddress = 0x0000

	RomBank0Size             = 0x4000
	RomBank0SizeStartAddress = InitialAddress

	SwitchableRomBankSize         = 0x4000
	SwitchableRomBankStartAddress = RomBank0SizeStartAddress + RomBank0Size

	VideoRamSize         = 0x2000
	VideoRamStartAddress = SwitchableRomBankStartAddress + SwitchableRomBankSize

	SwitchableRamBankSize         = 0x2000
	SwitchableRamBankStartAddress = VideoRamStartAddress + VideoRamSize

	InternalRamSize         = 0x2000
	InternalRamStartAddress = SwitchableRamBankStartAddress + SwitchableRamBankSize

	EchoInternalRamSize         = 0x2000
	EchoInternalRamStartAddress = InternalRamStartAddress + InternalRamSize

	OAMSize         = 0xA0
	OAMStartAddress = EchoInternalRamStartAddress + EchoInternalRamSize

	EmptyIO1Size         = 0x60
	EmptyIO1StartAddress = OAMStartAddress + OAMSize

	IOPortSize         = 0x4C
	IOPortStartAddress = EmptyIO1StartAddress + EmptyIO1Size

	EmptyIO2Size         = 0x34
	EmptyIO2StartAddress = IOPortStartAddress + IOPortSize

	HighRamSize         = 0x7F // also internal ram?
	HighRamStartAddress = EmptyIO2StartAddress + EmptyIO2Size
)

// Memory represents the memory system of the NES
type Memory struct {
	RomBank0          [RomBank0Size]byte
	BootRomBank0      [RomBank0Size]byte
	SwitchableRomBank [SwitchableRomBankSize]byte
	VideoRam          [VideoRamSize]byte
	SwitchableRamBank [SwitchableRamBankSize]byte
	InternalRam       [InternalRamSize]byte
	EchoInternalRam   [EchoInternalRamSize]byte
	OAM               [OAMSize]byte
	EmptyIO1          [EmptyIO1Size]byte
	IOPort            [IOPortSize]byte
	EmptyIO2          [EmptyIO2Size]byte
	HighRam           [HighRamSize]byte
	Boot              bool // if is true, the console is booting
}

// New creates a new Memory instance
func New() *Memory {
	// Create a new memory instance
	m := &Memory{Boot: true}

	// Initialize RAM (addresses 0x0000 to 0x07FF) with zeros
	//for i := 0; i < 0x0800; i++ {
	//	m.InternalRam[i] = 0
	//}

	return m
}

// Read returns a byte from the specified memory address
func (m *Memory) Read(address uint16) byte {
	switch {
	case address < RomBank0Size:
		if m.Boot {
			return m.BootRomBank0[address]
		} else {
			return m.RomBank0[address]
		}
	case address < SwitchableRomBankSize:
		return m.SwitchableRomBank[address-SwitchableRomBankStartAddress]
	case address < VideoRamSize:
		return m.VideoRam[address-VideoRamStartAddress]
	case address < SwitchableRamBankSize:
		return m.SwitchableRamBank[address-SwitchableRamBankStartAddress]
	case address < InternalRamSize:
		return m.InternalRam[address-InternalRamStartAddress]
	default:
		fmt.Printf("Invalid memory address: %x \n", address)
		panic("Invalid memory address")
	}

}

// Write writes a byte to the specified memory address
func (m *Memory) Write(address uint16, value byte) {

}

// LoadPRGROM loads the program ROM into memory
//func (m *Memory) LoadPRGROM(prgROM []byte) {
// Copy PRG ROM data into the appropriate location in cartridge space
// Typically starting at 0x8000, but this is a simplification
// In a real implementation, this would depend on the mapper
//	for i := 0; i < len(prgROM) && i < len(m.ROMCartridgeSpace); i++ {
//		m.ROMCartridgeSpace[i] = prgROM[i]
//	}
//}

// ReadWord reads a 16-bit word from the specified memory address
// NES is little-endian, so the first byte is the low byte
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

func (m *Memory) ReadAddressIndirectPageBoundaryBug(address uint16) uint16 {
	var addr uint16
	if (address & 0x00FF) == 0x00FF { // Si el puntero termina en XXFF, aplica el bug
		low := m.Read(address)
		high := m.Read(address & 0xFF00) // Lee desde XX00 en vez de XXFF+1
		addr = uint16(high)<<8 | uint16(low)
	} else {
		addr = m.ReadWord(address) // Lectura normal
	}

	return addr
}
