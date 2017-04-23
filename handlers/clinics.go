package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eggsbenjamin/clinics-microservice-go/constants"
	"github.com/eggsbenjamin/clinics-microservice-go/mappers"
	clinics "github.com/eggsbenjamin/clinics-microservice-go/services"
	"github.com/julienschmidt/httprouter"
)

type ClinicsHandlers struct {
	clinicsService clinics.IClinicsService
	mapper         mappers.IMapper
}

func NewClinicsHandlers(clinicsService clinics.IClinicsService, mapper mappers.IMapper) *ClinicsHandlers {
	return &ClinicsHandlers{
		clinicsService: clinicsService,
		mapper:         mapper,
	}
}

func (this *ClinicsHandlers) ClinicsByPostcode(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	postcode := params.ByName("postcode")
	results, err := this.clinicsService.GetByPostcode(postcode)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		var (
			status    int
			clientErr string
		)
		errMsg := err.Error()
		switch errMsg {
		case constants.POSTCODE_ERROR:
			status = http.StatusBadRequest
			clientErr = constants.POSTCODE_ERROR
		default:
			status = http.StatusInternalServerError
			clientErr = constants.INTERNAL_SERVER_ERROR
		}

		w.WriteHeader(status)
		fmt.Fprintf(w, fmt.Sprintf(`{ "error" : "%s" }`, clientErr))
		return
	}

	payload, err := json.Marshal(this.mapper.MapPartialPostcodeResponse(results))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, fmt.Sprintf(`{ "error" : "%s" }`, constants.INTERNAL_SERVER_ERROR))
		return
	}

	w.Write(payload)
	return
}
