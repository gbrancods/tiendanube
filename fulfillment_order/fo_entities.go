package fulfillmentorder

import "time"

type FulfillmentOrder struct {
	ID               string           `json:"id,omitempty"`
	Number           string           `json:"number,omitempty"`
	TotalQuantity    int              `json:"total_quantity,omitempty"`
	TotalWeight      float64          `json:"total_weight,omitempty"`
	TotalPrice       TotalPrice       `json:"total_price,omitempty"`
	AssignedLocation AssignedLocation `json:"assigned_location,omitempty"`
	LineItems        []LineItem       `json:"line_items,omitempty"`
	Recipient        Recipient        `json:"recipient,omitempty"`
	Shipping         Shipping         `json:"shipping,omitempty"`
	Destination      Address          `json:"destination,omitempty"`
	Status           string           `json:"status,omitempty"`
	StatusHistory    StatusHistory    `json:"status_history,omitempty"`
	TrackingInfo     TrackingInfo     `json:"tracking_info,omitempty"`
	TrackingEvents   []Tracking       `json:"tracking_events,omitempty"`
	FulfilledAt      time.Time        `json:"fulfilled_at,omitempty"`
	CreatedAt        time.Time        `json:"created_at,omitempty"`
	UpdatedAt        time.Time        `json:"updated_at,omitempty"`
}

type TotalPrice struct {
	Value    float64 `json:"value,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

type AssignedLocation struct {
	LocationID string  `json:"location_id,omitempty"`
	Name       string  `json:"name,omitempty"`
	Address    Address `json:"address,omitempty"`
}

type Recipient struct {
	Name              string `json:"name,omitempty"`
	Phone             string `json:"phone,omitempty"`
	Identifier        string `json:"identifier,omitempty"`
	AllowFreeShipping bool   `json:"allow_free_shipping,omitempty"`
}

// ------

type Shipping struct {
	Type            string        `json:"type,omitempty"`
	Carrier         Carrier       `json:"carrier,omitempty"`
	Option          Option        `json:"option,omitempty"`
	MerchantCost    MerchantCost  `json:"merchant_cost,omitempty"`
	ConsumerCost    ConsumerCost  `json:"consumer_cost,omitempty"`
	MinDeliveryDate time.Time     `json:"min_delivery_date,omitempty"`
	MaxDeliveryDate time.Time     `json:"max_delivery_date,omitempty"`
	PickupDetails   PickupDetails `json:"pickup_details,omitempty"`
	Extras          Extras        `json:"extras,omitempty"`
}

type Carrier struct {
	Name      string `json:"name,omitempty"`
	Code      string `json:"code,omitempty"`
	CarrierID string `json:"carrier_id,omitempty"`
}

type Option struct {
	Name      string `json:"name,omitempty"`
	Code      string `json:"code,omitempty"`
	Reference string `json:"reference,omitempty"`
}

type MerchantCost struct {
	Value    float64 `json:"value,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

type ConsumerCost struct {
	Value    float64 `json:"value,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

// -----

type PickupDetails struct {
	LocationID  string        `json:"location_id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Address     Address       `json:"address,omitempty"`
	PickupHours []PickupHours `json:"pickup_hours,omitempty"`
	Extras      Extras        `json:"extras,omitempty"`
}

type PickupHours struct {
	Day   string `json:"day,omitempty"`
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type Extras struct {
	FreeShippingInfo FreeShippingInfo `json:"free_shipping_info,omitempty"`
	PhoneRequired    bool             `json:"phone_required,omitempty"`
	IDRequired       bool             `json:"id_required,omitempty"`
	AcceptsCod       bool             `json:"accepts_cod,omitempty"`
	ShowTime         bool             `json:"show_time,omitempty"`
	Shippable        bool             `json:"shippable,omitempty"`
}

type FreeShippingInfo struct {
	FreeShippingID       string               `json:"free_shipping_id,omitempty"`
	ConsumerOriginalCost ConsumerOriginalCost `json:"consumer_original_cost,omitempty"`
}

type ConsumerOriginalCost struct {
	Value    float64 `json:"value,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

// -----

type Address struct {
	Zipcode        string   `json:"zipcode,omitempty"`
	Street         string   `json:"street,omitempty"`
	Number         string   `json:"number,omitempty"`
	Floor          string   `json:"floor,omitempty"`
	Locality       string   `json:"locality,omitempty"`
	City           string   `json:"city,omitempty"`
	Reference      string   `json:"reference,omitempty"`
	BetweenStreets string   `json:"between_streets,omitempty"`
	Province       Province `json:"province,omitempty"`
	Region         Region   `json:"region,omitempty"`
	Country        Country  `json:"country,omitempty"`
}

type Province struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

type Region struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

type Country struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

// -----

type StatusHistory struct {
	FromStatus string    `json:"from_status,omitempty"`
	ToStatus   string    `json:"to_status,omitempty"`
	HappenedAt time.Time `json:"happened_at,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

type LineItem struct {
	Quantity      int                   `json:"quantity,omitempty"`
	Variant       LineItemVariant       `json:"variant,omitempty"`
	Product       LineItemProduct       `json:"product,omitempty"`
	UnitPrice     LineItemUnitPrice     `json:"unit_price,omitempty"`
	UnitDimension LineItemUnitDimension `json:"unit_dimension,omitempty"`
	CreatedAt     time.Time             `json:"created_at,omitempty"`
	UpdatedAt     time.Time             `json:"updated_at,omitempty"`
}

type LineItemVariant struct {
	VariantID string `json:"variant_id,omitempty"`
}

type LineItemProduct struct {
	ProductID string `json:"product_id,omitempty"`
}

type LineItemUnitPrice struct {
	Value    float64 `json:"value,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

type LineItemUnitDimension struct {
	Weight float64 `json:"weight,omitempty"`
	Width  float64 `json:"width,omitempty"`
	Height float64 `json:"height,omitempty"`
	Depth  float64 `json:"depth,omitempty"`
}

type OptTra func(*TrackingOpts)

type Tracking struct {
	TrackingOpts
}

type TrackingOpts struct {
	ID                  string      `json:"id,omitempty"`
	Status              string      `json:"status,omitempty"`
	Description         string      `json:"description,omitempty"`
	Address             string      `json:"address,omitempty"`
	Geolocation         Geolocation `json:"geolocation,omitempty"`
	HappenedAt          time.Time   `json:"happened_at,omitempty"`
	EstimatedDeliveryAt time.Time   `json:"estimated_delivery_at,omitempty"`
	CreatedAt           time.Time   `json:"created_at,omitempty"`
	UpdatedAt           time.Time   `json:"updated_at,omitempty"`
}

type Geolocation struct {
	Longitude float64 `json:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
}

type TrackingInfo struct {
	URL  string `json:"url,omitempty"`
	Code string `json:"code,omitempty"`
}
