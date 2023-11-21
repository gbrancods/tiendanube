package abandonedcheckout

import (
	"github.com/gbrancods/tiendanube/coupons"
	fulfillmentorder "github.com/gbrancods/tiendanube/fulfillment_order"
	productvariant "github.com/gbrancods/tiendanube/products_variants"
)

type Coupon struct {
	CouponID int `json:"coupon_id,omitempty"`
}

type Checkout struct {
	ID                      int                          `json:"id,omitempty"`
	Token                   string                       `json:"token,omitempty"`
	StoreID                 int                          `json:"store_id,omitempty"`
	AbandonedCheckoutURL    string                       `json:"abandoned_checkout_url,omitempty"`
	ContactEmail            string                       `json:"contact_email,omitempty"`
	ContactName             string                       `json:"contact_name,omitempty"`
	ContactPhone            string                       `json:"contact_phone,omitempty"`
	ContactIdentification   string                       `json:"contact_identification,omitempty"`
	ShippingName            string                       `json:"shipping_name,omitempty"`
	ShippingPhone           string                       `json:"shipping_phone,omitempty"`
	ShippingAddress         string                       `json:"shipping_address,omitempty"`
	ShippingNumber          string                       `json:"shipping_number,omitempty"`
	ShippingFloor           string                       `json:"shipping_floor,omitempty"`
	ShippingLocality        string                       `json:"shipping_locality,omitempty"`
	ShippingZipcode         string                       `json:"shipping_zipcode,omitempty"`
	ShippingCity            string                       `json:"shipping_city,omitempty"`
	ShippingProvince        string                       `json:"shipping_province,omitempty"`
	ShippingCountry         string                       `json:"shipping_country,omitempty"`
	ShippingMinDays         int                          `json:"shipping_min_days,omitempty"`
	ShippingMaxDays         int                          `json:"shipping_max_days,omitempty"`
	BillingName             string                       `json:"billing_name,omitempty"`
	BillingPhone            string                       `json:"billing_phone,omitempty"`
	BillingAddress          string                       `json:"billing_address,omitempty"`
	BillingNumber           string                       `json:"billing_number,omitempty"`
	BillingFloor            string                       `json:"billing_floor,omitempty"`
	BillingLocality         string                       `json:"billing_locality,omitempty"`
	BillingZipcode          string                       `json:"billing_zipcode,omitempty"`
	BillingCity             string                       `json:"billing_city,omitempty"`
	BillingProvince         string                       `json:"billing_province,omitempty"`
	BillingCountry          string                       `json:"billing_country,omitempty"`
	ShippingCostOwner       string                       `json:"shipping_cost_owner,omitempty"`
	ShippingCostCustomer    string                       `json:"shipping_cost_customer,omitempty"`
	Coupon                  []coupons.CouponOpts         `json:"coupon,omitempty"`
	PromotionalDiscount     PromotionalDiscount          `json:"promotional_discount,omitempty"`
	Subtotal                string                       `json:"subtotal,omitempty"`
	Discount                string                       `json:"discount,omitempty"`
	DiscountCoupon          string                       `json:"discount_coupon,omitempty"`
	DiscountGateway         string                       `json:"discount_gateway,omitempty"`
	Total                   string                       `json:"total,omitempty"`
	TotalUsd                string                       `json:"total_usd,omitempty"`
	CheckoutEnabled         bool                         `json:"checkout_enabled,omitempty"`
	Weight                  string                       `json:"weight,omitempty"`
	Currency                string                       `json:"currency,omitempty"`
	Language                string                       `json:"language,omitempty"`
	Gateway                 string                       `json:"gateway,omitempty"`
	GatewayID               string                       `json:"gateway_id,omitempty"`
	Shipping                string                       `json:"shipping,omitempty"`
	ShippingOption          string                       `json:"shipping_option,omitempty"`
	ShippingOptionCode      string                       `json:"shipping_option_code,omitempty"`
	ShippingOptionReference string                       `json:"shipping_option_reference,omitempty"`
	ShippingPickupDetails   string                       `json:"shipping_pickup_details,omitempty"`
	ShippingTrackingNumber  string                       `json:"shipping_tracking_number,omitempty"`
	ShippingTrackingURL     string                       `json:"shipping_tracking_url,omitempty"`
	ShippingStoreBranchName string                       `json:"shipping_store_branch_name,omitempty"`
	ShippingPickupType      string                       `json:"shipping_pickup_type,omitempty"`
	ShippingSuboption       []string                     `json:"shipping_suboption,omitempty"`
	Extra                   interface{}                  `json:"extra,omitempty"`
	Storefront              string                       `json:"storefront,omitempty"`
	Note                    string                       `json:"note,omitempty"`
	CreatedAt               string                       `json:"created_at,omitempty"`
	UpdatedAt               string                       `json:"updated_at,omitempty"`
	CompletedAt             string                       `json:"completed_at,omitempty"`
	NextAction              string                       `json:"next_action,omitempty"`
	PaymentDetails          PaymentDetails               `json:"payment_details,omitempty"`
	Attributes              []fulfillmentorder.LineItem  `json:"attributes,omitempty"`
	Products                []productvariant.VariantOpts `json:"products,omitempty"`
}

type PromotionalDiscount struct {
	ID                  int        `json:"id,omitempty"`
	StoreID             int        `json:"store_id,omitempty"`
	OrderID             int        `json:"order_id,omitempty"`
	CreatedAt           string     `json:"created_at,omitempty"`
	TotalDiscountAmount string     `json:"total_discount_amount,omitempty"`
	Contents            []struct{} `json:"contents,omitempty"`
	PromotionsApplied   []struct{} `json:"promotions_applied,omitempty"`
}

type PaymentDetails struct {
	Method            string `json:"method,omitempty"`
	CreditCardCompany string `json:"credit_card_company,omitempty"`
	Installments      int    `json:"installments,omitempty"`
}
