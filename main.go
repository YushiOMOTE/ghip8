package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var Terminated = errors.New("terminated")

type Game struct {
	rom_picker  RomPicker
	game_runner GameRunner
}

func (g *Game) Update() error {
	if g.game_runner.IsRunning() {
		if g.game_runner.Update() {
			g.game_runner.Stop()
		}
	} else {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			return Terminated
		}

		if g.rom_picker.Update() {
			rom := g.rom_picker.LoadRom()
			g.game_runner.Start(rom)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.game_runner.IsRunning() {
		g.game_runner.Draw(screen)
	} else {
		g.rom_picker.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	if g.game_runner.IsRunning() {
		return g.game_runner.ScreenSize()
	} else {
		return outsideWidth, outsideHeight
	}
}

func main() {
	picker := NewRomPicker()
	runner := NewGameRunner()

	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("Chip-8")

	if err := ebiten.RunGame(&Game{picker, runner}); err != nil {
		switch err {
		case Terminated:
			fmt.Printf("Terminated")
		default:
			log.Fatal(err)
		}
	}
}
