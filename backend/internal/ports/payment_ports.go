package ports

import (
	"antfarm/backend/internal/domain"
)

type PaymentService interface {
	CreateOrder(amount string) (*domain.Order, error)
	CheckOrder(orderID int64) (*domain.Order, error)
}
