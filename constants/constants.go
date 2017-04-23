package constants

//	server constants
const (
	CLINICS_POSTCODE = "http://data.gov.uk/data/api/service/health/clinics/partial_postcode?partial_postcode="
	CLINICS_NAME     = "http://data.gov.uk/data/api/service/health/clinics/organisation_name?organisation_name="
	CLINICS_CITY     = "http://data.gov.uk/data/api/service/health/clinics?city="
	URL              = "localhost:8080"
)

//	error constants
const (
	SERVICE_UNAVAILABLE_ERROR = "SERVICE_UNAVAILABLE"
	CLIENT_ERROR              = "CLIENT_ERROR"
	POSTCODE_ERROR            = "POSTCODE_ERROR"
	INTERNAL_SERVER_ERROR     = "INTERNAL_SERVER_ERROR"
)
