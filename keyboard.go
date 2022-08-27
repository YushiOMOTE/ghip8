package main

type Keyboard struct {
	inputs chan uint8
	keys   []bool
}

func NewKeyboard() Keyboard {
	inputs := make(chan uint8)

	return Keyboard{
		inputs: inputs,
		keys:   make([]bool, 16),
	}
}

func (k *Keyboard) IsPressed(v uint8) bool {
	return k.keys[v]
}

func (k *Keyboard) WaitForPress() uint8 {
	return <-k.inputs
}

func (k *Keyboard) Press(v uint8) {
	k.keys[v] = true
	select {
	case k.inputs <- v:
	default:
	}
}

func (k *Keyboard) Release(v uint8) {
	k.keys[v] = false
	select {
	case _ = <- k.inputs:
	default:
	}
}
