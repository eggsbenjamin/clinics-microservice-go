package mappers

import (
	"fmt"

	"github.com/eggsbenjamin/clinics-microservice-go/models"
)

type Mapper struct{}

func (this *Mapper) formatAddressFields(input *models.PartialPostcodeResult) string {
	var (
		formattedAddress string
		result           = ""
	)

	fieldsToBeFormatted := []string{
		input.Address1,
		input.Address2,
		input.Address3,
		input.Postcode,
		input.City,
	}

	for i := 0; i < len(fieldsToBeFormatted); i++ {
		if fieldsToBeFormatted[i] != "" {
			formattedAddress += fieldsToBeFormatted[i]
		}

		if i != len(fieldsToBeFormatted)-1 {
			formattedAddress += ", "
		}
	}

	if input.Name != "" {
		result += input.Name
	}

	if formattedAddress != "" {
		result += fmt.Sprintf(" (%s)", formattedAddress)
	}

	return result
}

func (this *Mapper) MapPartialPostcodeResult(input *models.PartialPostcodeResult) *models.PartialPostcodeClientResult {
	outputCrsProperties := models.PartialPostcodeClientResultCrsProperties{
		Name: input.LatLong.Crs.Properties.Name,
	}

	outputCrs := models.PartialPostcodeClientResultCrs{
		Type:       input.LatLong.Crs.Type,
		Properties: outputCrsProperties,
	}

	outputLatLong := models.PartialPostcodeClientResultLatLong{
		Type:        input.LatLong.Type,
		Crs:         outputCrs,
		Coordinates: input.LatLong.Coordinates,
	}

	output := &models.PartialPostcodeClientResult{
		Website:            input.Website,
		SubType:            input.SubType,
		Postcode:           input.Postcode,
		Phone:              input.Phone,
		PartialPostCode:    input.PartialPostCode,
		OrganisationType:   input.OrganisationType,
		OrganisationStatus: input.OrganisationStatus,
		OrganisationId:     input.OrganisationId,
		OrganisationCode:   input.OrganisationCode,
		Name:               input.Name,
		Longitude:          input.Longitude,
		Latitude:           input.Latitude,
		IsPimsManaged:      input.IsPimsManaged,
		Fax:                input.Fax,
		Email:              input.Email,
		County:             input.County,
		City:               input.City,
		Address1:           input.Address1,
		Address2:           input.Address2,
		Address3:           input.Address3,
		LatLong:            outputLatLong,
		Formatted:          this.formatAddressFields(input),
	}

	return output
}
