package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Length", func() {
	It("should equate 1000 meters to 1000 meters", func() {
		length := NewLength(1000, Meter)
		Expect(length).To(Equal(NewLength(1000, Meter)))
	})
	It("should not equate 1 meter to 2 meters", func() {
		length := NewLength(1, Meter)
		Expect(length).To(Not(Equal(NewLength(1000, Meter))))
	})
	It("should equate 1000 meters to 1 kilometer", func() {
		length := NewLength(1000, Meter)
		Expect(length).To(Equal(NewLength(1, Kilometer)))
	})
	It("should equate 1 kilometer to 1000 meters", func() {
		length := NewLength(1, Kilometer)
		Expect(length).To(Equal(NewLength(1000, Meter)))
	})
	It("should be able to put and get values from a map", func() {
		m := make(map[Length]int)
		m[NewLength(1000, Meter)] = 5
		Expect(m[NewLength(1, Kilometer)]).To(Equal(5))
	})
	It("should equate 100 centimeter to 1 meter", func() {
		length := NewLength(100, Centimeter)
		Expect(length).To(Equal(NewLength(1, Meter)))
	})
	It("should equate 10 meter to 1000 centimeter", func() {
		length := NewLength(10, Meter)
		Expect(length).To(Equal(NewLength(1000, Centimeter)))
	})
	It("should equate 5 kilometer to 500000 centimeter", func() {
		length := NewLength(5, Kilometer)
		Expect(length).To(Equal(NewLength(500000, Centimeter)))
	})
	It("should equate 200000 centimeter to 2 kilometer", func() {
		length := NewLength(200000, Centimeter)
		Expect(length).To(Equal(NewLength(2, Kilometer)))
	})
})
