package handlers

import (
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
