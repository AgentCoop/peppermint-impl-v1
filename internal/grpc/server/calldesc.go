package server

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/runtime/deps"
)

type callDesc struct {
	context.Context
	Request
	Response
	deps.Configurator
}

type GrpcCallDescriptor interface {
	context.Context
	Request
	Response
	GetRequest() Request
	AssignNewRequest(Request) Request
	GetResponse() Response
	AssignNewResponse(Response) Response
	GetConfigurator() deps.Configurator
}

func NewCallDescriptor(ctx context.Context, cfg deps.Configurator) *callDesc {
	return &callDesc{
		ctx,
		NewRequest(ctx),
		NewResponse(ctx),
		cfg,
	}
}

func (p *callDesc) GetRequest() Request {
	return p.Request
}

func (p *callDesc) GetConfigurator() deps.Configurator {
	return p.Configurator
}

// Replaces the base request with an extended one
func (p *callDesc) AssignNewRequest(new Request) Request {
	base := p.Request
	p.Request = new
	return base
}

func (p *callDesc) GetResponse() Response {
	return p.Response
}

// Replaces the base response with an extended one
func (p *callDesc) AssignNewResponse(new Response) Response {
	base := p.Response
	p.Response = new
	return base
}
