package utils_test

import (
	"github.com/eggsbenjamin/clinics-microservice-go/constants"
	. "github.com/eggsbenjamin/clinics-microservice-go/utils"
	testUtils "github.com/eggsbenjamin/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Fixture struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

var _ = Describe("Utils", func() {
	It("returns the correct error if the postcode contains less than 5 non-whitespace characters", func() {
		utils := &Utils{}
		input := "w"
		actual, err := utils.GetOutwardCode(input)

		Expect(err.Error()).To(Equal(constants.POSTCODE_ERROR))
		Expect(actual).To(Equal(""))
	})

	It("returns the correct error if the postcode contains greater than 7 non-whitespace characters", func() {
		utils := &Utils{}
		input := "wqqqqqqq"
		actual, err := utils.GetOutwardCode(input)

		Expect(err.Error()).To(Equal(constants.POSTCODE_ERROR))
		Expect(actual).To(Equal(""))
	})

	It("returns the correct error if the postcode contains any non-alphanumeric chars", func() {
		utils := &Utils{}
		input := "W6 4$D"
		actual, err := utils.GetOutwardCode(input)

		Expect(err.Error()).To(Equal(constants.POSTCODE_ERROR))
		Expect(actual).To(Equal(""))
	})

	It("returns the correct outward code on valid postcode argument", func() {
		utils := &Utils{}
		testData := []Fixture{}

		testUtils.UnmarshalJsonFile("../fixtures/partial_postcode.json", &testData)

		for _, fixture := range testData {
			actual, err := utils.GetOutwardCode(fixture.Input)
			expected := fixture.Output

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		}
	})
})
