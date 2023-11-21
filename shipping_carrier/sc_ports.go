// A shipping carrier (or shipping company) service provides real-time shipping rates for merchants.
//
// Using these endpoints, you can add a shipping carrier to a store and provide shipping rates at checkout.
package shippingcarrier

import (
	"encoding/json"
)

func defaultShippingCarrier() ShippingCarrierOpts {
	return ShippingCarrierOpts{}
}

// The name of the Shipping Carrier as seen by both merchants and buyers.
func SetName(name string) OptShipC {
	return func(v *ShippingCarrierOpts) {
		v.Name = name
	}
}

// The URL endpoint that we need to retrieve shipping rates. Must be HTTPS.
func SetCallbackURL(url string) OptShipC {
	return func(v *ShippingCarrierOpts) {
		v.CallbackURL = url
	}
}

// The supported shipping types, can be one or both ship or pickup.
func SetTypes(types string) OptShipC {
	return func(v *ShippingCarrierOpts) {
		v.Types = types
	}
}

// Whether this Shipping Carrier is active.
func SetActive(active bool) OptShipC {
	return func(v *ShippingCarrierOpts) {
		v.Active = active
	}
}

// Initialize a shipping carrier instance
func New(opts ...OptShipC) *ShippingCarrier {
	dp := defaultShippingCarrier()
	for _, fn := range opts {
		fn(&dp)
	}
	return &ShippingCarrier{
		ShippingCarrierOpts: dp,
	}
}

