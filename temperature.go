package main

type temperatureUnit struct {
	conv   float32
	offset float32
}

func (u *temperatureUnit) toBaseUnit(val float32) float32 {
	return (val - u.offset) * u.conv
}

func (u *temperatureUnit) fromBaseUnit(val float32) float32 {
	return (val / u.conv) + u.offset
}

var CelsiusUnit = temperatureUnit{1, 0}
var FahrenheitUnit = temperatureUnit{5.0 / 9.0, 32}
var KelvinUnit = temperatureUnit{1, 273.15}

type Temperature struct {
	Measurement
}

func NewTemperature(value float32, u unit) Temperature {
	return Temperature{NewMeasurement(value, u)}
}

func Celsius(value float32) Temperature {
	return NewTemperature(value, &CelsiusUnit)
}

func Fahrenheit(value float32) Temperature {
	return NewTemperature(value, &FahrenheitUnit)
}

func Kelvin(value float32) Temperature {
	return NewTemperature(value, &KelvinUnit)
}
