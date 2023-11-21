package draftorder

type OptDrO func(*DraftOrderOpts)

type DraftOrder struct {
	DraftOrderOpts
}

type DraftOrderOpts struct {
	ID                       int                 `json:"id,omitempty"`
	Token                    string              `json:"token,omitempty"`
	StoreID                  any                 `json:"store_id,omitempty"` // TODO - divergence, to create a need for a int, to get need for a string
	ContactEmail             string              `json:"contact_email,omitempty"`
	ContactName              string              `json:"contact_name,omitempty"`
	ContactLastName          string              `json:"contact_lastname,omitempty"`
	ContactPhone             string              `json:"contact_phone,omitempty"`
	ContactIdentification    string              `json:"contact_identification,omitempty"`
	ShippingMinDays          any                 `json:"shipping_min_days,omitempty"`
	ShippingMaxDays          any                 `json:"shipping_max_days,omitempty"`
	BillingName              any                 `json:"billing_name,omitempty"`
	BillingPhone             any                 `json:"billing_phone,omitempty"`
	BillingAddress           any                 `json:"billing_address,omitempty"`
	BillingNumber            any                 `json:"billing_number,omitempty"`
	BillingFloor             any                 `json:"billing_floor,omitempty"`
	BillingLocality          any                 `json:"billing_locality,omitempty"`
	BillingZipcode           any                 `json:"billing_zipcode,omitempty"`
	BillingCity              any                 `json:"billing_city,omitempty"`
	BillingProvince          any                 `json:"billing_province,omitempty"`
	BillingCountry           any                 `json:"billing_country,omitempty"`
	ShippingCostOwner        string              `json:"shipping_cost_owner,omitempty"`
	ShippingCostCustomer     string              `json:"shipping_cost_customer,omitempty"`
	Coupon                   []any               `json:"coupon,omitempty"`
	PromotionalDiscount      PromotionalDiscount `json:"promotional_discount,omitempty"`
	Subtotal                 string              `json:"subtotal,omitempty"`
	Discount                 string              `json:"discount,omitempty"`
	DiscountCoupon           string              `json:"discount_coupon,omitempty"`
	DiscountGateway          string              `json:"discount_gateway,omitempty"`
	Total                    string              `json:"total,omitempty"`
	TotalUsd                 string              `json:"total_usd,omitempty"`
	CheckoutEnabled          bool                `json:"checkout_enabled,omitempty"`
	Weight                   string              `json:"weight,omitempty"`
	Currency                 string              `json:"currency,omitempty"`
	Language                 string              `json:"language,omitempty"`
	Gateway                  string              `json:"gateway,omitempty"`
	GatewayID                any                 `json:"gateway_id,omitempty"`
	GatewayName              string              `json:"gateway_name,omitempty"`
	Shipping                 string              `json:"shipping,omitempty"`
	ShippingOption           string              `json:"shipping_option,omitempty"`
	ShippingOptionCode       string              `json:"shipping_option_code,omitempty"`
	ShippingOptionReference  any                 `json:"shipping_option_reference,omitempty"`
	ShippingPickupDetails    any                 `json:"shipping_pickup_details,omitempty"`
	ShippingTrackingNumber   any                 `json:"shipping_tracking_number,omitempty"`
	ShippingTrackingURL      any                 `json:"shipping_tracking_url,omitempty"`
	ShippingStoreBranchName  any                 `json:"shipping_store_branch_name,omitempty"`
	ShippingStoreBranchExtra any                 `json:"shipping_store_branch_extra,omitempty"`
	ShippingPickupType       string              `json:"shipping_pickup_type,omitempty"`
	ShippingSuboption        []any               `json:"shipping_suboption,omitempty"`
	Extra                    struct{}            `json:"extra,omitempty"`
	Storefront               string              `json:"storefront,omitempty"`
	Note                     any                 `json:"note,omitempty"`
	CreatedAt                string              `json:"created_at,omitempty"`
	UpdatedAt                string              `json:"updated_at,omitempty"`
	CompletedAt              CompletedAt         `json:"completed_at,omitempty"`
	NextAction               string              `json:"next_action,omitempty"`
	PaymentDetails           PaymentDetails      `json:"payment_details,omitempty"`
	Attributes               []any               `json:"attributes,omitempty"`
	Customer                 Customer            `json:"customer,omitempty"`
	Products                 []Product           `json:"products,omitempty"`
	Fulfillments             any                 `json:"fulfillments,omitempty"`
	Number                   int                 `json:"number,omitempty"`
	CancelReason             any                 `json:"cancel_reason,omitempty"`
	OwnerNote                string              `json:"owner_note,omitempty"`
	CancelledAt              any                 `json:"cancelled_at,omitempty"`
	ClosedAt                 any                 `json:"closed_at,omitempty"`
	ReadAt                   any                 `json:"read_at,omitempty"`
	Status                   string              `json:"status,omitempty"`
	PaymentStatus            string              `json:"payment_status,omitempty"`
	GatewayLink              any                 `json:"gateway_link,omitempty"`
	HasShippableProducts     bool                `json:"has_shippable_products,omitempty"`
	ShippingCarrierName      any                 `json:"shipping_carrier_name,omitempty"`
	ShippingAddress          ShippingAddress     `json:"shipping_address,omitempty"`
	ShippingStatus           string              `json:"shipping_status,omitempty"`
	ShippedAt                any                 `json:"shipped_at,omitempty"`
	PaidAt                   any                 `json:"paid_at,omitempty"`
	LandingURL               any                 `json:"landing_url,omitempty"`
	ClientDetails            ClientDetails       `json:"client_details,omitempty"`
	AppID                    any                 `json:"app_id,omitempty"`
	CheckoutURL              string              `json:"checkout_url,omitempty"`
	//SaleChannel              string      `json:"sale_channel,omitempty"` //TODO not working
}

