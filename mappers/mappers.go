package mappers

import (
	"fmt"
	"strings"

	"github.com/eggsbenjamin/clinics-microservice-go/models"
)

type Mapper struct{}

func (this *Mapper) formatAddressFields(input *models.PartialPostcodeResult) string {
	var (
		populatedFields []string
		result          = ""
	)

	fields := []string{
		input.Address1,
		input.Address2,
		input.Address3,
		input.Postcode,
		input.City,
	}

	for i := 0; i < len(fields); i++ {
		if fields[i] != "" {
			populatedFields = append(populatedFields, fields[i])
		}
	}

	result += fmt.Sprintf("%s ", input.Name)

	if len(populatedFields) > 0 {
		result += fmt.Sprintf("(%s)", strings.Join(populatedFields, ", "))
	}

	return strings.TrimSpace(result)
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

func (this *Mapper) MapPartialPostcodeResponse(input *models.PartialPostcodeResponse) *models.PartialPostcodeClientResponse {
	output := &models.PartialPostcodeClientResponse{
		Success: input.Success,
		Results: []models.PartialPostcodeClientResult{},
	}

	for _, result := range input.Results {
		output.Results = append(output.Results, *this.MapPartialPostcodeResult(&result))
	}

	return output
}
