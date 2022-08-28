package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type RomPicker struct {
	font        font.Face
	sel         int
	items       []string
	last_update time.Time
}

func NewRomPicker() RomPicker {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72

	font, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	return RomPicker{font, 0, make([]string, 0), time.Now()}
}

func (p *RomPicker) Add(item string) {
	p.items = append(p.items, item)
}

func (p *RomPicker) Next() {
	if p.sel < len(p.items)-1 {
		p.sel += 1
	}
}

func (p *RomPicker) Prev() {
	if p.sel > 0 {
		p.sel -= 1
	}
}

func (p *RomPicker) LoadRom() []byte {
	item := p.items[p.sel]

	rom, err := res.ReadFile("games/" + item)
	if err != nil {
		panic(err)
	}

	return rom
}

func (p *RomPicker) HandleInput() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return true
	}

	dt := time.Since(p.last_update)
	if dt < 50*time.Millisecond {
		return false
	}

	p.last_update = time.Now()

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.Prev()
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.Next()
	}

	return false
}

func (p *RomPicker) Draw(screen *ebiten.Image) {
	for i, name := range p.items {
		msg := fmt.Sprintf("%2d: %s", i, name)
		col := color.RGBA{255, 255, 255, 255}
		if p.sel == i {
			col = color.RGBA{255, 0, 0, 255}
		}
		text.Draw(screen, msg, p.font, 0, i*30+30-(p.sel*30), col)
	}
}
