package main

const (
	Meter = iota
	Kilometer
	Centimeter
)

type Length struct {
	value int
}

var units = map[int]int{
	Centimeter: 1,
	Meter:      100,
	Kilometer:  100000,
}

func NewLength(value int, unit int) Length {
	value *= units[unit]
	return Length{value: value}
}
