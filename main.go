package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var Terminated = errors.New("terminated")

var KeyMap = [...]ebiten.Key{
	ebiten.KeyX,
	ebiten.Key1,
	ebiten.Key2,
	ebiten.Key3,
	ebiten.KeyQ,
	ebiten.KeyW,
	ebiten.KeyE,
	ebiten.KeyA,
	ebiten.KeyS,
	ebiten.KeyD,
	ebiten.KeyZ,
	ebiten.KeyC,
	ebiten.Key4,
	ebiten.KeyR,
	ebiten.KeyF,
	ebiten.KeyV,
}

type Game struct {
	cpu CPU
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return Terminated
	}

	for i, k := range KeyMap {
		if inpututil.IsKeyJustPressed(k) {
			g.cpu.Keyboard.Press(uint8(i))
		}
		if inpututil.IsKeyJustReleased(k) {
			g.cpu.Keyboard.Release(uint8(i))
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for x := 0; x < g.cpu.Vram.Width; x++ {
		for y := 0; y < g.cpu.Vram.Height; y++ {
			if g.cpu.Vram.Get(x, y) {
				screen.Set(x, y, color.White)
			} else {
				screen.Set(x, y, color.Black)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cpu.Vram.Width, g.cpu.Vram.Height
}

func main() {
	romfile := os.Args[1]

	f, err := os.Open(romfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rom := make([]byte, 4096)

	count, err := f.Read(rom)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Loaded %d bytes from %s\n", count, romfile)

	cpu := NewCPU(rom)

	go func() {
		for range time.Tick(time.Millisecond) {
			cpu.Run()
		}
	}()

	ebiten.SetWindowSize(640, 320)
	ebiten.SetWindowTitle("Chip-8")

	if err := ebiten.RunGame(&Game{cpu}); err != nil {
		switch err {
		case Terminated:
			fmt.Printf("Terminated")
		default:
			log.Fatal(err)
		}
	}
}
