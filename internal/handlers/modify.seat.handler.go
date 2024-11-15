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

func (h *ticketServiceHandlers) ModifySeat(ctx context.Context, mr *bookingv1.ModifySeatRequest) (*bookingv1.ModifySeatResponse, error) {
	err := h.svc.ModifySeat(ctx, mr.GetEmail(), mr.GetSection().String(), int(mr.GetSeatNo()))

	if err != nil {
		if errors.Is(err, exceptions.ErrNoBookingFoundForUser) {
			return nil, status.New(codes.NotFound, "No booking found").Err()
		}

		if errors.Is(err, exceptions.ErrSeatNotFound) {
			return nil, status.New(codes.InvalidArgument, "Requested seat is invalid").Err()
		}

		if errors.Is(err, exceptions.ErrSeatNotAvailableAnymore) {
			return nil, status.New(codes.FailedPrecondition, "Requested seat is not available anymore").Err()
		}

		return nil, status.New(codes.Unknown, "some unknow error happened").Err()
	}

	res := &bookingv1.ModifySeatResponse{}

	res.Message = constants.SuccessModifySeatMessage
	res.Success = true

	return res, nil
}
