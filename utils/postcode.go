package utils

import (
	"errors"
	"fmt"
	"strings"
)

type Utils struct{}

func (this *Utils) formatPostcode(postcode string) (string, error) {
	formatted := strings.Replace(postcode, " ", "", -1)
	length := len(formatted)

	if length < 5 || length > 7 {
		return "", errors.New(fmt.Sprintf("invalid postcode : %s", postcode))
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
