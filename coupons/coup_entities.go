package coupons

import "github.com/gbrancods/tiendanube/categories"

type OptCoup func(*CouponOpts)

// A discount coupon is a way for a store to provide discounts
// for its customers. There are three type of coupons. Percentage,
// absolute and shipping. The percentage type indicates that the
// value is a percentage discount. The type absolute indicates that
// the value is an absolute amount of discount. And finally the type
// shipping indicates that the discount value is on the shipping.
type Coupon struct {
	CouponOpts
}

type CouponOpts struct {
	ID         int                   `json:"id,omitempty"`
	Code       string                `json:"code,omitempty"`
	Type       string                `json:"type,omitempty"`
	Value      string                `json:"value,omitempty"`
	Valid      bool                  `json:"valid,omitempty"`
	Used       int                   `json:"used,omitempty"`
	MaxUses    int                   `json:"max_uses,omitempty"`
	StartDate  string                `json:"start_date,omitempty"`
	EndDate    string                `json:"end_date,omitempty"`
	MinPrice   int                   `json:"min_price,omitempty"`
	Categories []categories.Category `json:"categories,omitempty"`
}
