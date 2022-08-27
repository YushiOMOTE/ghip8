package main

type VRAM struct {
	Width  int
	Height int
	buffer []bool
}

func NewVRAM() VRAM {
	w, h := 64, 32

	return VRAM{
		Width:  w,
		Height: h,
		buffer: make([]bool, w*h),
	}
}

func (vram *VRAM) Set(x int, y int, v bool) bool {
	x %= vram.Width
	y %= vram.Height
	dot := &vram.buffer[x+y*vram.Width]
	erased := *dot && *dot == v
	*dot = *dot != v
	return erased
}

func (vram *VRAM) Get(x int, y int) bool {
	return vram.buffer[x+y*vram.Width]
}

func (vram *VRAM) Clear() {
	for i := 0; i < len(vram.buffer); i++ {
		vram.buffer[i] = false
	}
}
