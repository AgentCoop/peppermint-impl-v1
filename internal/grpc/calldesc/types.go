package calldesc

import (
	"context"
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc/metadata"
)

type DescriptorType int

const (
	ServerType DescriptorType = iota
	ClientType
)

type common struct {
	context.Context
	typ       DescriptorType
	meta      meta
	secPolicy *secPolicy
}

type secPolicy struct {
	encKey  []byte
	e2e_Enc bool
}

type meta struct {
	parent  *common
	header  metadata.MD
	trailer metadata.MD
	sId     internal.SessionId
	nodeId  internal.NodeId
}

type ServerDescriptor struct {
	common
	svcCfg  runtime.ServiceConfigurator
	method  runtime.Method
	reqData grpc.RequestData
	resData grpc.ResponseData
}

func (s *ServerDescriptor) Policy() runtime.MethodCallPolicy {
	panic("implement me")
}

type ClientDescriptor struct {
	common
	policy runtime.MethodCallPolicy
}

func (s *ClientDescriptor) Policy() runtime.MethodCallPolicy {
	panic("implement me")
}