type PromotionalDiscount struct {
	ID                  any    `json:"id,omitempty"`
	StoreID             int    `json:"store_id,omitempty"`
	OrderID             any    `json:"order_id,omitempty"` // TODO - divergence, to create a need for a int, to get need for a string
	CreatedAt           string `json:"created_at,omitempty"`
	TotalDiscountAmount string `json:"total_discount_amount,omitempty"`
	Contents            []any  `json:"contents,omitempty"`
	PromotionsApplied   []any  `json:"promotions_applied,omitempty"`
}

type CompletedAt struct {
	Date         string `json:"date,omitempty"`
	TimezoneType int    `json:"timezone_type,omitempty"`
	Timezone     string `json:"timezone,omitempty"`
}

type PaymentDetails struct {
	Method            any `json:"method,omitempty"`
	CreditCardCompany any `json:"credit_card_company,omitempty"`
	Installments      any `json:"installments,omitempty"` // TODO - DIVERGENT!
}

type Customer struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Email          string `json:"email,omitempty"`
	Identification string `json:"identification,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Note           any    `json:"note,omitempty"`
	DefaultAddress struct {
		Address   string `json:"address,omitempty"`
		City      string `json:"city,omitempty"`
		Country   string `json:"country,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
		Default   bool   `json:"default,omitempty"`
		Floor     string `json:"floor,omitempty"`
		ID        int    `json:"id,omitempty"`
		Locality  string `json:"locality,omitempty"`
		Name      string `json:"name,omitempty"`
		Number    string `json:"number,omitempty"`
		Phone     string `json:"phone,omitempty"`
		Province  string `json:"province,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
		Zipcode   string `json:"zipcode,omitempty"`
	} `json:"default_address,omitempty"`
	Addresses []struct {
		Address   string `json:"address,omitempty"`
		City      string `json:"city,omitempty"`
		Country   string `json:"country,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
		Default   bool   `json:"default,omitempty"`
		Floor     any    `json:"floor,omitempty"`
		ID        int    `json:"id,omitempty"`
		Locality  string `json:"locality,omitempty"`
		Name      any    `json:"name,omitempty"`
		Number    string `json:"number,omitempty"`
		Phone     string `json:"phone,omitempty"`
		Province  string `json:"province,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
		Zipcode   string `json:"zipcode,omitempty"`
	} `json:"addresses,omitempty"`
	BillingName     any `json:"billing_name,omitempty"`
	BillingPhone    any `json:"billing_phone,omitempty"`
	BillingAddress  any `json:"billing_address,omitempty"`
	BillingNumber   any `json:"billing_number,omitempty"`
	BillingFloor    any `json:"billing_floor,omitempty"`
	BillingLocality any `json:"billing_locality,omitempty"`
	BillingZipcode  any `json:"billing_zipcode,omitempty"`
	BillingCity     any `json:"billing_city,omitempty"`
	BillingProvince any `json:"billing_province,omitempty"`
	BillingCountry  any `json:"billing_country,omitempty"`
	Extra           struct {
	} `json:"extra,omitempty"`
	TotalSpent                string `json:"total_spent,omitempty"`
	TotalSpentCurrency        string `json:"total_spent_currency,omitempty"`
	LastOrderID               int    `json:"last_order_id,omitempty"`
	Active                    bool   `json:"active,omitempty"`
	FirstInteraction          string `json:"first_interaction,omitempty"`
	CreatedAt                 string `json:"created_at,omitempty"`
	UpdatedAt                 string `json:"updated_at,omitempty"`
	AcceptsMarketing          bool   `json:"accepts_marketing,omitempty"`
	AcceptsMarketingUpdatedAt string `json:"accepts_marketing_updated_at,omitempty"`
}

type Product struct {
	ID             int    `json:"id,omitempty"`
	Depth          string `json:"depth,omitempty"`
	Height         string `json:"height,omitempty"`
	Name           string `json:"name,omitempty"`
	Price          string `json:"price,omitempty"`
	CompareAtPrice string `json:"compare_at_price,omitempty"`
	ProductID      int    `json:"product_id,omitempty"`
	Image          struct {
		ID        int    `json:"id,omitempty"`
		ProductID int    `json:"product_id,omitempty"`
		Src       string `json:"src,omitempty"`
		Position  int    `json:"position,omitempty"`
		Alt       []any  `json:"alt,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
	} `json:"image,omitempty"`
	Quantity      int    `json:"quantity,omitempty"`
	FreeShipping  bool   `json:"free_shipping,omitempty"`
	Weight        string `json:"weight,omitempty"`
	Width         string `json:"width,omitempty"`
	VariantID     int    `json:"variant_id,omitempty"`
	VariantValues []any  `json:"variant_values,omitempty"`
	Properties    []any  `json:"properties,omitempty"`
	Sku           any    `json:"sku,omitempty"`
	Barcode       any    `json:"barcode,omitempty"`
	Cost          any    `json:"cost,omitempty"`
}

type ShippingAddress struct {
	Address   string `json:"address,omitempty"`
	City      string `json:"city,omitempty"`
	Country   string `json:"country,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Default   bool   `json:"default,omitempty"`
	Floor     string `json:"floor,omitempty"`
	ID        int    `json:"id,omitempty"`
	Locality  string `json:"locality,omitempty"`
	Name      string `json:"name,omitempty"`
	Number    string `json:"number,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Province  string `json:"province,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Zipcode   string `json:"zipcode,omitempty"`
	Customs   any    `json:"customs,omitempty"`
}

type ClientDetails struct {
	BrowserIP any `json:"browser_ip,omitempty"`
	UserAgent any `json:"user_agent,omitempty"`
}
