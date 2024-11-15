package handlers

import (
	"context"

	bookingv1 "github.com/codebind-luna/booking-service/gen/go/booking/v1"
)

const (
	successRemoveUserMessage string = "successfully removed user from train"
)

func (h *ticketServiceHandlers) RemoveUser(ctx context.Context, rr *bookingv1.RemoveUserRequest) (*bookingv1.RemoveUserResponse, error) {
	err := h.svc.RemoveUserfromTrain(ctx, rr.GetEmail())

	res := &bookingv1.RemoveUserResponse{}

	if err != nil {
		return nil, err
	}

	res.Message = successRemoveUserMessage
	res.Success = true

	return res, nil
}
