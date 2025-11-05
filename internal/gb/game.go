package gb

import "github.com/hajimehoshi/ebiten/v2"

// Game implements ebiten.Game for the NES emulator
type Game struct {
	gb *GB
	//renderer *ppu.Renderer
	paused bool
}

// NewGame creates a new Game instance
func NewGame(gb *GB) *Game {
	return &Game{
		gb: gb,
		//renderer: ppu.NewRenderer(gb.PPU),
		paused: false,
	}
}

// StartGame initializes and starts the NES game
func StartGame(gb *GB) error {
	//game := NewGame(gb)

	// Configure window
	ebiten.SetWindowSize(512, 480) // 256x240 scaled by 2
	ebiten.SetWindowTitle("GB Emulator")

	// Run the game
	return nil //ebiten.RunGame(game)
}
