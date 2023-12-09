package main

var Meter = unit{100}
var Kilometer = unit{100000}
var Centimeter = unit{1}

type unit struct {
	conv float32
}

type Length struct {
	value float32
	unit  unit
}

func (l *Length) Equals(o *Length) bool {
	return toBaseUnit(l) == toBaseUnit(o)
}

func NewLength(value float32, u unit) Length {
	return Length{value: value, unit: u}
}

func (l *Length) Id() float32 {
	return toBaseUnit(l)
}

func (l *Length) Add(l2 *Length) Length {
	result := toBaseUnit(l) + toBaseUnit(l2)
	return fromBaseUnit(result, l.unit)
}

func toBaseUnit(l *Length) float32 {
	return l.value * l.unit.conv
}

func fromBaseUnit(value float32, unit unit) Length {
	return NewLength(value/unit.conv, unit)
}
