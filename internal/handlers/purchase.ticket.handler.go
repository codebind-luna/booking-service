package handlers

import (
	"context"
	"errors"

	bookingv1 "github.com/codebind-luna/booking-service/gen/go/booking/v1"
	"github.com/codebind-luna/booking-service/internal/constants"
	"github.com/codebind-luna/booking-service/internal/exceptions"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *ticketServiceHandlers) PurchaseTicket(ctx context.Context, pr *bookingv1.PurchaseTicketRequest) (*bookingv1.PurchaseTicketResponse, error) {
	user := pr.GetUser()
	fromCity := pr.GetFromCity()
	toCity := pr.GetToCity()
	price := pr.GetPrice()

	bookingID, err := h.svc.PurchaseTicket(ctx, user.GetEmail(), user.GetFirstName(), user.GetLastName(), fromCity, toCity, float64(price))

	if err != nil {
		if errors.Is(err, exceptions.ErrNoSeatsAvailable) {
			return nil, status.New(codes.FailedPrecondition, "No seats available").Err()
		}
		if errors.Is(err, exceptions.ErrUserHasPurchasedTicketAlready) {
			return nil, status.New(codes.ResourceExhausted, "Already purchased a ticket").Err()
		}

		return nil, status.New(codes.Unknown, "some unknow error happened").Err()
	}

	r := &bookingv1.PurchaseTicketResponse{}
	r.BookingId = *bookingID
	r.Message = constants.SuccessPurchaseTicketMessage
	r.Success = true

	return r, nil
}
