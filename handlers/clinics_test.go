package handlers_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/eggsbenjamin/clinics-microservice-go/constants"
	"github.com/eggsbenjamin/clinics-microservice-go/fixtures"
	. "github.com/eggsbenjamin/clinics-microservice-go/handlers"
	"github.com/eggsbenjamin/clinics-microservice-go/mocks"
	"github.com/eggsbenjamin/clinics-microservice-go/models"
	"github.com/julienschmidt/httprouter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	resp *httptest.ResponseRecorder
)

var (
	MOCK_POSTCODE = "TW208DE"
	POSTCODE_URL  = fmt.Sprintf("/clinics/postcode/%s", MOCK_POSTCODE)
	postcodeParam = httprouter.Param{
		Key:   "postcode",
		Value: MOCK_POSTCODE,
	}
)

var _ = Describe("Clinics Handlers", func() {

	var _ = BeforeEach(func() {
		resp = httptest.NewRecorder()
	})

	var _ = Describe("Clinics By Postcode", func() {

		Context("When an invalid postcode is received", func() {
			It("should return a 400 with an 'Invalid Postcode' error", func() {
				mockClinicsService := &mocks.MockClinicsService{}
				mockClinicsService.On("GetByPostcode", MOCK_POSTCODE).
					Return(nil, errors.New(constants.POSTCODE_ERROR))

				clinicsHandlers := NewClinicsHandlers(mockClinicsService, nil)

				req := httptest.NewRequest("GET", POSTCODE_URL, nil)
				params := httprouter.Params{postcodeParam}

				clinicsHandlers.ClinicsByPostcode(resp, req, params)
				result := resp.Result()
				body, _ := ioutil.ReadAll(result.Body)

				defer result.Body.Close()

				Expect(result.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(body).To(Equal([]byte(fmt.Sprintf(`{ "error" : "%s" }`, constants.POSTCODE_ERROR))))
			})
		})

		Context("When a client error occurs when retreiving the clinics", func() {
			It("should return a 500 with an 'Internal Server Error' error", func() {
				mockClinicsService := &mocks.MockClinicsService{}
				mockClinicsService.On("GetByPostcode", MOCK_POSTCODE).
					Return(nil, errors.New(constants.CLIENT_ERROR))

				clinicsHandlers := NewClinicsHandlers(mockClinicsService, nil)

				req := httptest.NewRequest("GET", POSTCODE_URL, nil)
				params := httprouter.Params{postcodeParam}

				clinicsHandlers.ClinicsByPostcode(resp, req, params)
				result := resp.Result()
				body, _ := ioutil.ReadAll(result.Body)

				defer result.Body.Close()

				Expect(result.StatusCode).To(Equal(http.StatusInternalServerError))
				Expect(body).To(Equal([]byte(fmt.Sprintf(`{ "error" : "%s" }`, constants.INTERNAL_SERVER_ERROR))))
			})
		})

		Context("When an unexpected response code is returned from the clinics data source", func() {
			It("should return a 500 with an 'Internal Server Error' error", func() {
				mockClinicsService := &mocks.MockClinicsService{}
				mockClinicsService.On("GetByPostcode", MOCK_POSTCODE).
					Return(nil, errors.New(constants.SERVICE_UNAVAILABLE_ERROR))

				clinicsHandlers := NewClinicsHandlers(mockClinicsService, nil)

				req := httptest.NewRequest("GET", POSTCODE_URL, nil)
				params := httprouter.Params{postcodeParam}

				clinicsHandlers.ClinicsByPostcode(resp, req, params)
				result := resp.Result()
				body, _ := ioutil.ReadAll(result.Body)

				defer result.Body.Close()

				Expect(result.StatusCode).To(Equal(http.StatusInternalServerError))
				Expect(body).To(Equal([]byte(fmt.Sprintf(`{ "error" : "%s" }`, constants.INTERNAL_SERVER_ERROR))))
			})
		})

		Context("When an expected response code is returned from the clinics data source", func() {
			It("should return a 200 with the correct payload", func() {
				mockPartialPostcodeResponse := &models.PartialPostcodeResponse{}
				mockPartialPostcodeClientResponse := &models.PartialPostcodeClientResponse{}

				json.Unmarshal([]byte(fixtures.EXPECTED_FORMATTED_POSTCODE_RESPONSE), mockPartialPostcodeClientResponse)

				mockClinicsService := &mocks.MockClinicsService{}
				mockClinicsService.On("GetByPostcode", MOCK_POSTCODE).
					Return(mockPartialPostcodeResponse, nil)

				mockMapper := &mocks.MockMapper{}
				mockMapper.On("MapPartialPostcodeResponse", mockPartialPostcodeResponse).
					Return(mockPartialPostcodeClientResponse, nil)

				clinicsHandlers := NewClinicsHandlers(mockClinicsService, mockMapper)

				req := httptest.NewRequest("GET", POSTCODE_URL, nil)
				params := httprouter.Params{postcodeParam}

				clinicsHandlers.ClinicsByPostcode(resp, req, params)
				result := resp.Result()
				body, _ := ioutil.ReadAll(result.Body)

				defer result.Body.Close()

				Expect(result.StatusCode).To(Equal(http.StatusOK))
				Expect(body).To(MatchJSON(fixtures.EXPECTED_FORMATTED_POSTCODE_RESPONSE))
			})
		})
	})

})
