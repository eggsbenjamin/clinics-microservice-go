package mappers_test

import (
	"fmt"

	. "github.com/eggsbenjamin/clinics-microservice-go/mappers"

	"github.com/eggsbenjamin/clinics-microservice-go/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	MOCK_NAME                = "MOCK_NAME"
	MOCK_ADDRESS1            = "MOCK_ADDRESS1"
	MOCK_ADDRESS2            = "MOCK_ADDRESS2"
	MOCK_ADDRESS3            = "MOCK_ADDRESS3"
	MOCK_POSTCODE            = "MOCK_POSTCODE"
	MOCK_CITY                = "MOCK_CITY"
	MOCK_WEBSITE             = "MOCK_WEBSITE"
	MOCK_SUBTYPE             = "MOCK_SUBTYPE"
	MOCK_PHONE               = "MOCK_PHONE"
	MOCK_PARTIAL_POSTCODE    = "MOCK_PARTIAL_POSTCODE"
	MOCK_ORGANISATION_TYPE   = "MOCK_ORGANISATION_TYPE"
	MOCK_ORGANISATION_ID     = "MOCK_ORGANISATION_ID"
	MOCK_ORGANISATION_STATUS = "MOCK_ORGANISATION_STATUS"
	MOCK_ORGANISATION_CODE   = "MOCK_ORGANISATION_CODE"
	MOCK_LATITUDE            = "MOCK_LATITUDE"
	MOCK_LONGITUDE           = "MOCK_LONGITUDE"
	MOCK_IS_PIMS_MANAGED     = "MOCK_IS_PIMS_MANAGED"
	MOCK_FAX                 = "MOCK_FAX"
	MOCK_EMAIL               = "MOCK_EMAIL"
	MOCK_COUNTY              = "MOCK_COUNTY"
)

var (
	MOCK_PARTIAL_POSTCODE_RESULT_LAT_LONG              *models.PartialPostcodeResultLatLong
	MOCK_PARTIAL_POSTCODE_RESULT_CRS                   *models.PartialPostcodeResultCrs
	MOCK_PARTIAL_POSTCODE_RESULT_CRS_PROPERTIES        *models.PartialPostcodeResultCrsProperties
	MOCK_PARTIAL_POSTCODE_CLIENT_RESULT_LAT_LONG       models.PartialPostcodeClientResultLatLong
	MOCK_PARTIAL_POSTCODE_CLIENT_RESULT_CRS            models.PartialPostcodeClientResultCrs
	MOCK_PARTIAL_POSTCODE_CLIENT_RESULT_CRS_PROPERTIES models.PartialPostcodeClientResultCrsProperties
)

var _ = BeforeEach(func() {
	MOCK_PARTIAL_POSTCODE_RESULT_CRS_PROPERTIES = &models.PartialPostcodeResultCrsProperties{
		Name: "MOCK_CRS_PROPERTIES_NAME",
	}

	MOCK_PARTIAL_POSTCODE_RESULT_CRS = &models.PartialPostcodeResultCrs{
		Type:       "MOCK_CRS_TYPE",
		Properties: MOCK_PARTIAL_POSTCODE_RESULT_CRS_PROPERTIES,
	}

	MOCK_PARTIAL_POSTCODE_RESULT_LAT_LONG = &models.PartialPostcodeResultLatLong{
		Type:        "MOCK_LAT_LONG_TYPE",
		Crs:         MOCK_PARTIAL_POSTCODE_RESULT_CRS,
		Coordinates: []float64{1, 2},
	}

	MOCK_PARTIAL_POSTCODE_CLIENT_RESULT_CRS_PROPERTIES = models.PartialPostcodeClientResultCrsProperties{
		Name: "MOCK_CRS_PROPERTIES_NAME",
	}

	MOCK_PARTIAL_POSTCODE_CLIENT_RESULT_CRS = models.PartialPostcodeClientResultCrs{
		Type:       "MOCK_CRS_TYPE",
		Properties: MOCK_PARTIAL_POSTCODE_CLIENT_RESULT_CRS_PROPERTIES,
	}

	MOCK_PARTIAL_POSTCODE_CLIENT_RESULT_LAT_LONG = models.PartialPostcodeClientResultLatLong{
		Type:        "MOCK_LAT_LONG_TYPE",
		Crs:         MOCK_PARTIAL_POSTCODE_CLIENT_RESULT_CRS,
		Coordinates: []float64{1, 2},
	}
})

var _ = Describe("Mappers", func() {
	It("maps the 'formatted' field correctly when all relevant fields are populated and returns the correctly mapped struct", func() {
		mapper := Mapper{}

		input := &models.PartialPostcodeResult{
			Name:               MOCK_NAME,
			Address1:           MOCK_ADDRESS1,
			Address2:           MOCK_ADDRESS2,
			Address3:           MOCK_ADDRESS3,
			Postcode:           MOCK_POSTCODE,
			City:               MOCK_CITY,
			Website:            MOCK_WEBSITE,
			SubType:            MOCK_SUBTYPE,
			Phone:              MOCK_PHONE,
			PartialPostCode:    MOCK_PARTIAL_POSTCODE,
			OrganisationType:   MOCK_ORGANISATION_TYPE,
			OrganisationId:     MOCK_ORGANISATION_ID,
			OrganisationStatus: MOCK_ORGANISATION_STATUS,
			OrganisationCode:   MOCK_ORGANISATION_CODE,
			Latitude:           MOCK_LATITUDE,
			Longitude:          MOCK_LONGITUDE,
			IsPimsManaged:      MOCK_IS_PIMS_MANAGED,
			Fax:                MOCK_FAX,
			Email:              MOCK_EMAIL,
			County:             MOCK_COUNTY,
			LatLong:            MOCK_PARTIAL_POSTCODE_RESULT_LAT_LONG,
		}

		expectedFormatted := fmt.Sprintf("%s (%s, %s, %s, %s, %s)", MOCK_NAME, MOCK_ADDRESS1, MOCK_ADDRESS2, MOCK_ADDRESS3, MOCK_POSTCODE, MOCK_CITY)

		expected := &models.PartialPostcodeClientResult{
			Name:               MOCK_NAME,
			Address1:           MOCK_ADDRESS1,
			Address2:           MOCK_ADDRESS2,
			Address3:           MOCK_ADDRESS3,
			Postcode:           MOCK_POSTCODE,
			City:               MOCK_CITY,
			Website:            MOCK_WEBSITE,
			SubType:            MOCK_SUBTYPE,
			Phone:              MOCK_PHONE,
			PartialPostCode:    MOCK_PARTIAL_POSTCODE,
			OrganisationType:   MOCK_ORGANISATION_TYPE,
			OrganisationId:     MOCK_ORGANISATION_ID,
			OrganisationStatus: MOCK_ORGANISATION_STATUS,
			OrganisationCode:   MOCK_ORGANISATION_CODE,
			Latitude:           MOCK_LATITUDE,
			Longitude:          MOCK_LONGITUDE,
			IsPimsManaged:      MOCK_IS_PIMS_MANAGED,
			Fax:                MOCK_FAX,
			Email:              MOCK_EMAIL,
			County:             MOCK_COUNTY,
			LatLong:            MOCK_PARTIAL_POSTCODE_CLIENT_RESULT_LAT_LONG,
			Formatted:          expectedFormatted,
		}

		actual := mapper.MapPartialPostcodeResult(input)

		Expect(actual).To(Equal(expected))
	})
})
