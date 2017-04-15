package mocks

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

const (
	MockInvalidPostcodeError = "MOCK_INVALID_POSTCODE_ERROR"
	MockClientError          = "MOCK_CLIENT_ERROR"
	MockGetResponse          = `{ "MOCK" : "GET_RESPONSE" }`
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
	return "", nil
}

type MockUtilsWithInvalidPostcodeError struct{}

func (this *MockUtilsWithInvalidPostcodeError) GetOutwardCode(postcode string) (string, error) {
	return "", errors.New(MockInvalidPostcodeError)
}

//	HTTPClient mocks

type MockHTTPClient struct{}

func (this *MockHTTPClient) Get(url string) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       &NoopCloser{bytes.NewBufferString(MockGetResponse)},
	}, nil
}

type MockHTTPClientWithError struct{}

func (this *MockHTTPClientWithError) Get(url string) (*http.Response, error) {
	return nil, errors.New(MockClientError)
}

type MockHTTPClientWithNon200Response struct{}

func (this *MockHTTPClientWithNon200Response) Get(url string) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       &NoopCloser{bytes.NewBufferString(MockGetResponse)},
	}, nil
}
