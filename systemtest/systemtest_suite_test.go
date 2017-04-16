package system_test

import (
	"net"
	"os/exec"

	"github.com/eggsbenjamin/clinics-microservice-go/constants"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var (
	binaryPath string
	session    *gexec.Session
)

var _ = BeforeSuite(func() {
	var err error

	binaryPath, err = gexec.Build("github.com/eggsbenjamin/clinics-microservice-go")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func verifyServerIsListening() error {
	_, err := net.Dial("tcp", constants.URL)
	return err
}

var _ = BeforeEach(func() {
	var err error

	session, err = gexec.Start(exec.Command(binaryPath), GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	Eventually(verifyServerIsListening).Should(Succeed())
})

var _ = AfterEach(func() {
	session.Interrupt()
	Eventually(session).Should(gexec.Exit())
})

func TestSystemtest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "System Test Suite")
}
