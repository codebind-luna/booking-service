package coupon

import (
	"context"

	"github.com/codebind-luna/booking-service/internal/domain/models"
)

type tenPercentDiscountCoupon struct {
	code          models.CouponCode
	discount      float64
	isAlwaysValid bool
}

func NewTenPercentDiscountCoupon(code models.CouponCode) *tenPercentDiscountCoupon {
	return &tenPercentDiscountCoupon{
		isAlwaysValid: true,
		discount:      10,
		code:          code,
	}
}

func (tc *tenPercentDiscountCoupon) Valid(ctx context.Context) bool {
	return tc.isAlwaysValid
}

func (tc *tenPercentDiscountCoupon) ApplyCoupon(ctx context.Context, price float64) float64 {
	return price * (1 - tc.discount/100)
}
