package handlers

import (
	"context"

	bookingv1 "github.com/codebind-luna/booking-service/gen/go/booking/v1"
)

const (
	successGetReceiptMessage string = "successfully purchased ticket"
)

func (h *ticketServiceHandlers) GetReceipt(ctx context.Context, rr *bookingv1.GetReceiptRequest) (*bookingv1.GetReceiptResponse, error) {
	ticket, err := h.svc.GetReceipt(ctx, rr.GetEmail())

	res := &bookingv1.GetReceiptResponse{}

	if err != nil {
		return res, nil
	}

	res.Message = successGetReceiptMessage
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
