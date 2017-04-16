package clinics

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eggsbenjamin/clinics-microservice-go/constants"
	"github.com/eggsbenjamin/clinics-microservice-go/models"
	"github.com/eggsbenjamin/clinics-microservice-go/utils"
)

type IHTTPClient interface {
	Get(string) (*http.Response, error)
}

type ClinicsService struct {
	Client IHTTPClient
	Utils  utils.IUtils
}

func (this *ClinicsService) GetByPostcode(postcode string) (*models.PartialPostcodeClientResponse, error) {
	outwardCode, err := this.Utils.GetOutwardCode(postcode)

	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s", constants.CLINICS_POSTCODE, outwardCode)

	res, err := this.Client.Get(url)

	if err != nil {
		return nil, errors.New(constants.CLIENT_ERROR)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(constants.SERVICE_UNAVAILABLE_ERROR)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	result := &models.PartialPostcodeClientResponse{}

	err = json.Unmarshal(body, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
