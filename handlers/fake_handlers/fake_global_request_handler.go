// This file was generated by counterfeiter
package fake_handlers

import (
	"sync"

	"code.cloudfoundry.org/diego-ssh/handlers"
	"github.com/pivotal-golang/lager"
	"golang.org/x/crypto/ssh"
)

type FakeGlobalRequestHandler struct {
	HandleRequestStub        func(logger lager.Logger, request *ssh.Request)
	handleRequestMutex       sync.RWMutex
	handleRequestArgsForCall []struct {
		logger  lager.Logger
		request *ssh.Request
	}
}

func (fake *FakeGlobalRequestHandler) HandleRequest(logger lager.Logger, request *ssh.Request) {
	fake.handleRequestMutex.Lock()
	fake.handleRequestArgsForCall = append(fake.handleRequestArgsForCall, struct {
		logger  lager.Logger
		request *ssh.Request
	}{logger, request})
	fake.handleRequestMutex.Unlock()
	if fake.HandleRequestStub != nil {
		fake.HandleRequestStub(logger, request)
	}
}

func (fake *FakeGlobalRequestHandler) HandleRequestCallCount() int {
	fake.handleRequestMutex.RLock()
	defer fake.handleRequestMutex.RUnlock()
	return len(fake.handleRequestArgsForCall)
}

func (fake *FakeGlobalRequestHandler) HandleRequestArgsForCall(i int) (lager.Logger, *ssh.Request) {
	fake.handleRequestMutex.RLock()
	defer fake.handleRequestMutex.RUnlock()
	return fake.handleRequestArgsForCall[i].logger, fake.handleRequestArgsForCall[i].request
}

var _ handlers.GlobalRequestHandler = new(FakeGlobalRequestHandler)
