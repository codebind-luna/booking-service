package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codebind-luna/booking-service/internal/app"
	"github.com/codebind-luna/booking-service/internal/domain"
	"github.com/codebind-luna/booking-service/internal/handlers"
	"github.com/codebind-luna/booking-service/internal/repositories"
	"github.com/codebind-luna/booking-service/pkg/logger"
	"github.com/codebind-luna/booking-service/pkg/transport"
)

const (
	DefaultShutdownContextWaitSeconds = 5
)

func main() {
	logger := logger.ConfigureLogging()

	repoType, repoErr := domain.ParseRepository("in-memory")
	if repoErr != nil {
		logger.Fatal(repoErr.Error())
	}

	repo, rErr := repositories.New(logger, repoType)
	if rErr != nil {
		logger.Fatal(rErr.Error())
	}

	server := transport.NewServer(logger, 50051, handlers.NewTicketService(app.NewService(logger, repo)))

	server.Start()

	// Create shutdown chan
	startShutdown := make(chan struct{})
	// Gracefully shut down
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigs
		signal.Stop(sigs)
		logger.Debugf("received signal %v\n", sig)
		close(startShutdown)

		// Allow pressing Ctrl+C again to exit, otherwise the developer must manually kill the process
		if sig == syscall.SIGINT {
			sigs2 := make(chan os.Signal, 1)
			signal.Notify(sigs2, syscall.SIGINT)
			logger.Debugf("press Ctrl+C again to exit")
			<-sigs2
			os.Exit(0)
		}
	}()

	// Block until a signal is received
	<-startShutdown
	// Create the cancelation context
	ctx, cancel := context.WithTimeout(context.Background(), DefaultShutdownContextWaitSeconds*time.Second)
	defer cancel()

	// Wait until context times out or the server shuts down
	<-ctx.Done()
	logger.Info("Server stopped successfully.")

}
