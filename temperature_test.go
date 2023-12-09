package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Temperature", func() {
	It("should equate 0 degree celsius to 32 fahrenheit", func() {
		temperature := Celsius(0)
		expected := Fahrenheit(32)
		Expect(temperature.Equals(&expected.Measurement)).To(BeTrue())
	})
	It("should not equate 0 degree celsius to 1 fahrenheit", func() {
		temperature := Celsius(0)
		expected := Fahrenheit(1)
		Expect(temperature.Equals(&expected.Measurement)).To(BeFalse())
	})
	It("should equate 0 degree celsius to 273.15 kelvin", func() {
		temperature := Celsius(0)
		expected := Kelvin(273.15)
		Expect(temperature.Equals(&expected.Measurement)).To(BeTrue())
	})

})
