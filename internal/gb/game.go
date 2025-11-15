package gb

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

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

// Update updates the game logic
func (g *Game) Update() error {
	if g.paused {
		return nil
	}

	// Run CPU cycles for one frame (approximately 70224 cycles for Game Boy)
	// This will be adjusted based on actual frame timing
	for i := 0; i < 70224; i++ {
		g.gb.Cpu.Step()
	}

	return nil
}

// Draw draws the game screen
func (g *Game) Draw(screen *ebiten.Image) {
	// TODO: Implement PPU rendering
	// For now, just clear the screen
	screen.Fill(color.Black)
}

// Layout returns the game's logical screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Game Boy screen resolution is 160x144
	return 160, 144
}

// StartGame initializes and starts the NES game
func StartGame(gb *GB) error {
	game := NewGame(gb)

	// Configure window
	ebiten.SetWindowSize(512, 480) // 256x240 scaled by 2
	ebiten.SetWindowTitle("GB Emulator")

	// Run the game
	return ebiten.RunGame(game)
}
