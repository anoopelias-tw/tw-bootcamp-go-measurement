package main

var Meter = unit{100}
var Kilometer = unit{100000}
var Centimeter = unit{1}

type unit struct {
	conv int
}

type Length struct {
	value int
}

func (l *Length) Equals(o *Length) bool {
	return l.value == o.value
}

func NewLength(value int, unit unit) Length {
	value *= unit.conv
	return Length{value: value}
}

func (l *Length) Id() int {
	return l.value
}
