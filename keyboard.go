package main

// Represents 4x4 keypads
type Keyboard struct {
	inputs chan uint8
	keys   []bool
}

// Creates Keyboard instance
func NewKeyboard() Keyboard {
	inputs := make(chan uint8)

	return Keyboard{
		inputs: inputs,
		keys:   make([]bool, 16),
	}
}

// Return `true` if the key is just pressed
func (k *Keyboard) IsPressed(v uint8) bool {
	return k.keys[v]
}

// Waits until a key is pressed
func (k *Keyboard) WaitForPress() uint8 {
	return <-k.inputs
}

// Presses a key. This wakes up the task waiting for a key press if it exists
func (k *Keyboard) Press(v uint8) {
	k.keys[v] = true

	// Try send a signal to wake up a task
	select {
	case k.inputs <- v:
	default:
	}
}

// Releases a key
func (k *Keyboard) Release(v uint8) {
	k.keys[v] = false

	// Clear the signal
	select {
	case _ = <-k.inputs:
	default:
	}
}
