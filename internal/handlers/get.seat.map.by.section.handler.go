package handlers

import (
	"context"
	"errors"

	bookingv1 "github.com/codebind-luna/booking-service/gen/go/booking/v1"
	"github.com/codebind-luna/booking-service/internal/domain/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	successViewMapMessage string = "successfully fetched seat map details"
)

func (h *ticketServiceHandlers) ViewSeatMap(ctx context.Context, vsm *bookingv1.ViewSeatMapRequest) (*bookingv1.ViewSeatMapResponse, error) {
	seats, err := h.svc.ViewSeatMap(ctx, vsm.GetSection().String())

	if err != nil {
		if errors.Is(err, models.ErrInvalidSection) {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		return nil, status.New(codes.Unknown, "some unknow error happened").Err()
	}

	s := []*bookingv1.Seat{}
	vsmResponse := &bookingv1.ViewSeatMapResponse{}

	vsmResponse.Success = true
	vsmResponse.Message = successViewMapMessage

	for _, v := range seats {
		sm := bookingv1.Seat{
			Section: v.Section().Section(),
			SeatNo:  int32(v.SeatNo()),
			Status:  v.Status().Status(),
		}

		if v.User() != nil {
			sm.Email = v.User().Email()
		}

		s = append(s, &sm)
	}
	vsmResponse.Seats = s
	return vsmResponse, nil
}
