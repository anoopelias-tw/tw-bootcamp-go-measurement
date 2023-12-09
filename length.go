package main

const (
	Meter = iota
	Kilometer
)

type Length struct {
	value int
}

func NewLength(value int, unit int) Length {
	if unit == Kilometer {
		value *= 1000
	}
	return Length{value: value}
}
