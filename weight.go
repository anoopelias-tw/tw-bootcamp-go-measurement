package main

var GramUnit = measurementUnit{1000}
var KilogramUnit = measurementUnit{1000000}
var MilligramUnit = measurementUnit{1}

type Weight struct {
	Measurement
}

func (w Weight) Add(w2 *Weight) Weight {
	result := w.add(&w2.Measurement)
	return NewWeight(w.unit.fromBaseUnit(result), w.unit)
}

func NewWeight(value float32, u unit) Weight {
	return Weight{NewMeasurement(value, u)}
}

func Gram(value float32) Weight {
	return NewWeight(value, &GramUnit)
}

func Kilogram(value float32) Weight {
	return NewWeight(value, &KilogramUnit)
}

func Milligram(value float32) Weight {
	return NewWeight(value, &MilligramUnit)
}
