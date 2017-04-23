package system_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eggsbenjamin/clinics-microservice-go/constants"
	"github.com/eggsbenjamin/clinics-microservice-go/fixtures"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("System", func() {
	It("responds to GET /clinics/postcode/:postcode with the correct formatted response", func() {
		url := fmt.Sprintf("http://%s/clinics/postcode/TW208DE", constants.URL)

		res, err := http.Get(url)

		Expect(err).NotTo(HaveOccurred())

		defer res.Body.Close()

		Expect(res.StatusCode).To(Equal(200))

		body, err := ioutil.ReadAll(res.Body)

		Expect(err).NotTo(HaveOccurred())

		Expect(body).To(MatchJSON(fixtures.EXPECTED_FORMATTED_POSTCODE_RESPONSE))
	})
})