// Create a new Shipping Carrier
func (s ShippingCarrier) Create() (r ShippingCarrier, err error) {
	b, err := json.Marshal(s)
	if err != nil {
		return
	}
	br, err := create(b)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Receive a list of all Shipping Carriers.
func GetAll() (r []ShippingCarrier, err error) {
	b, err := getAll()
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Receive a single Shipping Carrier
func GetByID(shippingCarrierID int) (r ShippingCarrier, err error) {
	b, err := getByID(shippingCarrierID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Modify an existing Shipping Carrier
func (s ShippingCarrier) Update(shippingCarrierID int) (r ShippingCarrier, err error) {
	b, err := json.Marshal(s)
	if err != nil {
		return
	}
	br, err := update(b, shippingCarrierID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Remove a Shipping Carrier
func Delete(shippingCarrierID int) (err error) {
	err = delete(shippingCarrierID)
	return
}

// |-----------|
// |  OPTIONS  |
// |-----------|

func defaultOptions() OptionsOpts {
	return OptionsOpts{}
}

// A unique code associated with the Shipping Carrier Option. The carrier option code is used for matching
// against the shipping rates. This allows merchant to add additional day and cost settings to the retrieved carrier options.
func SetOptionsCode(code string) OptOpt {
	return func(v *OptionsOpts) {
		v.Code = code
	}
}

// The name of the Shipping Carrier Option as seen by both merchants and buyers.
func SetOptionsName(name string) OptOpt {
	return func(v *OptionsOpts) {
		v.Name = name
	}
}

// The additional days configurable value that will be added to the option's estimated delivery time.
func SetOptionsAdditionalDays(additionalDays int) OptOpt {
	return func(v *OptionsOpts) {
		v.AdditionalDays = additionalDays
	}
}

// The additional cost configurable value that will be added to the option's consumer price.
func SetOptionsAdditionalCost(additionalCost float64) OptOpt {
	return func(v *OptionsOpts) {
		v.AdditionalCost = additionalCost
	}
}

// The configurable free shipping eligible parameter that specifies that an option allows free shipping.
func SetOptionsAllowFreeShipping(allowFreeShipping bool) OptOpt {
	return func(v *OptionsOpts) {
		v.AllowFreeShipping = allowFreeShipping
	}
}

// The avaiability status of the Shipping Carrier Option. Defaults to true
func SetOptionsActive(active bool) OptOpt {
	return func(v *OptionsOpts) {
		v.Active = active
	}
}

// Initialize a shipping option instance
func NewOption(opts ...OptOpt) *Options {
	dp := defaultOptions()
	for _, fn := range opts {
		fn(&dp)
	}
	return &Options{
		OptionsOpts: dp,
	}
}

// Create a new Shipping Carrier Option
func (o Options) CreateOption(carrierID int) (r Options, err error) {
	b, err := json.Marshal(o)
	if err != nil {
		return
	}
	br, err := createOption(b, carrierID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Receive a list of all Shipping Carriers Options.
func GetAllOptions(carrierID int) (r []Options, err error) {
	b, err := getOptions(carrierID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Receive a single Shipping Carrier Option.
func GetOptionByID(carrierID, optionID int) (r []Options, err error) {
	b, err := getOptionByID(carrierID, optionID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Modify an existing Shipping Carrier Option.
func (o Options) UpdateOption(carrierID, optionID int) (r []Options, err error) {
	b, err := json.Marshal(o)
	if err != nil {
		return
	}
	br, err := updateOption(b, carrierID, optionID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Remove a Shipping Carrier Option
func DeleteOption(carrierID, optionID int) (err error) {
	err = deleteOption(carrierID, optionID)
	return
}

// |----------------|
// |  FULFILLMENTS  |
// |----------------|

func defaultFulfillment() FulfillmentOpts {
	return FulfillmentOpts{}
}

// Fulfillment Event's status (see below for accepted values).
func SetFulfillmentStatus(status string) OptFulfi {
	return func(v *FulfillmentOpts) {
		v.Status = status
	}
}

// Message describing the status.
func SetFulfillmentDescription(description string) OptFulfi {
	return func(v *FulfillmentOpts) {
		v.Description = description
	}
}

// Fulfillment Event's city.
func SetFulfillmentCity(city string) OptFulfi {
	return func(v *FulfillmentOpts) {
		v.City = city
	}
}

// Fulfillment Event's province.
func SetFulfillmentProvince(province string) OptFulfi {
	return func(v *FulfillmentOpts) {
		v.Province = province
	}
}

// Fulfillment Event's country in ISO 3166-1 format.
func SetFulfillmentCountry(country string) OptFulfi {
	return func(v *FulfillmentOpts) {
		v.Country = country
	}
}

// Date when the Fulfillment Event occured was created in ISO 8601 format.
func SetFulfillmentHappenedAt(happenedAt string) OptFulfi {
	return func(v *FulfillmentOpts) {
		v.HappenedAt = happenedAt
	}
}

// Estimated date when the package will be delivered in ISO 8601 format.
func SetFulfillmentEstimatedDeliveryAt(estimatedDeliveryAt string) OptFulfi {
	return func(v *FulfillmentOpts) {
		v.EstimatedDeliveryAt = estimatedDeliveryAt
	}
}

// Initialize a shipping carrier fulfillment instance
func NewFulfillment(opts ...OptFulfi) *Fulfillment {
	dp := defaultFulfillment()
	for _, fn := range opts {
		fn(&dp)
	}
	return &Fulfillment{
		FulfillmentOpts: dp,
	}
}

// Receive a list of all Fulfillment Events for an order.
func GetOrderFulfillment(orderID int) (r Fulfillment, err error) {
	b, err := getAllFulfillments(orderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Receive a single Fulfillment Event
func GetOrderFulfillmentByID(orderID, fulfillmentID int) (r Fulfillment, err error) {
	b, err := getFulfillmentByID(orderID, fulfillmentID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Create a new Fulfillment Event
func (f Fulfillment) CreateOrderFulfillment(orderID int) (r Fulfillment, err error) {
	b, err := json.Marshal(f)
	if err != nil {
		return
	}
	br, err := createFulfillment(b, orderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Remove a Fulfillment Event
func DeleteOrderFulfillment(orderID, fulfillmentID int) (r Fulfillment, err error) {
	err = deleteFulfillment(orderID, fulfillmentID)
	return
}
