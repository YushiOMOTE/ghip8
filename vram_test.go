package main

import (
	"testing"
)

func TestVRAM(t *testing.T) {
	vram := NewVRAM()

	// false ^ true -> true
	if vram.Set(0, 0, true) {
		t.Fatalf(`Unexpected erase on false ^ true`)
	}
	if !vram.Get(0, 0) {
		t.Fatalf(`VRAM cannot be set on false ^ true`)
	}

	// true ^ false -> true
	if vram.Set(0, 0, false) {
		t.Fatalf(`Unexpected erase on true ^ false`)
	}
	if !vram.Get(0, 0) {
		t.Fatalf(`VRAM cannot be set on true ^ false`)
	}

	// true ^ true -> false
	if !vram.Set(0, 0, true) {
		t.Fatalf(`VRAM cannot be erased on true ^ true`)
	}
	if vram.Get(0, 0) {
		t.Fatalf(`VRAM cannot be reset on true ^ true`)
	}

	// false ^ false -> false
	if vram.Set(0, 0, false) {
		t.Fatalf(`Unexpected erase on false ^ false`)
	}
	if vram.Get(0, 0) {
		t.Fatalf(`VRAM cannot be reset on false ^ false`)
	}
}
