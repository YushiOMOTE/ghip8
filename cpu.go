package main

import (
	"math/rand"
)

type CPU struct {
	pc          uint16
	i           uint16
	v           []uint8
	mem         []uint8
	stack       Stack
	delay_timer uint8
	sound_timer uint8
	Vram        VRAM
	Keyboard    Keyboard
	time        int
}

func NewCPU(rom []byte) CPU {
	ep := uint16(0x200)

	c := CPU{
		pc:          ep,
		v:           make([]uint8, 16),
		mem:         make([]uint8, 4096),
		stack:       NewStack(),
		delay_timer: 0,
		sound_timer: 0,
		Vram:        NewVRAM(),
		Keyboard:    NewKeyboard(),
		time:        0,
	}

	copy(c.mem, Sprites[0:])
	copy(c.mem[ep:], rom)

	return c
}

func (c *CPU) Run() {
	op := c.fetch()
	c.exec(op)
	c.tick()
}

func (c *CPU) tick() {
	c.time += 1
	if c.time < 10 {
		return
	}
	c.time = 0
	if c.delay_timer > 0 {
		c.delay_timer -= 1
	}
	if c.sound_timer > 0 {
		c.sound_timer -= 1
	}
}

func (c *CPU) fetch() uint16 {
	hb := c.mem[c.pc]
	lb := c.mem[c.pc+1]
	op := uint16(hb)<<8 | uint16(lb)
	c.pc += 2
	return op
}

func (c *CPU) exec(op uint16) {
	nnn := op & 0x0fff
	p := (op >> 12) & 0xf
	x := (op >> 8) & 0xf
	y := (op >> 4) & 0xf
	n := op & 0xf
	kk := byte(op & 0xff)

	switch p {
	case 0x0:
		switch kk {
		case 0xe0:
			c.Vram.Clear()
		case 0xee:
			c.pc = c.stack.Pop()
		default:
			c.pc = nnn
			panic("old inst")
		}
	case 0x1:
		c.pc = nnn
	case 0x2:
		c.stack.Push(c.pc)
		c.pc = nnn
	case 0x3:
		if c.v[x] == kk {
			c.pc += 2
		}
	case 0x4:
		if c.v[x] != kk {
			c.pc += 2
		}
	case 0x5:
		if c.v[x] == c.v[y] {
			c.pc += 2
		}
	case 0x6:
		c.v[x] = kk
	case 0x7:
		c.v[x] += kk
	case 0x8:
		switch n {
		case 0x0:
			c.v[x] = c.v[y]
		case 0x1:
			c.v[x] |= c.v[y]
		case 0x2:
			c.v[x] &= c.v[y]
		case 0x3:
			c.v[x] ^= c.v[y]
		case 0x4:
			c.v[x], c.v[0xf] = add(c.v[x], c.v[y])
		case 0x5:
			c.v[x], c.v[0xf] = sub(c.v[x], c.v[y])
		case 0x6:
			c.v[x], c.v[0xf] = shr(c.v[x])
		case 0x7:
			c.v[x], c.v[0xf] = sub(c.v[y], c.v[x])
		case 0xe:
			c.v[x], c.v[0xf] = shl(c.v[x])
		}
	case 0x9:
		if c.v[x] != c.v[y] {
			c.pc += 2
		}
	case 0xa:
		c.i = nnn
	case 0xb:
		c.pc = uint16(c.v[0]) + nnn
	case 0xc:
		c.v[x] = uint8(rand.Int()) & kk
	case 0xd:
		base_x := int(c.v[x])
		base_y := int(c.v[y])

		c.v[0xf] = 0

		for yy := 0; yy < int(n); yy++ {
			row := c.mem[int(c.i)+yy]

			for xx := 0; xx < 8; xx++ {
				vx := base_x + xx
				vy := base_y + yy

				v := (row>>(7-xx))&1 == 1

				if c.Vram.Set(vx, vy, v) {
					c.v[0xf] = 1
				}
			}
		}
	case 0xe:
		switch kk {
		case 0x9e:
			if c.Keyboard.IsPressed(c.v[x]) {
				c.pc += 2
			}
		case 0xa1:
			if !c.Keyboard.IsPressed(c.v[x]) {
				c.pc += 2
			}
		}
	case 0xf:
		switch kk {
		case 0x07:
			c.v[x] = c.delay_timer
		case 0x0a:
			c.v[x] = c.Keyboard.WaitForPress()
		case 0x15:
			c.delay_timer = c.v[x]
		case 0x18:
			c.sound_timer = c.v[x]
		case 0x1e:
			c.i += uint16(c.v[x])
		case 0x29:
			c.i = uint16(c.v[x] * 5)
		case 0x33:
			c.mem[c.i], c.mem[c.i+1], c.mem[c.i+2] = bcd(c.v[x])
		case 0x55:
			copy(c.mem[c.i:c.i+x+1], c.v)
		case 0x65:
			copy(c.v, c.mem[c.i:c.i+x+1])
		}
	}
}

func add(a byte, b byte) (byte, byte) {
	sum := a + b
	if int(a)+int(b) > 255 {
		return sum, 1
	} else {
		return sum, 0
	}
}

func sub(a byte, b byte) (byte, byte) {
	diff := a - b
	if a > b {
		return diff, 1
	} else {
		return diff, 0
	}
}

func shr(a byte) (byte, byte) {
	return a >> 1, a & 1
}

func shl(a byte) (byte, byte) {
	return a << 1, (a >> 7) & 1
}

func bcd(a byte) (byte, byte, byte) {
	return (a / 100) % 10, (a / 10) % 10, a % 10
}
