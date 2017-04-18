package models

type PartialPostcodeClientResponse struct {
	Success bool                          `json:"success"`
	Results []PartialPostcodeClientResult `json:"result"`
}

type PartialPostcodeClientResult struct {
	Website            string                             `json:"website"`
	SubType            string                             `json:"sub_type"`
	Postcode           string                             `json:"postcode"`
	Phone              string                             `json:"phone"`
	PartialPostCode    string                             `json:"partial_postcode"`
	OrganisationType   string                             `json:"organisation_type"`
	OrganisationStatus string                             `json:"organisation_status"`
	OrganisationId     string                             `json:"organisation_id"`
	OrganisationCode   string                             `json:"organisation_code"`
	Name               string                             `json:"name"`
	Longitude          string                             `json:"longitude"`
	Latitude           string                             `json:"latitude"`
	IsPimsManaged      string                             `json:"is_pims_managed"`
	Fax                string                             `json:"fax"`
	Email              string                             `json:"email"`
	County             string                             `json:"county"`
	City               string                             `json:"city"`
	Address1           string                             `json:"address1"`
	Address2           string                             `json:"address2"`
	Address3           string                             `json:"address3"`
	LatLong            PartialPostcodeClientResultLatLong `json:"latlong"`
	Formatted          string                             `json:"formatted"`
}

type PartialPostcodeClientResultLatLong struct {
	Type        string                         `json:"type"`
	Crs         PartialPostcodeClientResultCrs `json:"crs"`
	Coordinates []float64                      `json:"coordinates"`
}

type PartialPostcodeClientResultCrs struct {
	Type       string                                   `json:"type"`
	Properties PartialPostcodeClientResultCrsProperties `json:"properties"`
}

type PartialPostcodeClientResultCrsProperties struct {
	Name string `json:"name"`
}
