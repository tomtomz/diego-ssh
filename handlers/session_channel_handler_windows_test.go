// +build windows

package handlers_test

import (
	"github.com/cloudfoundry-incubator/diego-ssh/daemon"
	"github.com/cloudfoundry-incubator/diego-ssh/handlers"
	"github.com/cloudfoundry-incubator/diego-ssh/handlers/fakes"
	"github.com/cloudfoundry-incubator/diego-ssh/test_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-golang/lager/lagertest"
	"golang.org/x/crypto/ssh"
)

var _ = Describe("SessionChannelHandler", func() {
	var (
		sshd   *daemon.Daemon
		client *ssh.Client

		logger          *lagertest.TestLogger
		serverSSHConfig *ssh.ServerConfig

		runner                *fakes.FakeRunner
		shellLocator          *fakes.FakeShellLocator
		sessionChannelHandler *handlers.SessionChannelHandler

		newChannelHandlers map[string]handlers.NewChannelHandler
		defaultEnv         map[string]string
		connectionFinished chan struct{}
	)

	BeforeEach(func() {
		logger = lagertest.NewTestLogger("test")
		serverSSHConfig = &ssh.ServerConfig{
			NoClientAuth: true,
		}
		serverSSHConfig.AddHostKey(TestHostKey)

		runner = &fakes.FakeRunner{}
		shellLocator = &fakes.FakeShellLocator{}
		defaultEnv = map[string]string{}

		sessionChannelHandler = handlers.NewSessionChannelHandler()

		newChannelHandlers = map[string]handlers.NewChannelHandler{
			"session": sessionChannelHandler,
		}

		serverNetConn, clientNetConn := test_helpers.Pipe()

		sshd = daemon.New(logger, serverSSHConfig, nil, newChannelHandlers)
		connectionFinished = make(chan struct{})
		go func() {
			sshd.HandleConnection(serverNetConn)
			close(connectionFinished)
		}()

		client = test_helpers.NewClient(clientNetConn, nil)
	})

	AfterEach(func() {
		if client != nil {
			err := client.Close()
			Expect(err).NotTo(HaveOccurred())
		}
		Eventually(connectionFinished).Should(BeClosed())
	})

	Context("when a session is opened", func() {

		It("doesn't accept sessions", func() {
			_, sessionErr := client.NewSession()

			Expect(sessionErr).To(MatchError(ContainSubstring("not supported on windows")))
		})
	})
})
