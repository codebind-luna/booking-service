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

func (h *ticketServiceHandlers) RemoveUser(ctx context.Context, rr *bookingv1.RemoveUserRequest) (*bookingv1.RemoveUserResponse, error) {
	err := h.svc.RemoveUserfromTrain(ctx, rr.GetEmail())

	if err != nil {
		if errors.Is(err, exceptions.ErrNoBookingFoundForUser) {
			return nil, status.New(codes.NotFound, "No booking found").Err()
		}

		return nil, status.New(codes.Unknown, "some unknow error happened").Err()
	}

	res := &bookingv1.RemoveUserResponse{}
	res.Message = constants.SuccessRemoveUserMessage
	res.Success = true

	return res, nil
}
