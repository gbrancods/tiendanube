// A discount coupon is a way for a store to provide discounts
// for its customers. There are three type of coupons. Percentage,
// absolute and shipping. The percentage type indicates that the
// value is a percentage discount. The type absolute indicates that
// the value is an absolute amount of discount. And finally the type
// shipping indicates that the discount value is on the shipping.
package coupons

import (
	"encoding/json"
)

func defaultCoupon() CouponOpts {
	return CouponOpts{}
}

// String that identifies the coupon.
func SetCode(code string) OptCoup {
	return func(c *CouponOpts) {
		c.Code = code
	}
}

// Type of the coupon. Can take the following values: percentage, absolute or shipping.
func SetType(cType string) OptCoup {
	return func(c *CouponOpts) {
		c.Type = cType
	}
}

// Value of the discount
func SetValue(value string) OptCoup {
	return func(c *CouponOpts) {
		c.Value = value
	}
}

// Max number of times the coupon can be used.
func SetMaxUses(maxUses int) OptCoup {
	return func(c *CouponOpts) {
		c.MaxUses = maxUses
	}
}

// Indicates the minimun value of the bill for applying the discount
func SetMinPrice(minPrice int) OptCoup {
	return func(c *CouponOpts) {
		c.MinPrice = minPrice
	}
}

// Date from which the coupon is valid.
func SetStartDate(startDate string) OptCoup {
	return func(c *CouponOpts) {
		c.StartDate = startDate
	}
}

// Date of overdue of the coupon.
func SetEndDate(endDate string) OptCoup {
	return func(c *CouponOpts) {
		c.EndDate = endDate
	}
}

// Initialize a coupon instance
func New(opts ...OptCoup) *Coupon {

	dc := defaultCoupon()

	for _, fn := range opts {
		fn(&dc)
	}

	return &Coupon{
		CouponOpts: dc,
	}
}

// Create a new Coupon
func (c Coupon) Create() (r Coupon, err error) {
	b, err := json.Marshal(c)
	if err != nil {
		return
	}
	br, err := create(b)
	json.Unmarshal(br, &r)
	return
}

// Modify an existing coupon
func (c Coupon) Update(couponID int) (r Coupon, err error) {
	b, err := json.Marshal(c)
	if err != nil {
		return
	}
	br, err := update(b, couponID)
	json.Unmarshal(br, &r)
	return
}

// Receive a list of all Coupon
func GetAll() (r []Coupon, err error) {
	b, err := getAll()
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Receive a single Coupon
func GetByID(id int) (r Coupon, err error) {
	b, err := getByID(id)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Delete an existing coupon
func Delete(id int) (err error) {
	err = delete(id)
	return
}
