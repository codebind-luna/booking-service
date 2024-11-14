package handlers

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	bookingv1 "github.com/codebind-luna/booking-service/gen/go/booking/v1"
	"github.com/codebind-luna/booking-service/internal/app"
	"github.com/codebind-luna/booking-service/internal/domain"
	"github.com/codebind-luna/booking-service/internal/repositories"
	"github.com/codebind-luna/booking-service/pkg/logger"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/test/bufconn"
)

// setup function to initialize server, repo, and service
func setup(t *testing.T) (context.Context, *grpc.ClientConn, bookingv1.TicketServiceClient, func()) {
	// Server Initialization
	lis := bufconn.Listen(1024 * 1024) // Allocate a buffer for the connection
	t.Cleanup(func() {
		lis.Close() // Ensure the listener is closed after the test
	})

	srv := grpc.NewServer() // Create a new gRPC server
	t.Cleanup(func() {
		srv.Stop() // Ensure the server is stopped after the test
	})

	// Initialize logger
	logger := logger.ConfigureLogging()

	// Set up the repository (in-memory in this case)
	repoType, repoErr := domain.ParseRepository("in-memory")
	if repoErr != nil {
		t.Fatal(repoErr.Error()) // Fail the test if there is an error parsing the repository type
	}

	repo, rErr := repositories.New(logger, repoType)
	if rErr != nil {
		t.Fatal(rErr.Error()) // Fail the test if the repository creation fails
	}

	// Initialize the service and register it with the server
	svc := NewTicketService(app.NewService(logger, repo))
	bookingv1.RegisterTicketServiceServer(srv, svc) // Register the service on the server

	// Start the server in a goroutine
	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("srv.Serve %v", err) // Log an error if the server fails to serve
		}
	}()

	// Test Setup: Dial the in-memory server using bufconn
	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial() // Use bufconn for in-memory communication
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	t.Cleanup(func() {
		cancel() // Ensure the context is canceled after the test
	})

	resolver.SetDefaultScheme("passthrough")
	// Create a gRPC connection using the dialer
	conn, err := grpc.NewClient("bufnet", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	t.Cleanup(func() {
		conn.Close() // Ensure the connection is closed after the test
	})
	if err != nil {
		t.Fatalf("grpc.DialContext %v", err) // Fail the test if dialing the gRPC server fails
	}

	// Create a client for the TicketService
	client := bookingv1.NewTicketServiceClient(conn)

	// Return the context, client, and a cleanup function
	return ctx, conn, client, func() {
		// Cleanup any additional resources if necessary
		cancel()
		conn.Close()
		lis.Close()
		srv.Stop()
	}
}

// Test function that uses the setup function
func TestTicketService_PurchaseTicket(t *testing.T) {
	// Call the setup function to initialize resources
	ctx, _, c, cleanup := setup(t)
	// Ensure cleanup happens after the test finishes
	defer cleanup()

	// Call the PurchaseTicket method on the client
	res, err := c.PurchaseTicket(ctx, &bookingv1.PurchaseTicketRequest{
		User: &bookingv1.User{
			Email:     "john@xyz.com",
			FirstName: "John",
			LastName:  "Smith",
		},
		FromCity: "New York",
		ToCity:   "New Jersey",
		Price:    20.00,
	})

	// Assert that there were no errors and that the response is as expected
	assert.NoError(t, err, "Expected no error, but got one")
	assert.Equal(t, successPurchaseTicketMessage, res.GetMessage(), "Expected success message, but got a different message")
	assert.Equal(t, true, res.GetSuccess(), "Expected success flag to be true")
}

func TestTicketService_GetReceipt(t *testing.T) {
	// Call the setup function to initialize resources
	ctx, _, c, cleanup := setup(t)
	// Ensure cleanup happens after the test finishes
	defer cleanup()

	details := []struct {
		email string
		fn    string
		ln    string
		fc    string
		tc    string
		p     float32
	}{
		{
			email: "adam@gmail.com",
			fn:    "Adam",
			ln:    "Lesiuk",
		},
	}

	// Call the PurchaseTicket method on the client
	c.PurchaseTicket(ctx, &bookingv1.PurchaseTicketRequest{
		User: &bookingv1.User{
			Email:     details[0].email,
			FirstName: details[0].fn,
			LastName:  details[0].ln,
		},
		FromCity: details[0].fc,
		ToCity:   details[0].tc,
		Price:    details[0].p,
	})

	receipt, rErr := c.GetReceipt(ctx, &bookingv1.GetReceiptRequest{
		Email: details[0].email,
	})

	// Assert that there were no errors and that the response is as expected
	assert.NoError(t, rErr, "Expected no error, but got one")
	assert.Equal(t, successGetReceiptMessage, receipt.GetMessage(), "Expected success message, but got a different message")
	assert.Equal(t, true, receipt.GetSuccess(), "Expected success flag to be true")
	assert.Equal(t, details[0].email, receipt.GetDetails().GetUser().GetEmail())
	assert.Equal(t, details[0].fn, receipt.GetDetails().GetUser().GetFirstName())
	assert.Equal(t, details[0].ln, receipt.GetDetails().GetUser().GetLastName())
	assert.Equal(t, details[0].fc, receipt.GetDetails().GetFromCity())
	assert.Equal(t, details[0].tc, receipt.GetDetails().GetToCity())
	assert.Equal(t, details[0].p, receipt.GetDetails().GetPrice())
	assert.NotEmpty(t, receipt.GetDetails().GetBookingId(), "Expected booking id to be not empty")
}

func TestTicketService_ViewSeatMap(t *testing.T) {
	// Call the setup function to initialize resources
	ctx, _, c, cleanup := setup(t)
	// Ensure cleanup happens after the test finishes
	defer cleanup()

	// Call the PurchaseTicket method on the client
	details := []struct {
		email string
		fn    string
		ln    string
		fc    string
		tc    string
		p     float32
	}{
		{
			email: "adam@gmail.com",
			fn:    "Adam",
			ln:    "Lesiuk",
		},
	}

	// Call the PurchaseTicket method on the client
	c.PurchaseTicket(ctx, &bookingv1.PurchaseTicketRequest{
		User: &bookingv1.User{
			Email:     details[0].email,
			FirstName: details[0].fn,
			LastName:  details[0].ln,
		},
		FromCity: details[0].fc,
		ToCity:   details[0].tc,
		Price:    details[0].p,
	})

	receipt, _ := c.GetReceipt(ctx, &bookingv1.GetReceiptRequest{
		Email: details[0].email,
	})

	section := receipt.GetDetails().GetSection()
	seatNo := receipt.GetDetails().GetSeatNo()

	res, sErr := c.ViewSeatMap(ctx, &bookingv1.ViewSeatMapRequest{
		Section: section.String(),
	})

	// sectionIndexMap := map[bookingv1.Section]int{
	// 	bookingv1.Section_SECTION_A: 0,
	// 	bookingv1.Section_SECTION_B: 1,
	// }

	// Assert that there were no errors and that the response is as expected
	assert.NoError(t, sErr, "Expected no error, but got one")
	assert.Equal(t, successViewMapMessage, res.GetMessage(), "Expected success message, but got a different message")
	assert.Equal(t, true, res.GetSuccess(), "Expected success flag to be true")
	assert.Equal(t, section, res.GetSeats()[seatNo-1].Section, "Expected section to be set correctly")
	assert.Equal(t, bookingv1.Status_STATUS_BOOKED, res.GetSeats()[seatNo-1].Status, "Expected status set to be booked")
}
