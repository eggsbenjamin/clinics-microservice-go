package clinics_test

import (
	"fmt"
	"github.com/eggsbenjamin/clinics-microservice-go/constants"
	"github.com/eggsbenjamin/clinics-microservice-go/mocks"
	. "github.com/eggsbenjamin/clinics-microservice-go/services"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Clinics", func() {
	It("returns invalid postcode error if postcode is invalid", func() {
		mockClient := &mocks.MockHTTPClient{}
		mockUtilsWithPostcodeError := &mocks.MockUtilsWithInvalidPostcodeError{}

		clinicsService := ClinicsService{
			mockClient,
			mockUtilsWithPostcodeError,
		}

		res, err := clinicsService.GetByPostcode("")

		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal(constants.POSTCODE_ERROR))
		Expect(res).To(BeNil())
	})

	It("calls the http client with the correct url", func() {
		expectedURL := fmt.Sprintf("%s%s", constants.CLINICS_POSTCODE, mocks.MOCK_OUTWARD_CODE)
		mockClient := &mocks.MockHTTPClient{}
		mockUtils := &mocks.MockUtils{}

		clinicsService := ClinicsService{
			Client: mockClient,
			Utils:  mockUtils,
		}

		_, _ = clinicsService.GetByPostcode("")

		Expect(mockClient.CalledWith).To(Equal(expectedURL))
	})

	It("returns client error on client error", func() {
		mockClientWithError := &mocks.MockHTTPClientWithError{}
		mockUtils := &mocks.MockUtils{}

		clinicsService := ClinicsService{
			Client: mockClientWithError,
			Utils:  mockUtils,
		}

		res, err := clinicsService.GetByPostcode("")

		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal(constants.CLIENT_ERROR))
		Expect(res).To(BeNil())

	})

	It("returns service unavailable error on non-200 status from remote service", func() {
		mockClientWithNon200Response := &mocks.MockHTTPClientWithNon200Response{}
		mockUtils := &mocks.MockUtils{}

		clinicsService := ClinicsService{
			Client: mockClientWithNon200Response,
			Utils:  mockUtils,
		}

		res, err := clinicsService.GetByPostcode("")

		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal(constants.SERVICE_UNAVAILABLE_ERROR))
		Expect(res).To(BeNil())
	})

	It("returns a result without error on succes", func() {
		mockClient := &mocks.MockHTTPClient{}
		mockUtils := &mocks.MockUtils{}

		clinicsService := ClinicsService{
			Client: mockClient,
			Utils:  mockUtils,
		}

		res, err := clinicsService.GetByPostcode("")

		Expect(err).To(BeNil())
		Expect(res).ToNot(BeNil())
	})
})
