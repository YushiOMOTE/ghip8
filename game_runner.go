package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameRunner struct {
	stop    chan struct{}
	running bool
	cpu     *CPU
}

func NewGameRunner() GameRunner {
	return GameRunner{make(chan struct{}), false, nil}
}

func (g *GameRunner) Start(rom []byte) {
	fmt.Println("Starting game")

	g.cpu = NewCPU(rom)
	g.running = true

	go func() {
		for {
			select {
			case <-time.Tick(time.Millisecond):
				g.cpu.Run()
			case <-g.stop:
				fmt.Println("Stopped game")
				g.running = false
				break
			}
		}
	}()
}

func (g *GameRunner) IsRunning() bool {
	return g.running
}

func (g *GameRunner) Stop() {
	fmt.Println("Stopping game")
	g.stop <- struct{}{}
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

func (g *GameRunner) HandleInput() bool {
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

	return false
}

func (g *GameRunner) Draw(screen *ebiten.Image) {
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
