// A Payment Provider can implement multiple payment options to be available at the checkout of a store.
// Payment options are the different alternatives through which a consumer can pay for their order in the store.
// These options can be integrated in the store checkout in two ways:
//
// - transparent integration: the payment through this option is processed within the checkout of the store;
//
// - redirect integration: by using this option, the consumer is redirected to an external checkout provided by the payment provider.
package paymentoptions

import (
	"encoding/json"
)

// Get the payment options available for the checkout of a store based on the Payment Providers activated by the merchant.
func PaymentOptions() (r []PaymentOption, err error) {
	br, err := getAll()
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}
