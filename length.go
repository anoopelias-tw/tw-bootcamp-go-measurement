package main

var MeterUnit = measurementUnit{100}
var KilometerUnit = measurementUnit{100000}
var CentimeterUnit = measurementUnit{1}

type Length struct {
	Measurement
}

func NewLength(value float32, u unit) Length {
	return Length{NewMeasurement(value, u)}
}

func Meter(value float32) Length {
	return NewLength(value, &MeterUnit)
}

func Kilometer(value float32) Length {
	return NewLength(value, &KilometerUnit)
}

func Centimeter(value float32) Length {
	return NewLength(value, &CentimeterUnit)
}

func (l *Length) Add(l2 *Length) Length {
	result := l.add(&l2.Measurement)
	return NewLength(l.unit.fromBaseUnit(result), l.unit)
}
