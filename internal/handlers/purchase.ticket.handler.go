package handlers

import (
	"context"

	bookingv1 "github.com/codebind-luna/booking-service/gen/go/booking/v1"
)

const (
	successPurchaseTicketMessage string = "successfully purchased ticket"
)

func (h *ticketServiceHandlers) PurchaseTicket(ctx context.Context, pr *bookingv1.PurchaseTicketRequest) (*bookingv1.PurchaseTicketResponse, error) {
	user := pr.GetUser()
	fromCity := pr.GetFromCity()
	toCity := pr.GetToCity()
	price := pr.GetPrice()

	bookingID, err := h.svc.PurchaseTicket(ctx, user.GetEmail(), user.GetFirstName(), user.GetLastName(), fromCity, toCity, float64(price))

	r := &bookingv1.PurchaseTicketResponse{}
	if err != nil {
		r.Message = err.Error()
		return r, nil
	}

	r.BookingId = *bookingID
	r.Message = successPurchaseTicketMessage
	r.Success = true

	return r, nil
}
