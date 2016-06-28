// This file was generated by counterfeiter
package fake_handlers

import (
	"sync"

	"code.cloudfoundry.org/diego-ssh/handlers"
	"github.com/pivotal-golang/lager"
	"golang.org/x/crypto/ssh"
)

type FakeNewChannelHandler struct {
	HandleNewChannelStub        func(logger lager.Logger, newChannel ssh.NewChannel)
	handleNewChannelMutex       sync.RWMutex
	handleNewChannelArgsForCall []struct {
		logger     lager.Logger
		newChannel ssh.NewChannel
	}
}

func (fake *FakeNewChannelHandler) HandleNewChannel(logger lager.Logger, newChannel ssh.NewChannel) {
	fake.handleNewChannelMutex.Lock()
	fake.handleNewChannelArgsForCall = append(fake.handleNewChannelArgsForCall, struct {
		logger     lager.Logger
		newChannel ssh.NewChannel
	}{logger, newChannel})
	fake.handleNewChannelMutex.Unlock()
	if fake.HandleNewChannelStub != nil {
		fake.HandleNewChannelStub(logger, newChannel)
	}
}

func (fake *FakeNewChannelHandler) HandleNewChannelCallCount() int {
	fake.handleNewChannelMutex.RLock()
	defer fake.handleNewChannelMutex.RUnlock()
	return len(fake.handleNewChannelArgsForCall)
}

func (fake *FakeNewChannelHandler) HandleNewChannelArgsForCall(i int) (lager.Logger, ssh.NewChannel) {
	fake.handleNewChannelMutex.RLock()
	defer fake.handleNewChannelMutex.RUnlock()
	return fake.handleNewChannelArgsForCall[i].logger, fake.handleNewChannelArgsForCall[i].newChannel
}

var _ handlers.NewChannelHandler = new(FakeNewChannelHandler)
