package main

// Represents no key input
const NotFound = uint8(255)

// Represents 4x4 keypads
type Keyboard struct {
	keys []bool
}

// Creates Keyboard instance
func NewKeyboard() Keyboard {
	return Keyboard{make([]bool, 16)}
}

// Return `true` if the key is just pressed
func (k *Keyboard) IsPressed(v uint8) bool {
	return k.keys[v]
}

// Return `true` if the key is just pressed
func (k *Keyboard) GetPressed() uint8 {
	for i, key := range k.keys {
		if key {
			return uint8(i)
		}
	}
	return NotFound
}

// Presses a key. This wakes up the task waiting for a key press if it exists
func (k *Keyboard) Press(v uint8) {
	k.keys[v] = true
}

// Releases a key
func (k *Keyboard) Release(v uint8) {
	k.keys[v] = false
}
