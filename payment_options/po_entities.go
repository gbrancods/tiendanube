package paymentoptions

type PaymentOption struct {
	ID                     string                  `json:"id,omitempty"`
	Name                   string                  `json:"name,omitempty"`
	LogoURL                string                  `json:"logo_url,omitempty"`
	CheckoutPaymentOptions []CheckoutPaymentOption `json:"checkout_payment_options,omitempty"`
}

type CheckoutPaymentOption struct {
	ID                          string   `json:"id,omitempty"`
	Name                        string   `json:"name,omitempty"`
	SupportedPaymentMethodTypes []string `json:"supported_payment_method_types,omitempty"`
	IntegrationType             string   `json:"integration_type,omitempty"`
}
