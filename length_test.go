package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Length", func() {
	It("should equate 1000 meters to 1000 meters", func() {
		length := NewLength(1000, Meter)
		expected := NewLength(1000, Meter)
		Expect(length.Equals(&expected)).To(BeTrue())
	})
	It("should not equate 1 meter to 2 meters", func() {
		length := NewLength(1, Meter)
		expected := NewLength(2, Meter)
		Expect(length.Equals(&expected)).To(BeFalse())
	})
	It("should equate 1000 meters to 1 kilometer", func() {
		length := NewLength(1000, Meter)
		expected := NewLength(1, Kilometer)
		Expect(length.Equals(&expected)).To(BeTrue())
	})
	It("should equate 1 kilometer to 1000 meters", func() {
		length := NewLength(1, Kilometer)
		expected := NewLength(1000, Meter)
		Expect(length.Equals(&expected)).To(BeTrue())
	})
	It("should be able to put and get values from a map", func() {
		m := make(map[int]int)
		length := NewLength(1000, Meter)
		m[length.Id()] = 5
		expected := NewLength(1, Kilometer)
		Expect(m[expected.Id()]).To(Equal(5))
	})
	It("should equate 100 centimeter to 1 meter", func() {
		length := NewLength(100, Centimeter)
		expected := NewLength(1, Meter)
		Expect(length.Equals(&expected)).To(BeTrue())
	})
	It("should equate 10 meter to 1000 centimeter", func() {
		length := NewLength(10, Meter)
		expected := NewLength(1000, Centimeter)
		Expect(length.Equals(&expected)).To(BeTrue())
	})
	It("should equate 5 kilometer to 500000 centimeter", func() {
		length := NewLength(5, Kilometer)
		expected := NewLength(500000, Centimeter)
		Expect(length.Equals(&expected)).To(BeTrue())
	})
	It("should equate 200000 centimeter to 2 kilometer", func() {
		length := NewLength(200000, Centimeter)
		expected := NewLength(2, Kilometer)
		Expect(length.Equals(&expected)).To(BeTrue())
	})
})
