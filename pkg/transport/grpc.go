package transport

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	bookingv1 "github.com/codebind-luna/booking-service/gen/go/booking/v1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server interface {
	Start()
	Stop(ctx context.Context)
}

var _ Server = (*serverImp)(nil)

type serverImp struct {
	logger *log.Logger
	port   int
	l      net.Listener
	srv    *grpc.Server
}

func NewServer(logger *log.Logger, host string, port int, svc bookingv1.TicketServiceServer) Server {
	// Set up a listener on port 50051
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the Ticket service with the server
	bookingv1.RegisterTicketServiceServer(s, svc)

	return &serverImp{
		logger: logger,
		l:      listener,
		port:   port,
		srv:    s,
	}
}

func (s *serverImp) Start() {
	s.logger.Infof("starting internal grpc server on :%d", s.port)

	go func() {
		if err := s.srv.Serve(s.l); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Fatalf("error starting grpc server on port %d : %s\n", s.port, err.Error())
		}
	}()
}

func (s *serverImp) Stop(ctx context.Context) {
	s.logger.Infof("initializing stop of grpc server on port %d", s.port)
	s.srv.GracefulStop()
}
