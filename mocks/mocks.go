package mocks

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/eggsbenjamin/clinics-microservice-go/constants"
)

const (
	MOCK_CLIENT_ERROR = "MOCK_CLIENT_ERROR"
	MOCK_GET_RESPONSE = `{ "MOCK" : "GET_RESPONSE" }`
	MOCK_OUTWARD_CODE = "MOCK_OUTWARD_CODE"
)

//	Supplementary mocks

type NoopCloser struct {
	io.Reader
}

func (this *NoopCloser) Close() error {
	return nil
}

//	Utils mocks

type MockUtils struct{}

func (this *MockUtils) GetOutwardCode(postcode string) (string, error) {
	return MOCK_OUTWARD_CODE, nil
}

type MockUtilsWithInvalidPostcodeError struct{}

func (this *MockUtilsWithInvalidPostcodeError) GetOutwardCode(postcode string) (string, error) {
	return "", errors.New(constants.POSTCODE_ERROR)
}

//	HTTPClient mocks

type MockHTTPClient struct {
	CalledWith string
}

func (this *MockHTTPClient) Get(url string) (*http.Response, error) {
	this.CalledWith = url

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       &NoopCloser{bytes.NewBufferString(MOCK_GET_RESPONSE)},
	}, nil
}

type MockHTTPClientWithError struct{}

func (this *MockHTTPClientWithError) Get(url string) (*http.Response, error) {
	return nil, errors.New(MOCK_CLIENT_ERROR)
}

type MockHTTPClientWithNon200Response struct{}

func (this *MockHTTPClientWithNon200Response) Get(url string) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       &NoopCloser{bytes.NewBufferString(MOCK_GET_RESPONSE)},
	}, nil
}
