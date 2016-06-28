// This file was generated by counterfeiter
package fake_authenticators

import (
	"sync"

	"code.cloudfoundry.org/diego-ssh/authenticators"
	"github.com/pivotal-golang/lager"
	"golang.org/x/crypto/ssh"
)

type FakePermissionsBuilder struct {
	BuildStub        func(logger lager.Logger, processGuid string, index int, metadata ssh.ConnMetadata) (*ssh.Permissions, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		logger      lager.Logger
		processGuid string
		index       int
		metadata    ssh.ConnMetadata
	}
	buildReturns struct {
		result1 *ssh.Permissions
		result2 error
	}
}

func (fake *FakePermissionsBuilder) Build(logger lager.Logger, processGuid string, index int, metadata ssh.ConnMetadata) (*ssh.Permissions, error) {
	fake.buildMutex.Lock()
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		logger      lager.Logger
		processGuid string
		index       int
		metadata    ssh.ConnMetadata
	}{logger, processGuid, index, metadata})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub(logger, processGuid, index, metadata)
	} else {
		return fake.buildReturns.result1, fake.buildReturns.result2
	}
}

func (fake *FakePermissionsBuilder) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakePermissionsBuilder) BuildArgsForCall(i int) (lager.Logger, string, int, ssh.ConnMetadata) {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return fake.buildArgsForCall[i].logger, fake.buildArgsForCall[i].processGuid, fake.buildArgsForCall[i].index, fake.buildArgsForCall[i].metadata
}

func (fake *FakePermissionsBuilder) BuildReturns(result1 *ssh.Permissions, result2 error) {
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 *ssh.Permissions
		result2 error
	}{result1, result2}
}

var _ authenticators.PermissionsBuilder = new(FakePermissionsBuilder)
