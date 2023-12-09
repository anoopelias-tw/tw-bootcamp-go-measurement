package main

type unit interface {
	toBaseUnit(val float32) float32
	fromBaseUnit(val float32) float32
}
