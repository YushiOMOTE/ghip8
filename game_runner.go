package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameRunner struct {
	running bool
	freq    int
	time    time.Time
	cpu     *CPU
}

func NewGameRunner() GameRunner {
	return GameRunner{false, 1000, time.Now(), NewCPU(make([]byte, 0))}
}

func (g *GameRunner) Start(rom []byte) {
	if g.running {
		return
	}

	fmt.Println("Starting game")
	g.cpu = NewCPU(rom)
	g.running = true
}

func (g *GameRunner) IsRunning() bool {
	return g.running
}

func (g *GameRunner) Stop() {
	g.running = false
}

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

func (g *GameRunner) Update() bool {
	if !g.running {
		return false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return true
	}

	for i, k := range KeyMap {
		if inpututil.IsKeyJustPressed(k) {
			g.cpu.Keyboard.Press(uint8(i))
		}
		if inpututil.IsKeyJustReleased(k) {
			g.cpu.Keyboard.Release(uint8(i))
		}
	}

	dt := int(time.Since(g.time) / 1000)
	g.time = time.Now()

	for i := 0; i < dt*g.freq/1000000; i++ {
		g.cpu.Run()
	}

	return false
}

func (g *GameRunner) Draw(screen *ebiten.Image) {
	if !g.running {
		return
	}

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

func (g *GameRunner) ScreenSize() (int, int) {
	return g.cpu.Vram.Width, g.cpu.Vram.Height
}
