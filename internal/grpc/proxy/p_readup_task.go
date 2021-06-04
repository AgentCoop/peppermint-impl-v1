package proxy

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"io"
)

func (c *proxyConn) readUpstreamTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		var err error
		recvRaw := codec.NewRawPacket(nil, c.upstream.EncKey())
		err = c.upstream.Recv(recvRaw)
		task.Assert(err)
		if err == io.EOF {
			task.Done()
			return
		}
		c.upstreamChan <- recvRaw
		task.Tick()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}


