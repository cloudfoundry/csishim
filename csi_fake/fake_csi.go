// This file was generated by counterfeiter
package csi_fake

import (
	"sync"

	"code.cloudfoundry.org/csishim"
	"code.cloudfoundry.org/goshims/grpcshim"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
)

type FakeCsi struct {
	NewIdentityClientStub        func(cc *grpc.ClientConn) csi.IdentityClient
	newIdentityClientMutex       sync.RWMutex
	newIdentityClientArgsForCall []struct {
		cc *grpc.ClientConn
	}
	newIdentityClientReturns struct {
		result1 csi.IdentityClient
	}
	RegisterIdentityServerStub        func(s *grpc.Server, srv csi.IdentityServer)
	registerIdentityServerMutex       sync.RWMutex
	registerIdentityServerArgsForCall []struct {
		s   *grpc.Server
		srv csi.IdentityServer
	}
	NewControllerClientStub        func(cc *grpc.ClientConn) csi.ControllerClient
	newControllerClientMutex       sync.RWMutex
	newControllerClientArgsForCall []struct {
		cc *grpc.ClientConn
	}
	newControllerClientReturns struct {
		result1 csi.ControllerClient
	}
	RegisterControllerServerStub        func(s *grpc.Server, srv csi.ControllerServer)
	registerControllerServerMutex       sync.RWMutex
	registerControllerServerArgsForCall []struct {
		s   *grpc.Server
		srv csi.ControllerServer
	}
	NewNodeClientStub        func(cc grpcshim.ClientConn) csi.NodeClient
	newNodeClientMutex       sync.RWMutex
	newNodeClientArgsForCall []struct {
		cc grpcshim.ClientConn
	}
	newNodeClientReturns struct {
		result1 csi.NodeClient
	}
	RegisterNodeServerStub        func(s *grpc.Server, srv csi.NodeServer)
	registerNodeServerMutex       sync.RWMutex
	registerNodeServerArgsForCall []struct {
		s   *grpc.Server
		srv csi.NodeServer
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCsi) NewIdentityClient(cc *grpc.ClientConn) csi.IdentityClient {
	fake.newIdentityClientMutex.Lock()
	fake.newIdentityClientArgsForCall = append(fake.newIdentityClientArgsForCall, struct {
		cc *grpc.ClientConn
	}{cc})
	fake.recordInvocation("NewIdentityClient", []interface{}{cc})
	fake.newIdentityClientMutex.Unlock()
	if fake.NewIdentityClientStub != nil {
		return fake.NewIdentityClientStub(cc)
	}
	return fake.newIdentityClientReturns.result1
}

func (fake *FakeCsi) NewIdentityClientCallCount() int {
	fake.newIdentityClientMutex.RLock()
	defer fake.newIdentityClientMutex.RUnlock()
	return len(fake.newIdentityClientArgsForCall)
}

func (fake *FakeCsi) NewIdentityClientArgsForCall(i int) *grpc.ClientConn {
	fake.newIdentityClientMutex.RLock()
	defer fake.newIdentityClientMutex.RUnlock()
	return fake.newIdentityClientArgsForCall[i].cc
}

func (fake *FakeCsi) NewIdentityClientReturns(result1 csi.IdentityClient) {
	fake.NewIdentityClientStub = nil
	fake.newIdentityClientReturns = struct {
		result1 csi.IdentityClient
	}{result1}
}

func (fake *FakeCsi) RegisterIdentityServer(s *grpc.Server, srv csi.IdentityServer) {
	fake.registerIdentityServerMutex.Lock()
	fake.registerIdentityServerArgsForCall = append(fake.registerIdentityServerArgsForCall, struct {
		s   *grpc.Server
		srv csi.IdentityServer
	}{s, srv})
	fake.recordInvocation("RegisterIdentityServer", []interface{}{s, srv})
	fake.registerIdentityServerMutex.Unlock()
	if fake.RegisterIdentityServerStub != nil {
		fake.RegisterIdentityServerStub(s, srv)
	}
}

func (fake *FakeCsi) RegisterIdentityServerCallCount() int {
	fake.registerIdentityServerMutex.RLock()
	defer fake.registerIdentityServerMutex.RUnlock()
	return len(fake.registerIdentityServerArgsForCall)
}

func (fake *FakeCsi) RegisterIdentityServerArgsForCall(i int) (*grpc.Server, csi.IdentityServer) {
	fake.registerIdentityServerMutex.RLock()
	defer fake.registerIdentityServerMutex.RUnlock()
	return fake.registerIdentityServerArgsForCall[i].s, fake.registerIdentityServerArgsForCall[i].srv
}

func (fake *FakeCsi) NewControllerClient(cc *grpc.ClientConn) csi.ControllerClient {
	fake.newControllerClientMutex.Lock()
	fake.newControllerClientArgsForCall = append(fake.newControllerClientArgsForCall, struct {
		cc *grpc.ClientConn
	}{cc})
	fake.recordInvocation("NewControllerClient", []interface{}{cc})
	fake.newControllerClientMutex.Unlock()
	if fake.NewControllerClientStub != nil {
		return fake.NewControllerClientStub(cc)
	}
	return fake.newControllerClientReturns.result1
}

func (fake *FakeCsi) NewControllerClientCallCount() int {
	fake.newControllerClientMutex.RLock()
	defer fake.newControllerClientMutex.RUnlock()
	return len(fake.newControllerClientArgsForCall)
}

func (fake *FakeCsi) NewControllerClientArgsForCall(i int) *grpc.ClientConn {
	fake.newControllerClientMutex.RLock()
	defer fake.newControllerClientMutex.RUnlock()
	return fake.newControllerClientArgsForCall[i].cc
}

func (fake *FakeCsi) NewControllerClientReturns(result1 csi.ControllerClient) {
	fake.NewControllerClientStub = nil
	fake.newControllerClientReturns = struct {
		result1 csi.ControllerClient
	}{result1}
}

func (fake *FakeCsi) RegisterControllerServer(s *grpc.Server, srv csi.ControllerServer) {
	fake.registerControllerServerMutex.Lock()
	fake.registerControllerServerArgsForCall = append(fake.registerControllerServerArgsForCall, struct {
		s   *grpc.Server
		srv csi.ControllerServer
	}{s, srv})
	fake.recordInvocation("RegisterControllerServer", []interface{}{s, srv})
	fake.registerControllerServerMutex.Unlock()
	if fake.RegisterControllerServerStub != nil {
		fake.RegisterControllerServerStub(s, srv)
	}
}

func (fake *FakeCsi) RegisterControllerServerCallCount() int {
	fake.registerControllerServerMutex.RLock()
	defer fake.registerControllerServerMutex.RUnlock()
	return len(fake.registerControllerServerArgsForCall)
}

func (fake *FakeCsi) RegisterControllerServerArgsForCall(i int) (*grpc.Server, csi.ControllerServer) {
	fake.registerControllerServerMutex.RLock()
	defer fake.registerControllerServerMutex.RUnlock()
	return fake.registerControllerServerArgsForCall[i].s, fake.registerControllerServerArgsForCall[i].srv
}

func (fake *FakeCsi) NewNodeClient(cc grpcshim.ClientConn) csi.NodeClient {
	fake.newNodeClientMutex.Lock()
	fake.newNodeClientArgsForCall = append(fake.newNodeClientArgsForCall, struct {
		cc grpcshim.ClientConn
	}{cc})
	fake.recordInvocation("NewNodeClient", []interface{}{cc})
	fake.newNodeClientMutex.Unlock()
	if fake.NewNodeClientStub != nil {
		return fake.NewNodeClientStub(cc)
	}
	return fake.newNodeClientReturns.result1
}

func (fake *FakeCsi) NewNodeClientCallCount() int {
	fake.newNodeClientMutex.RLock()
	defer fake.newNodeClientMutex.RUnlock()
	return len(fake.newNodeClientArgsForCall)
}

func (fake *FakeCsi) NewNodeClientArgsForCall(i int) grpcshim.ClientConn {
	fake.newNodeClientMutex.RLock()
	defer fake.newNodeClientMutex.RUnlock()
	return fake.newNodeClientArgsForCall[i].cc
}

func (fake *FakeCsi) NewNodeClientReturns(result1 csi.NodeClient) {
	fake.NewNodeClientStub = nil
	fake.newNodeClientReturns = struct {
		result1 csi.NodeClient
	}{result1}
}

func (fake *FakeCsi) RegisterNodeServer(s *grpc.Server, srv csi.NodeServer) {
	fake.registerNodeServerMutex.Lock()
	fake.registerNodeServerArgsForCall = append(fake.registerNodeServerArgsForCall, struct {
		s   *grpc.Server
		srv csi.NodeServer
	}{s, srv})
	fake.recordInvocation("RegisterNodeServer", []interface{}{s, srv})
	fake.registerNodeServerMutex.Unlock()
	if fake.RegisterNodeServerStub != nil {
		fake.RegisterNodeServerStub(s, srv)
	}
}

func (fake *FakeCsi) RegisterNodeServerCallCount() int {
	fake.registerNodeServerMutex.RLock()
	defer fake.registerNodeServerMutex.RUnlock()
	return len(fake.registerNodeServerArgsForCall)
}

func (fake *FakeCsi) RegisterNodeServerArgsForCall(i int) (*grpc.Server, csi.NodeServer) {
	fake.registerNodeServerMutex.RLock()
	defer fake.registerNodeServerMutex.RUnlock()
	return fake.registerNodeServerArgsForCall[i].s, fake.registerNodeServerArgsForCall[i].srv
}

func (fake *FakeCsi) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newIdentityClientMutex.RLock()
	defer fake.newIdentityClientMutex.RUnlock()
	fake.registerIdentityServerMutex.RLock()
	defer fake.registerIdentityServerMutex.RUnlock()
	fake.newControllerClientMutex.RLock()
	defer fake.newControllerClientMutex.RUnlock()
	fake.registerControllerServerMutex.RLock()
	defer fake.registerControllerServerMutex.RUnlock()
	fake.newNodeClientMutex.RLock()
	defer fake.newNodeClientMutex.RUnlock()
	fake.registerNodeServerMutex.RLock()
	defer fake.registerNodeServerMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCsi) recordInvocation(key string, args []interface{}) {
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

var _ csishim.Csi = new(FakeCsi)
