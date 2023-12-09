package main

var MeterUnit = unit{100}
var KilometerUnit = unit{100000}
var CentimeterUnit = unit{1}

type unit struct {
	conv float32
}

func (u *unit) toBaseUnit(val float32) float32 {
	return val * u.conv
}

func (u *unit) fromBaseUnit(val float32) float32 {
	return val / u.conv
}

type Length struct {
	value float32
	unit  unit
}

func (l *Length) Equals(o *Length) bool {
	return l.toBaseUnit() == o.toBaseUnit()
}

func NewLength(value float32, u unit) Length {
	return Length{value: value, unit: u}
}

func Meter(value float32) Length {
	return NewLength(value, MeterUnit)
}

func Kilometer(value float32) Length {
	return NewLength(value, KilometerUnit)
}

func Centimeter(value float32) Length {
	return NewLength(value, CentimeterUnit)
}

func (l *Length) toBaseUnit() float32 {
	return l.unit.toBaseUnit(l.value)
}

func (l *Length) Id() float32 {
	return l.toBaseUnit()
}

func (l *Length) Add(l2 *Length) Length {
	result := l.toBaseUnit() + l2.toBaseUnit()
	return NewLength(l.unit.fromBaseUnit(result), l.unit)
}
