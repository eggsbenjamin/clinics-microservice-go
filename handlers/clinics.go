package handlers

import (
	"github.com/eggsbenjamin/clinics-microservice-go/mappers"
	clinics "github.com/eggsbenjamin/clinics-microservice-go/services"
	"github.com/eggsbenjamin/clinics-microservice-go/utils"
)

type ClinicsByPostcode struct {
	clinics clinics.ClinicsService
	mapper  mappers.Mapper
}

func (this *ClinicsByPostcode) Handle(c *gin.Context) {

}
