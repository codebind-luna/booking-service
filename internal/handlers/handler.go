package handlers

import (
	"context"

	bookingv1 "github.com/codebind-luna/booking-service/gen/go/booking/v1"
	"github.com/codebind-luna/booking-service/internal/domain"
)

type ticketServiceHandlers struct {
	bookingv1.UnimplementedTicketServiceServer
	svc domain.Service
}

func NewTicketService(svc domain.Service) *ticketServiceHandlers {
	return &ticketServiceHandlers{
		svc: svc,
	}
}

func (h *ticketServiceHandlers) RemoveUser(context.Context, *bookingv1.RemoveUserRequest) (*bookingv1.RemoveUserResponse, error) {
	panic("unimplemented")
}

func (h *ticketServiceHandlers) ModifySeat(context.Context, *bookingv1.ModifySeatRequest) (*bookingv1.ModifySeatResponse, error) {
	panic("unimplemented")
}
