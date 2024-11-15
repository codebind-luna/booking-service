package handlers

import (
	"context"

	bookingv1 "github.com/codebind-luna/booking-service/gen/go/booking/v1"
)

const (
	successModifySeatMessage string = "successfully purchased ticket"
)

func (h *ticketServiceHandlers) ModifySeat(ctx context.Context, mr *bookingv1.ModifySeatRequest) (*bookingv1.ModifySeatResponse, error) {
	err := h.svc.ModifySeat(ctx, mr.GetEmail(), mr.GetSection().String(), int(mr.GetSeatNo()))

	res := &bookingv1.ModifySeatResponse{}

	if err != nil {
		return nil, err
	}

	res.Message = successModifySeatMessage
	res.Success = true

	return res, nil
}
