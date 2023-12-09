package main

type measurementUnit struct {
	conv float32
}

func (u *measurementUnit) toBaseUnit(val float32) float32 {
	return val * u.conv
}

func (u *measurementUnit) fromBaseUnit(val float32) float32 {
	return val / u.conv
}

type Measurement struct {
	value float32
	unit  unit
}

func (m *Measurement) Equals(o *Measurement) bool {
	return m.toBaseUnit() == o.toBaseUnit()
}

func NewMeasurement(value float32, u unit) Measurement {
	return Measurement{value: value, unit: u}
}

func (m *Measurement) toBaseUnit() float32 {
	return m.unit.toBaseUnit(m.value)
}

func (m *Measurement) Id() float32 {
	return m.toBaseUnit()
}

func (m *Measurement) add(m2 *Measurement) float32 {
	return m.toBaseUnit() + m2.toBaseUnit()
}
