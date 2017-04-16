package clinics_test

import (
	"github.com/eggsbenjamin/clinics-microservice-go/mocks"
	. "github.com/eggsbenjamin/clinics-microservice-go/services"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Clinics", func() {
	It("returns 'Invalid Postcode' if postcode is invalid", func() {
		mockClient := &mocks.MockHTTPClient{}
		mockUtilsWithPostcodeError := &mocks.MockUtilsWithInvalidPostcodeError{}

		clinicsService := ClinicsService{
			mockClient,
			mockUtilsWithPostcodeError,
		}

		res, err := clinicsService.GetByPostcode("")

		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal(mocks.MockInvalidPostcodeError))
		Expect(res).To(BeNil())
	})
})
