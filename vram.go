package main

// Represents video RAM with two-colored pixels
type VRAM struct {
	Width  int
	Height int
	buffer []bool
}

// Creates VRAM instance
func NewVRAM() VRAM {
	w, h := 64, 32

	return VRAM{
		Width:  w,
		Height: h,
		buffer: make([]bool, w*h),
	}
}

// Set the pixel in the video RAM
func (vram *VRAM) Set(x int, y int, v bool) bool {
	x %= vram.Width
	y %= vram.Height
	dot := &vram.buffer[x+y*vram.Width]
	erased := *dot && *dot == v
	*dot = *dot != v
	return erased
}

// Get the pixel in the video RAM
func (vram *VRAM) Get(x int, y int) bool {
	return vram.buffer[x+y*vram.Width]
}

// Clear all the pixel in the video RAM
func (vram *VRAM) Clear() {
	for i := 0; i < len(vram.buffer); i++ {
		vram.buffer[i] = false
	}
}
