// This file was generated by counterfeiter
package tsafakes

import (
	"sync"

	"github.com/concourse/tsa"
)

type FakeTokenGenerator struct {
	GenerateTokenStub        func() (string, error)
	generateTokenMutex       sync.RWMutex
	generateTokenArgsForCall []struct{}
	generateTokenReturns     struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTokenGenerator) GenerateToken() (string, error) {
	fake.generateTokenMutex.Lock()
	fake.generateTokenArgsForCall = append(fake.generateTokenArgsForCall, struct{}{})
	fake.recordInvocation("GenerateToken", []interface{}{})
	fake.generateTokenMutex.Unlock()
	if fake.GenerateTokenStub != nil {
		return fake.GenerateTokenStub()
	} else {
		return fake.generateTokenReturns.result1, fake.generateTokenReturns.result2
	}
}

func (fake *FakeTokenGenerator) GenerateTokenCallCount() int {
	fake.generateTokenMutex.RLock()
	defer fake.generateTokenMutex.RUnlock()
	return len(fake.generateTokenArgsForCall)
}

func (fake *FakeTokenGenerator) GenerateTokenReturns(result1 string, result2 error) {
	fake.GenerateTokenStub = nil
	fake.generateTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeTokenGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateTokenMutex.RLock()
	defer fake.generateTokenMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTokenGenerator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ tsa.TokenGenerator = new(FakeTokenGenerator)