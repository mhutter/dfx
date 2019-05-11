package queue

import (
	"context"
	"log"
	"net"

	"github.com/mhutter/dfx"
	"google.golang.org/grpc"
)

var (
	void = &dfx.Void{}
)

// Server implements the QueueServer interface
type Server struct {
	Queue chan *dfx.Deployable

	grpcServer *grpc.Server
}

// NewServer creates a new Server instance
func NewServer() *Server {
	return &Server{
		Queue: make(chan *dfx.Deployable, 10),
	}
}

// Listen starts the GRPC server and listens for incoming calls
func (s *Server) Listen(addr string) {
	sock, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on '%s': %s\n", addr, err)
	}

	s.grpcServer = grpc.NewServer()
	dfx.RegisterQueueServer(s.grpcServer, s)

	go func() {
		log.Fatalln(s.grpcServer.Serve(sock))
	}()
}

// Stop stops the GRPC server gracefully
func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}

// PostEvent receives an event containing a deployable
func (s *Server) PostEvent(ctx context.Context, d *dfx.Deployable) (*dfx.Void, error) {
	log.Printf("Received event: %v\n", d)
	return void, nil
}
