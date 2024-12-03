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

func (h *ticketServiceHandlers) GetReceipt(ctx context.Context, rr *bookingv1.GetReceiptRequest) (*bookingv1.GetReceiptResponse, error) {
	ticket, err := h.svc.GetReceipt(ctx, rr.GetEmail())

	if err != nil {
		if errors.Is(err, exceptions.ErrNoBookingFoundForUser) {
			return nil, status.New(codes.NotFound, "No booking found").Err()
		}

		return nil, status.New(codes.Unknown, "some unknow error happened").Err()
	}
	res := &bookingv1.GetReceiptResponse{}

	res.Message = constants.SuccessGetReceiptMessage
	res.Success = true

	res.Details = &bookingv1.Receipt{
		BookingId: ticket.BookingID(),
		FromCity:  ticket.FromCity(),
		ToCity:    ticket.ToCity(),
		User: &bookingv1.User{
			Email:     ticket.User().Email(),
			FirstName: ticket.User().FirstName(),
			LastName:  ticket.User().LastName(),
		},
		Price:   float32(ticket.PricePaid()),
		Section: ticket.Seat().Section().Section(),
		SeatNo:  int32(ticket.Seat().SeatNo()),
	}
	return res, nil
}
