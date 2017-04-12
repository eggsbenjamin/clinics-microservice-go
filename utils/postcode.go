package utils

import (
	"strings"
	"errors"
	"fmt"
)

func formatPostcode(postcode string) (string, error) {
	formatted := strings.Replace(postcode, " ", "", -1)
	length := len(formatted)

	if length < 5 || length > 7 {
		return "", errors.New(fmt.Sprintf("invalid postcode : %s", postcode))
	}

	formatted = formatted[0:len(formatted)-3] + " " + formatted[len(formatted)-3:len(formatted)]

	return strings.ToUpper(formatted), nil
}

func GetOutwardCode(postcode string) (string, error) {
	formatted, err := formatPostcode(postcode)

	if err != nil {
		return "", err
	}

	return strings.Split(formatted, " ")[0], nil
}
