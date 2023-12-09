package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Length", func() {
	It("should equate 1000 meters to 1000 meters", func() {
		length := Meter(1000)
		expected := Meter(1000)
		Expect(length.Equals(&expected.Measurement)).To(BeTrue())
	})
	It("should not equate 1 meter to 2 meters", func() {
		length := Meter(1)
		expected := Meter(2)
		Expect(length.Equals(&expected.Measurement)).To(BeFalse())
	})
	It("should equate 1000 meters to 1 kilometer", func() {
		length := Meter(1000)
		expected := Kilometer(1)
		Expect(length.Equals(&expected.Measurement)).To(BeTrue())
	})
	It("should equate 1 kilometer to 1000 meters", func() {
		length := Kilometer(1)
		expected := Meter(1000)
		Expect(length.Equals(&expected.Measurement)).To(BeTrue())
	})
	It("should be able to put and get values from a map", func() {
		m := make(map[float32]int)
		length := Meter(1000)
		m[length.Id()] = 5
		expected := Kilometer(1)
		Expect(m[expected.Id()]).To(Equal(5))
	})
	It("should equate 100 centimeter to 1 meter", func() {
		length := Centimeter(100)
		expected := Meter(1)
		Expect(length.Equals(&expected.Measurement)).To(BeTrue())
	})
	It("should equate 10 meter to 1000 centimeter", func() {
		length := Meter(10)
		expected := Centimeter(1000)
		Expect(length.Equals(&expected.Measurement)).To(BeTrue())
	})
	It("should equate 5 kilometer to 500000 centimeter", func() {
		length := Kilometer(5)
		expected := Centimeter(500000)
		Expect(length.Equals(&expected.Measurement)).To(BeTrue())
	})
	It("should equate 200000 centimeter to 2 kilometer", func() {
		length := Centimeter(200000)
		expected := Kilometer(2)
		Expect(length.Equals(&expected.Measurement)).To(BeTrue())
	})
	It("should add 1000 meter to 1000 meter and give 2000 meters", func() {
		length1 := Meter(1000)
		length2 := Meter(1000)
		expected := Meter(2000)
		Expect(length1.Add(&length2)).To(Equal(expected))
	})
	It("should add 1 kilometer to 300 meters and give 1.3 kilometer", func() {
		length1 := Kilometer(1)
		length2 := Meter(300)
		expected := Kilometer(1.3)
		Expect(length1.Add(&length2)).To(Equal(expected))
	})
	It("should add 500 meter to 1 kilometer and give 1500 meter", func() {
		length1 := Meter(500)
		length2 := Kilometer(1)
		expected := Meter(1500)
		Expect(length1.Add(&length2)).To(Equal(expected))
	})
})
