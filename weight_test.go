package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Weight", func() {
	It("should equate 1000 gram to 1 kilogram", func() {
		weight := Gram(1000)
		expected := Kilogram(1)
		Expect(weight.Equals(&expected.Measurement)).To(BeTrue())
	})
	It("should not equate 1 kilogram to 2 kilogram", func() {
		weight := Kilogram(1)
		expected := Kilogram(2)
		Expect(weight.Equals(&expected.Measurement)).To(BeFalse())
	})
	It("should equate 2000000 milligram to 2 kilogram", func() {
		weight := Milligram(2000000)
		expected := Kilogram(2)
		Expect(weight.Equals(&expected.Measurement)).To(BeTrue())
	})
	It("should equate 2 kilogram to 2000000 milligram ", func() {
		weight := Kilogram(2)
		expected := Milligram(2000000)
		Expect(weight.Equals(&expected.Measurement)).To(BeTrue())
	})
	It("should add 1000 grams to 1000 grams and give 2000 grams", func() {
		weight1 := Gram(1000)
		weight2 := Gram(1000)
		expected := Gram(2000)
		Expect(weight1.Add(&weight2)).To(Equal(expected))
	})
	It("should add 1 kilograms to 300 grams and give 1.3 kilograms", func() {
		weight1 := Kilogram(1)
		weight2 := Gram(300)
		expected := Kilogram(1.3)
		Expect(weight1.Add(&weight2)).To(Equal(expected))
	})
	It("should add 500 grams to 1 kilograms and give 1500 grams", func() {
		weight1 := Gram(500)
		weight2 := Kilogram(1)
		expected := Gram(1500)
		Expect(weight1.Add(&weight2)).To(Equal(expected))
	})
})
