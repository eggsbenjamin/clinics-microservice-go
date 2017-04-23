package utils

import (
	"errors"
	"regexp"
	"strings"

	"github.com/eggsbenjamin/clinics-microservice-go/constants"
)

type IUtils interface {
	GetOutwardCode(string) (string, error)
}

type Utils struct{}

func (this *Utils) formatPostcode(postcode string) (string, error) {
	formatted := strings.Replace(postcode, " ", "", -1)

	if !regexp.MustCompile(`^[A-Za-z0-9]{5,7}$`).MatchString(formatted) {
		return "", errors.New(constants.POSTCODE_ERROR)
	}

	formatted = formatted[0:len(formatted)-3] + " " + formatted[len(formatted)-3:len(formatted)]

	return strings.ToUpper(formatted), nil
}

func (this *Utils) GetOutwardCode(postcode string) (string, error) {
	formatted, err := this.formatPostcode(postcode)

	if err != nil {
		return "", err
	}

	return strings.Split(formatted, " ")[0], nil
}
