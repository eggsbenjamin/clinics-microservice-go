package models

type PartialPostcodeClientResponse struct {
	Success bool                    `json:"success",omitempty`
	Results []PartialPostcodeResult `json:"result"`
}

type PartialPostcodeClientResult struct {
	Website            string                        `json:"website",omitempty`
	SubType            string                        `json:"sub_type",omitempty`
	Postcode           string                        `json:"postcode,omitempty"`
	Phone              string                        `json:"phone",omitempty`
	PartialPostCode    string                        `json:"partial_postcode",omitempty`
	OrganisationType   string                        `json:"organisation_type",omitempty`
	OrganisationStatus string                        `json:"organisation_status",omitempty`
	OrganisationId     string                        `json:"organisation_id",omitempty`
	OrganisationCode   string                        `json:"organisation_code",omitempty`
	Name               string                        `json:"name",omitempty`
	Longitude          string                        `json:"longitude",omitempty`
	Latitude           string                        `json:"latitude",omitempty`
	IsPimsManaged      string                        `json:"is_pims_managed",omitempty`
	Fax                string                        `json:"fax",omitempty`
	Email              string                        `json:"email",omitempty`
	County             string                        `json:"county",omitempty`
	City               string                        `json:"city",omitempty`
	Address1           string                        `json:"address1",omitempty`
	Address2           string                        `json:"address2",omitempty`
	Address3           string                        `json:"address3",omitempty`
	LatLong            *PartialPostcodeResultLatLong `json:"latong"`
	Formatted          string                        `json:"formatted",omitempty`
}
