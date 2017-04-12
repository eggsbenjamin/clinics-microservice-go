package utils

import (
	"testing"
	testUtils "github.com/eggsbenjamin/utils"
)

func TestGetOutwardCode(t *testing.T) {
	type Fixture struct {
		Input string `json:"input"`
		Output string `json:"output"`
	}

	testData := []Fixture{}

	testUtils.UnmarshalJsonFile("../fixtures/partial_postcode.json", &testData)

	for _, fixture := range testData {
		actual, err := GetOutwardCode(fixture.Input)
		expected := fixture.Output

		if err != nil {
			t.Logf("unexpected error %v", err)
			t.Fail()
		}

		if actual != expected {
			t.Logf("expected %s, received %s", expected, actual)
			t.Fail()
		}
	}
}
