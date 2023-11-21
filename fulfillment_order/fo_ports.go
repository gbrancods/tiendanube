// An order can now have multiple shipments. Each shipment is described in a new entity called
// Fulfillment Order.
package fulfillmentorder

import (
	"encoding/json"
	"time"
)

func defaultTracking() TrackingOpts {
	return TrackingOpts{}
}

// The fulfillment order tracking event input status.
func SetStatus(status string) OptTra {
	return func(t *TrackingOpts) {
		t.Status = status
	}
}

// The fulfillment order tracking event input description.
func SetDescription(description string) OptTra {
	return func(t *TrackingOpts) {
		t.Description = description
	}
}

// The fulfillment order tracking event input address as one liner address. Ex: St. Julio 123, Ciudad, Argentina.
func SetAddress(address string) OptTra {
	return func(t *TrackingOpts) {
		t.Address = address
	}
}

// The fulfillment order tracking event geolocation input.
func SetGeolocation(geolocation Geolocation) OptTra {
	return func(t *TrackingOpts) {
		t.Geolocation = geolocation
	}
}

// The fulfillment order tracking event input happened at the event. If null, the event was taken as now.
func SetHappenedAt(happenedAt time.Time) OptTra {
	return func(t *TrackingOpts) {
		t.HappenedAt = happenedAt
	}
}

// The fulfillment order tracking event input estimated delivery date time to arrive.
func SetEstimatedDeliveryAt(estimatedDeliveryAt time.Time) OptTra {
	return func(t *TrackingOpts) {
		t.EstimatedDeliveryAt = estimatedDeliveryAt
	}
}

// Initialize a tracking event instance
func New(opts ...OptTra) *Tracking {
	t := defaultTracking()
	for _, fn := range opts {
		fn(&t)
	}
	return &Tracking{
		TrackingOpts: t,
	}
}

// Retrive all Order Fulfillments from a specific Order.
func GetAll(orderID string) (r []FulfillmentOrder, err error) {
	b, err := getAllFulfillmentOrders(orderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Get a Fulfillment Order By Identifier
func GetByID(orderID, fulfillmentOrderID string) (r FulfillmentOrder, err error) {
	b, err := getFulfillmenteOrderByID(orderID, fulfillmentOrderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Get a Fulfillment Order By Identifier
func Delete(orderID, fulfillmentOrderID string) (err error) {
	err = deleteFulfillmentOrder(orderID, fulfillmentOrderID)
	return
}

// Update Fulfillment Order Status, Tracking Info, Destination, Recipient, Shipping, Assigned Location
func (fp FulfillmentOrder) Patch(orderID, fulfillmentOrderID string) (r FulfillmentOrder, err error) {
	b, err := json.Marshal(fp)
	if err != nil {
		return
	}
	br, err := patchFulfillmentOrder(b, orderID, fulfillmentOrderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Create Fulfillment Order Tracking Event
func (t Tracking) TrackingPost(orderID, fulfillmentOrderID string) (r FulfillmentOrder, err error) {
	b, err := json.Marshal(t)
	if err != nil {
		return
	}
	br, err := trackingPost(b, orderID, fulfillmentOrderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Update Fulfillment Order Tracking Event
func (t Tracking) TrackingPut(orderID, fulfillmentOrderID, trackingID string) (r FulfillmentOrder, err error) {
	b, err := json.Marshal(t)
	if err != nil {
		return
	}
	br, err := trackingPut(b, orderID, fulfillmentOrderID, trackingID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Delete Fulfillment Order Tracking Event
func DeleteTracking(orderID, fulfillmentOrderID, trackingID string) (err error) {
	err = trackingDelete(orderID, fulfillmentOrderID, trackingID)
	return
}

// Get All Fulfillment Order Tracking Events By Fulfillment Order
func GetAllTrackings(orderID, fulfillmentOrderID string) (f FulfillmentOrder, err error) {
	r, err := getAllTrackings(orderID, fulfillmentOrderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &f)
	return
}

// Get Fulfillment Order Tracking Event
func GetTrackingByID(orderID, fulfillmentOrderID, trackingID string) (r FulfillmentOrder, err error) {
	b, err := getTrackingByID(orderID, fulfillmentOrderID, trackingID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}
