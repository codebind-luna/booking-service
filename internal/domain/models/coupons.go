package models

type CouponCode string

const (
	DiscountFivePerCent CouponCode = "DISCOUNT_FIVE"
	DiscountTenPercent  CouponCode = "DISCOUNT_TEN"
)

func (c CouponCode) String() string {
	return string(c)
}

var CouponCodeMap = map[string]CouponCode{
	DiscountFivePerCent.String(): DiscountFivePerCent,
	DiscountTenPercent.String():  DiscountTenPercent,
}
