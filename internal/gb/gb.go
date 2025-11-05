// Package nes implements the NES system integration
package gb

import (
	"gb-emulator/internal/cpu"
)

// NES represents the Nintendo Entertainment System
type GB struct {
	Cpu *cpu.Cpu
	//PPU    *ppu.PPU
	//Memory *memory.Memory

	// System state
	Running bool
	//Cycles  uint64
}

// New creates a new NES instance
func New() *GB {
	cpuInstance := cpu.NewCPU()

	gb := &GB{
		Cpu: cpuInstance,
		//PPU:     ppu.NewPPU(),
		//Memory:  memory.New(),
		Running: false,
		//Cycles:  0,
	}

	// Connect components
	//nes.CPU.SetMemory(nes.Memory)
	//nes.PPU.SetCPU(nes.CPU)
	//nes.Memory.SetPPU(nes.PPU)

	return gb
}

// Reset resets the NES to its initial state
func (n *GB) Reset() {
	//n.Memory.Reset()
	//n.PPU.Reset()
	//n.CPU.Reset()
	//n.Cycles = 0
}

func (n *GB) LoadBootROM(bootRomData []byte) error {

	for i := 0; i < len(bootRomData); i++ {
		if i == len(n.Cpu.RomBank0) {
			panic("exceed memory")
		}

		n.Cpu.BootRomBank0[i] = bootRomData[i]
	}

	return nil
}

// LoadROM loads a ROM file into memory
func (n *GB) LoadROM(romData []byte) error {

	for i := 0; i < len(romData) && i < len(n.Cpu.RomBank0); i++ {
		if i == len(n.Cpu.RomBank0) {
			panic("exceed memory")
		}

		n.Cpu.RomBank0[i] = romData[i]
	}

	return nil
}

// Step advances the NES emulation by one CPU instruction
func (n *GB) Step() error {
	// Execute one CPU instruction
	_, err := n.Cpu.Step()
	if err != nil {
		return err
	}

	return nil
}

// Run runs the NES emulation until stopped
func (gb *GB) Run() error {
	gb.Running = true

	for gb.Running {
		err := gb.Step()
		if err != nil {
			return err
		}
	}

	return nil
}

// Stop stops the NES emulation
func (gb *GB) Stop() {
	gb.Running = false
}
