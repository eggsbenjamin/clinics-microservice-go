package clinics

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/eggsbenjamin/clinics-microservice-go/models"
	"github.com/eggsbenjamin/clinics-microservice-go/utils"
)

const (
	serviceUnavailableError = "SERVICE UNAVAILABLE"
)

type IHTTPClient interface {
	Get(string) (*http.Response, error)
}

type ClinicsService struct {
	Client IHTTPClient
	Utils  utils.IUtils
}

func (this *ClinicsService) GetByPostcode(postcode string) (*models.PartialPostcodeClientResponse, error) {
	_, err := this.Utils.GetOutwardCode(postcode)

	if err != nil {
		return nil, err
	}

	res, err := this.Client.Get("")

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(serviceUnavailableError)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	result := &models.PartialPostcodeClientResponse{}

	json.Unmarshal(body, result)

	return result, nil
}
