package main

import (
	"context"
	"log"
	"time"

	bookingv1 "github.com/codebind-luna/booking-service/gen/go/booking/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := "localhost:50051"

	// Set up a connection to the server.
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	c := bookingv1.NewTicketServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	details := []struct {
		email string
		fn    string
		ln    string
	}{
		{
			email: "john@gmail.com",
			fn:    "John",
			ln:    "Smith",
		},
		{
			email: "adam@gmail.com",
			fn:    "Adam",
			ln:    "Lesiuk",
		},
		{
			email: "Brian@gmail.com",
			fn:    "Brian",
			ln:    "Nwidge",
		},
		{
			email: "andrew@gmail.com",
			fn:    "Andrew",
			ln:    "Cahil",
		},
		{
			email: "Shireen@gmail.com",
			fn:    "Shireen",
			ln:    "Bailey",
		},
		{
			email: "Matt@gmail.com",
			fn:    "Matt",
			ln:    "Earnest",
		},
	}

	for _, d := range details {

		r, err := c.PurchaseTicket(ctx, &bookingv1.PurchaseTicketRequest{
			User: &bookingv1.User{
				Email:     d.email,
				FirstName: d.fn,
				LastName:  d.ln,
			},
			FromCity: "New York",
			ToCity:   "New Jersey",
			Price:    20.00,
		})
		if err != nil {
			log.Printf("Received error: %v\n", err)
			return
		}
		log.Printf("Received response: %s\n", r.Message)
		log.Printf("Booking ID: %s\n", r.BookingId)

		receipt, rErr := c.GetReceipt(ctx, &bookingv1.GetReceiptRequest{
			Email: d.email,
		})

		if rErr != nil {
			log.Printf("failed to get receipt for user with email %s\n: %v\n", d.email, rErr)
			return
		}

		log.Printf("receipt of the ticket details for user with email %s\n: %+v\n", d.email, receipt.Details)
	}

	ss := []bookingv1.Section{bookingv1.Section_SECTION_A, bookingv1.Section_SECTION_B}

	for _, s := range ss {
		seatM, sErr := c.ViewSeatMap(ctx, &bookingv1.ViewSeatMapRequest{
			Section: s,
		})

		if sErr != nil {
			log.Printf("failed to get seat map by section %s\n: %v\n", s, sErr)
			return
		}

		seats := map[string][]struct {
			user   string
			seatNo int
			status string
		}{}

		for _, st := range seatM.Seats {
			s1 := struct {
				user   string
				seatNo int
				status string
			}{

				seatNo: int(st.GetSeatNo()),
				status: st.GetStatus().String(),
			}
			if st.GetEmail() != "" {
				s1.user = st.GetEmail()
			}
			seats[s.String()] = append(seats[s.String()], s1)
		}

		log.Printf("seat map details by section %s:\n %+v\n", s, seats)
	}
}
