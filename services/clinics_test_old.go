package clinics

import (
	"github.com/eggsbenjamin/clinics-microservice-go/mocks"
	"testing"
)

func TestInvalidPostcodeErrorHandling(t *testing.T) {
	mockClient := &mocks.MockHTTPClient{}
	mockUtilsWithPostcodeError := &mocks.MockUtilsWithInvalidPostcodeError{}

	clinicsService := ClinicsService{
		Client: mockClient,
		Utils:  mockUtilsWithPostcodeError,
	}

	res, err := clinicsService.GetByPostcode("")

	if err == nil || string(err.Error()) != mocks.MockInvalidPostcodeError {
		t.Logf("expected error not returned : %v", err)
		t.Fail()
	}

	if res != nil {
		t.Logf("unexpected response : %v", res)
		t.Fail()
	}
}

func TestClientErrorHandling(t *testing.T) {
	mockClientWithError := &mocks.MockHTTPClientWithError{}
	mockUtils := &mocks.MockUtils{}

	clinicsService := ClinicsService{
		Client: mockClientWithError,
		Utils:  mockUtils,
	}

	res, err := clinicsService.GetByPostcode("")

	if err == nil || string(err.Error()) != mocks.MockClientError {
		t.Logf("expected error not returned : %v", err)
		t.Fail()
	}

	if res != nil {
		t.Logf("unexpected response : %v", res)
		t.Fail()
	}
}

func TestClientNon200ResponseHandling(t *testing.T) {
	mockClientWithNon200Response := &mocks.MockHTTPClientWithNon200Response{}
	mockUtils := &mocks.MockUtils{}

	clinicsService := ClinicsService{
		Client: mockClientWithNon200Response,
		Utils:  mockUtils,
	}

	res, err := clinicsService.GetByPostcode("")

	if err == nil || string(err.Error()) != serviceUnavailableError {
		t.Logf("expected error not returned : %v", err)
		t.Fail()
	}

	if res != nil {
		t.Logf("unexpected response : %v", res)
		t.Fail()
	}
}

func TestHappyPath(t *testing.T) {
	mockClient := &mocks.MockHTTPClient{}
	mockUtils := &mocks.MockUtils{}

	clinicsService := ClinicsService{
		Client: mockClient,
		Utils:  mockUtils,
	}

	res, err := clinicsService.GetByPostcode("")

	if err != nil {
		t.Logf("unexpected error returned : %v", err)
		t.Fail()
	}

	if res == nil {
		t.Logf("expected response not returned : %v", res)
		t.Fail()
	}
}