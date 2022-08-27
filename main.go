package main

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	// "golang.org/x/image/colornames"
)

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

	fmt.Printf("Loaded %d bytes\n", count)

	cpu := NewCPU(rom)

	pixelgl.Run(func() {
		run(&cpu)
	})
}

var KeyMap = [...]pixelgl.Button{
	pixelgl.KeyX,
	pixelgl.Key1,
	pixelgl.Key2,
	pixelgl.Key3,
	pixelgl.KeyQ,
	pixelgl.KeyW,
	pixelgl.KeyE,
	pixelgl.KeyA,
	pixelgl.KeyS,
	pixelgl.KeyD,
	pixelgl.KeyZ,
	pixelgl.KeyC,
	pixelgl.Key4,
	pixelgl.KeyR,
	pixelgl.KeyF,
	pixelgl.KeyV,
}

func run(cpu *CPU) {
	cfg := pixelgl.WindowConfig{
		Title:  "Chip-8",
		Bounds: pixel.R(0, 0, 640, 320),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	go func() {
		for range time.Tick(time.Millisecond) {
			cpu.Run()
		}
	}()

	for !win.Closed() && !win.Pressed(pixelgl.KeyEscape) {
		for i, k := range KeyMap {
			if win.JustPressed(k) {
				cpu.Keyboard.Press(uint8(i))
			}
			if win.JustReleased(k) {
				cpu.Keyboard.Release(uint8(i))
			}
		}

		scale := 10.0
		imd := imdraw.New(nil)
		for x := 0; x < cpu.Vram.Width; x++ {
			for y := 0; y < cpu.Vram.Height; y++ {
				if cpu.Vram.Get(x, y) {
					imd.Color = color.White
				} else {
					imd.Color = color.Black
				}

				px := float64(x) * scale
				py := float64((cpu.Vram.Height - y - 1)) * scale
				imd.Push(pixel.V(px, py), pixel.V(px+scale, py+scale))
				imd.Rectangle(0.0)
			}
		}
		imd.Draw(win)

		win.Update()
	}
}
