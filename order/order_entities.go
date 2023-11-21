package order

type OptOrd func(*OrderOpts)

type Order struct {
	OrderOpts
}

type OrderOpts struct {
	ID                      int                 `json:"id,omitempty"`
	Token                   string              `json:"token,omitempty"`
	StoreID                 int                 `json:"store_id,omitempty"`
	ContactEmail            string              `json:"contact_email,omitempty"`
	ContactName             string              `json:"contact_name,omitempty"`
	ContactPhone            string              `json:"contact_phone,omitempty"`
	ContactIdentification   interface{}         `json:"contact_identification,omitempty"`
	ShippingMinDays         int                 `json:"shipping_min_days,omitempty"`
	ShippingMaxDays         int                 `json:"shipping_max_days,omitempty"`
	BillingName             string              `json:"billing_name,omitempty"`
	BillingPhone            string              `json:"billing_phone,omitempty"`
	BillingAddress          string              `json:"billing_address,omitempty"`
	BillingNumber           string              `json:"billing_number,omitempty"`
	BillingFloor            string              `json:"billing_floor,omitempty"`
	BillingLocality         string              `json:"billing_locality,omitempty"`
	BillingZipcode          string              `json:"billing_zipcode,omitempty"`
	BillingCity             string              `json:"billing_city,omitempty"`
	BillingProvince         string              `json:"billing_province,omitempty"`
	BillingCountry          string              `json:"billing_country,omitempty"`
	ShippingCostOwner       string              `json:"shipping_cost_owner,omitempty"`
	ShippingCostCustomer    string              `json:"shipping_cost_customer,omitempty"`
	Coupon                  []interface{}       `json:"coupon,omitempty"`
	PromotionalDiscount     PromotionalDiscount `json:"promotional_discount,omitempty"`
	Subtotal                string              `json:"subtotal,omitempty"`
	Discount                string              `json:"discount,omitempty"`
	DiscountCoupon          string              `json:"discount_coupon,omitempty"`
	DiscountGateway         string              `json:"discount_gateway,omitempty"`
	Total                   string              `json:"total,omitempty"`
	TotalUsd                string              `json:"total_usd,omitempty"`
	CheckoutEnabled         bool                `json:"checkout_enabled,omitempty"`
	Weight                  string              `json:"weight,omitempty"`
	Currency                string              `json:"currency,omitempty"`
	Language                string              `json:"language,omitempty"`
	Gateway                 string              `json:"gateway,omitempty"`
	GatewayID               interface{}         `json:"gateway_id,omitempty"`
	GatewayName             string              `json:"gateway_name,omitempty"`
	Shipping                string              `json:"shipping,omitempty"`
	ShippingOption          string              `json:"shipping_option,omitempty"`
	ShippingOptionCode      interface{}         `json:"shipping_option_code,omitempty"`
	ShippingOptionReference interface{}         `json:"shipping_option_reference,omitempty"`
	ShippingPickupDetails   interface{}         `json:"shipping_pickup_details,omitempty"`
	ShippingTrackingNumber  interface{}         `json:"shipping_tracking_number,omitempty"`
	ShippingTrackingURL     interface{}         `json:"shipping_tracking_url,omitempty"`
	ShippingStoreBranchName interface{}         `json:"shipping_store_branch_name,omitempty"`
	ShippingPickupType      string              `json:"shipping_pickup_type,omitempty"`
	ShippingSuboption       []interface{}       `json:"shipping_suboption,omitempty"`
	Extra                   struct{}            `json:"extra,omitempty"`
	Storefront              string              `json:"storefront,omitempty"`
	Note                    interface{}         `json:"note,omitempty"`
	CreatedAt               string              `json:"created_at,omitempty"`
	UpdatedAt               string              `json:"updated_at,omitempty"`
	CompletedAt             CompletedAt         `json:"completed_at,omitempty"`
	NextAction              string              `json:"next_action,omitempty"`
	PaymentDetails          PaymentDetails      `json:"payment_details,omitempty"`
	Attributes              []interface{}       `json:"attributes,omitempty"`
	Customer                Customer            `json:"customer,omitempty"`
	Products                []Product           `json:"products,omitempty"`
	Number                  int                 `json:"number,omitempty"`
	CancelReason            interface{}         `json:"cancel_reason,omitempty"`
	OwnerNote               interface{}         `json:"owner_note,omitempty"`
	CancelledAt             interface{}         `json:"cancelled_at,omitempty"`
	ClosedAt                interface{}         `json:"closed_at,omitempty"`
	ReadAt                  interface{}         `json:"read_at,omitempty"`
	Status                  string              `json:"status,omitempty"`
	PaymentStatus           string              `json:"payment_status,omitempty"`
	GatewayLink             interface{}         `json:"gateway_link,omitempty"`
	ShippingCarrierName     interface{}         `json:"shipping_carrier_name,omitempty"`
	ShippingAddress         ShippingAddress     `json:"shipping_address,omitempty"`
	ShippingStatus          string              `json:"shipping_status,omitempty"`
	ShippedAt               interface{}         `json:"shipped_at,omitempty"`
	PaidAt                  interface{}         `json:"paid_at,omitempty"`
	LandingURL              interface{}         `json:"landing_url,omitempty"`
	ClientDetails           ClientDetails       `json:"client_details,omitempty"`
	AppID                   int                 `json:"app_id,omitempty"`
}

