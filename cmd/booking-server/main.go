package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codebind-luna/booking-service/internal/app/services/ticket"
	"github.com/codebind-luna/booking-service/internal/config"
	"github.com/codebind-luna/booking-service/internal/constants"
	"github.com/codebind-luna/booking-service/internal/domain"
	"github.com/codebind-luna/booking-service/internal/handlers"
	"github.com/codebind-luna/booking-service/internal/repositories"
	"github.com/codebind-luna/booking-service/pkg/logger"
	"github.com/codebind-luna/booking-service/pkg/transport"
)

func main() {
	config, createConfigErr := config.NewConfig()

	if createConfigErr != nil {
		log.Fatal(createConfigErr.Error())
	}

	logger := logger.ConfigureLogging()

	repoType, repoErr := domain.ParseRepository(config.Repository.Type)
	if repoErr != nil {
		logger.Fatal(repoErr.Error())
	}

	repo, rErr := repositories.New(logger, repoType)
	if rErr != nil {
		logger.Fatal(rErr.Error())
	}

	server := transport.NewServer(
		logger,
		config.Server.Host,
		config.Server.Port,
		handlers.NewTicketService(ticket.NewService(logger, repo)))

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
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultShutdownContextWaitSeconds*time.Second)
	defer cancel()

	// Wait until context times out or the server shuts down
	<-ctx.Done()
	logger.Info("Server stopped successfully.")

}
