package impl

import (
	"github.com/btwiuse/wetty/pkg/utils"
	"google.golang.org/grpc"

	"github.com/btwiuse/conntroll/pkg/agent"
	"github.com/btwiuse/conntroll/pkg/api"
	grpcimpl "github.com/btwiuse/conntroll/pkg/api/grpc/impl"
)

type RPC struct {
	Agent agent.Agent
}

type RPCRequest struct{}

type RPCResponse struct{}

func (c *RPC) New(req RPCRequest, res *RPCResponse) error {
	c.Agent.Go(c.Agent.ConnectAndServe)
	println("rpc.new called")
	return nil
}

type Session struct {
	Agent agent.Agent
}

type SessionRequest struct{}

type SessionResponse struct{}

func (c *Session) New(req SessionRequest, res *SessionResponse) error {
	conn, err := c.Agent.CreateSession()
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	api.RegisterSessionServer(grpcServer, &grpcimpl.Session{
		TtyFactory: c.Agent,
	})
	grpcServer.Serve(&utils.SingleListener{conn})
	return nil
}