type PromotionalDiscount struct {
	ID                  interface{}   `json:"id,omitempty"`
	StoreID             int           `json:"store_id,omitempty"`
	OrderID             int           `json:"order_id,omitempty"`
	CreatedAt           string        `json:"created_at,omitempty"`
	TotalDiscountAmount string        `json:"total_discount_amount,omitempty"`
	Contents            []interface{} `json:"contents,omitempty"`
	PromotionsApplied   []interface{} `json:"promotions_applied,omitempty"`
}

type CompletedAt struct {
	Date         string `json:"date,omitempty"`
	TimezoneType int    `json:"timezone_type,omitempty"`
	Timezone     string `json:"timezone,omitempty"`
}

type PaymentDetails struct {
	Method            interface{} `json:"method,omitempty"`
	CreditCardCompany interface{} `json:"credit_card_company,omitempty"`
	Installments      int         `json:"installments,omitempty"`
}

type Customer struct {
	ID             int         `json:"id,omitempty"`
	Name           string      `json:"name,omitempty"`
	Email          string      `json:"email,omitempty"`
	Identification interface{} `json:"identification,omitempty"`
	Phone          string      `json:"phone,omitempty"`
	Note           interface{} `json:"note,omitempty"`
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
		Address   string      `json:"address,omitempty"`
		City      string      `json:"city,omitempty"`
		Country   interface{} `json:"country,omitempty"`
		CreatedAt string      `json:"created_at,omitempty"`
		Default   bool        `json:"default,omitempty"`
		Floor     string      `json:"floor,omitempty"`
		ID        int         `json:"id,omitempty"`
		Locality  string      `json:"locality,omitempty"`
		Name      string      `json:"name,omitempty"`
		Number    string      `json:"number,omitempty"`
		Phone     string      `json:"phone,omitempty"`
		Province  string      `json:"province,omitempty"`
		UpdatedAt string      `json:"updated_at,omitempty"`
		Zipcode   string      `json:"zipcode,omitempty"`
	} `json:"addresses,omitempty"`
	BillingName        string   `json:"billing_name,omitempty"`
	BillingPhone       string   `json:"billing_phone,omitempty"`
	BillingAddress     string   `json:"billing_address,omitempty"`
	BillingNumber      string   `json:"billing_number,omitempty"`
	BillingFloor       string   `json:"billing_floor,omitempty"`
	BillingLocality    string   `json:"billing_locality,omitempty"`
	BillingZipcode     string   `json:"billing_zipcode,omitempty"`
	BillingCity        string   `json:"billing_city,omitempty"`
	BillingProvince    string   `json:"billing_province,omitempty"`
	BillingCountry     string   `json:"billing_country,omitempty"`
	Extra              struct{} `json:"extra,omitempty"`
	TotalSpent         string   `json:"total_spent,omitempty"`
	TotalSpentCurrency string   `json:"total_spent_currency,omitempty"`
	LastOrderID        int      `json:"last_order_id,omitempty"`
	Active             bool     `json:"active,omitempty"`
	FirstInteraction   string   `json:"first_interaction,omitempty"`
	CreatedAt          string   `json:"created_at,omitempty"`
	UpdatedAt          string   `json:"updated_at,omitempty"`
}

type Product struct {
	ID             int           `json:"id,omitempty"`
	Depth          string        `json:"depth,omitempty"`
	Height         string        `json:"height,omitempty"`
	Name           string        `json:"name,omitempty"`
	Price          string        `json:"price,omitempty"`
	CompareAtPrice string        `json:"compare_at_price,omitempty"`
	ProductID      int           `json:"product_id,omitempty"`
	Image          Image         `json:"image,omitempty"`
	Quantity       int           `json:"quantity,omitempty"`
	FreeShipping   bool          `json:"free_shipping,omitempty"`
	Weight         string        `json:"weight,omitempty"`
	Width          string        `json:"width,omitempty"`
	VariantID      int           `json:"variant_id,omitempty"`
	VariantValues  []interface{} `json:"variant_values,omitempty"`
	Properties     []interface{} `json:"properties,omitempty"`
	Sku            interface{}   `json:"sku,omitempty"`
	Barcode        interface{}   `json:"barcode,omitempty"`
}

type Image struct {
	ID        int           `json:"id,omitempty"`
	ProductID int           `json:"product_id,omitempty"`
	Src       string        `json:"src,omitempty"`
	Position  int           `json:"position,omitempty"`
	Alt       []interface{} `json:"alt,omitempty"`
	CreatedAt string        `json:"created_at,omitempty"`
	UpdatedAt string        `json:"updated_at,omitempty"`
}

type ShippingAddress struct {
	Address   string      `json:"address,omitempty"`
	City      string      `json:"city,omitempty"`
	Country   string      `json:"country,omitempty"`
	CreatedAt string      `json:"created_at,omitempty"`
	Default   bool        `json:"default,omitempty"`
	Floor     string      `json:"floor,omitempty"`
	ID        int         `json:"id,omitempty"`
	Locality  string      `json:"locality,omitempty"`
	Name      string      `json:"name,omitempty"`
	Number    string      `json:"number,omitempty"`
	Phone     string      `json:"phone,omitempty"`
	Province  string      `json:"province,omitempty"`
	UpdatedAt string      `json:"updated_at,omitempty"`
	Zipcode   string      `json:"zipcode,omitempty"`
	Customs   interface{} `json:"customs,omitempty"`
}

type ClientDetails struct {
	BrowserIP interface{} `json:"browser_ip,omitempty"`
	UserAgent interface{} `json:"user_agent,omitempty"`
}
