package main

import (
	"testing"
)

func TestCpuAdd(t *testing.T) {
	v, vf := add(1, 2)
	if v != 3 {
		t.Fatalf(`1 + 2 should be 3`)
	}
	if vf != 0 {
		t.Fatalf(`1 + 2 shouldn't overflow`)
	}
}

func TestCpuAddMax(t *testing.T) {
	v, vf := add(1, 254)
	if v != 255 {
		t.Fatalf(`1 + 254 should be 255`)
	}
	if vf != 0 {
		t.Fatalf(`1 + 254 shouldn't overflow`)
	}
}

func TestCpuAddOverflow(t *testing.T) {
	v, vf := add(1, 255)
	if v != 0 {
		t.Fatalf(`1 + 255 should be 0`)
	}
	if vf != 1 {
		t.Fatalf(`1 + 255 should overflow`)
	}
}

func TestCpuAddOverflowZero(t *testing.T) {
	v, vf := add(1, 255)
	if v != 0 {
		t.Fatalf(`1 + 255 should be 0`)
	}
	if vf != 1 {
		t.Fatalf(`1 + 255 should overflow`)
	}
}

func TestCpuSub(t *testing.T) {
	v, vf := sub(3, 1)
	if v != 2 {
		t.Fatalf(`3 - 1 should be 2`)
	}
	if vf != 1 {
		t.Fatalf(`3 - 1 shouldn't borrow`)
	}
}

func TestCpuSubBorrow(t *testing.T) {
	v, vf := sub(1, 3)
	if v != 254 {
		t.Fatalf(`1 - 3 should be 254`)
	}
	if vf != 0 {
		t.Fatalf(`1 - 3 should borrow`)
	}
}

func TestCpuSubBorrowZero(t *testing.T) {
	v, vf := sub(1, 1)
	if v != 0 {
		t.Fatalf(`1 - 1 should be 2`)
	}
	if vf != 0 {
		t.Fatalf(`1 - 1 should borrow (as spec)`)
	}
}

func TestShr(t *testing.T) {
	v, vf := shr(12)
	if v != 6 {
		t.Fatalf(`12 >> 1 should be 6`)
	}
	if vf != 0 {
		t.Fatalf(`12 >> 1 shouldn't loose bit`)
	}
}

func TestShrLooseBit(t *testing.T) {
	v, vf := shr(1)
	if v != 0 {
		t.Fatalf(`1 >> 1 should be 0`)
	}
	if vf != 1 {
		t.Fatalf(`1 >> 1 should loose bit`)
	}
}

func TestShl(t *testing.T) {
	v, vf := shl(1)
	if v != 2 {
		t.Fatalf(`1 << 1 should be 2`)
	}
	if vf != 0 {
		t.Fatalf(`1 << 1 shouldn't loose bit`)
	}
}

func TestShlLooseBit(t *testing.T) {
	v, vf := shl(128)
	if v != 0 {
		t.Fatalf(`128 << 1 should be 0`)
	}
	if vf != 1 {
		t.Fatalf(`128 << 1 should loose bit`)
	}
}

func TestBcd(t *testing.T) {
	a, b, c := bcd(123)

	if a != 1 {
		t.Fatalf(`hundreds digit is wrong`)
	}
	if b != 2 {
		t.Fatalf(`tens digit is wrong`)
	}
	if c != 3 {
		t.Fatalf(`ones digit is wrong`)
	}
}
