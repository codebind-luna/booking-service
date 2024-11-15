package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	bookingv1 "github.com/codebind-luna/booking-service/gen/go/booking/v1"
	"github.com/codebind-luna/booking-service/internal/config"
	"github.com/codebind-luna/booking-service/pkg/logger"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BookingServiceClient struct {
	logger *logrus.Logger
	client bookingv1.TicketServiceClient
	conn   *grpc.ClientConn
}

func NewBookingServiceClient() (*BookingServiceClient, error) {
	config, createConfigErr := config.NewConfig()

	if createConfigErr != nil {
		return nil, createConfigErr
	}

	logger := logger.ConfigureLogging()

	addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}

	client := bookingv1.NewTicketServiceClient(conn)

	return &BookingServiceClient{
		client: client,
		conn:   conn,
		logger: logger,
	}, nil
}

// Close closes the gRPC connection
func (bsc *BookingServiceClient) Close() {
	if err := bsc.conn.Close(); err != nil {
		bsc.logger.Fatalf("Error closing connection: %v", err)
	}
}

type purchaseDetails struct {
	email string
	fn    string
	ln    string
	fc    string
	tc    string
	p     float64
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const domains = "com net org io co"

// generateRandomString generates a random string of a given length
func generateRandomString(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}

// generateRandomEmail generates a random email address
func generateRandomEmail() string {
	username := generateRandomString(10)                // Random username
	domainName := generateRandomString(5)               // Random domain name
	domainExtension := domains[rand.Intn(len(domains))] // Random domain extension
	return fmt.Sprintf("%s@%s.%s", username, domainName, domainExtension)
}

var (
	details purchaseDetails = purchaseDetails{
		email: generateRandomEmail(),
		fn:    "John",
		ln:    "Smith",
		fc:    "New York",
		tc:    "New Jersey",
		p:     20.00,
	}
)

var (
	section bookingv1.Section
	seatNo  int
)

func (bsc *BookingServiceClient) PurchaseTicket() {
	// Call the SayHello method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := bsc.client.PurchaseTicket(ctx, &bookingv1.PurchaseTicketRequest{
		User: &bookingv1.User{
			Email:     details.email,
			FirstName: details.fn,
			LastName:  details.ln,
		},
		FromCity: details.fc,
		ToCity:   details.tc,
		Price:    float32(details.p),
	})

	if err != nil {
		bsc.logger.Fatalf("Received error: %+v", err)
	}

	bsc.logger.Infof("output purchase ticket:\n %+v", res)
}

func (bsc *BookingServiceClient) GetReceipt() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, rErr := bsc.client.GetReceipt(ctx, &bookingv1.GetReceiptRequest{
		Email: details.email,
	})

	if rErr != nil {
		bsc.logger.Fatalf("failed to get receipt %v", rErr)
	}

	section = res.GetDetails().Section
	seatNo = int(res.GetDetails().GetSeatNo())

	bsc.logger.Printf("output get receipt:\n %+v", res)
}

func (bsc *BookingServiceClient) ViewSeatMap() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, sErr := bsc.client.ViewSeatMap(ctx, &bookingv1.ViewSeatMapRequest{
		Section: bookingv1.Section_SECTION_A,
	})

	if sErr != nil {
		bsc.logger.Fatalf("failed to view seat map %v", sErr)
	}

	bsc.logger.Printf("output view seat map for:%s\n %+v", bookingv1.Section_SECTION_A.String(), res)

	res, sErr = bsc.client.ViewSeatMap(ctx, &bookingv1.ViewSeatMapRequest{
		Section: bookingv1.Section_SECTION_B,
	})

	if sErr != nil {
		bsc.logger.Fatalf("failed to view seat map for:%s\n %v", bookingv1.Section_SECTION_B.String(), sErr)
	}

	bsc.logger.Printf("output view seat map for:%s\n %+v", bookingv1.Section_SECTION_B.String(), res)
}

func (bsc *BookingServiceClient) ModifySeat() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	seat := 1
	if seatNo == 1 {
		seat = 2
	}
	res, err := bsc.client.ModifySeat(ctx, &bookingv1.ModifySeatRequest{
		Email:   details.email,
		Section: section,
		SeatNo:  int32(seat),
	})

	if err != nil {
		bsc.logger.Fatalf("failed to modify seat for user with email:%s\n %v", details.email, err)
	}

	bsc.logger.Printf("output modify seat for user with emaail:%s\n %+v", details.email, res)
}

func (bsc *BookingServiceClient) RemoveUser() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := bsc.client.RemoveUser(ctx, &bookingv1.RemoveUserRequest{
		Email: details.email,
	})

	if err != nil {
		bsc.logger.Fatalf("failed to remove user with email %s from train:\n %v", details.email, err)
	}

	bsc.logger.Printf("output remove user with email %s from train:\n %+v", details.email, res)
}

func main() {
	client, err := NewBookingServiceClient()
	if err != nil {
		log.Fatalf("could not create client: %v", err)
	}
	defer client.Close()

	client.PurchaseTicket()
	client.GetReceipt()
	client.ViewSeatMap()
	client.ModifySeat()
	client.RemoveUser()
}
