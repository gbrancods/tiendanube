package shippingcarrier

type OptShipC func(*ShippingCarrierOpts)

type ShippingCarrier struct {
	ShippingCarrierOpts
}

type ShippingCarrierOpts struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Active      bool   `json:"active"`
	CallbackURL string `json:"callback_url"`
	Types       string `json:"types"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type OptOpt func(*OptionsOpts)

type Options struct {
	OptionsOpts
}

type OptionsOpts struct {
	ID                int     `json:"id"`
	Code              string  `json:"code"`
	Name              string  `json:"name"`
	AdditionalDays    int     `json:"additional_days"`
	AdditionalCost    float64 `json:"additional_cost"`
	AllowFreeShipping bool    `json:"allow_free_shipping"`
	Active            bool    `json:"active"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}

type OptFulfi func(*FulfillmentOpts)

type Fulfillment struct {
	FulfillmentOpts
}

type FulfillmentOpts struct {
	ID                  int    `json:"id"`
	OrderID             int    `json:"order_id"`
	Status              string `json:"status"`
	Description         string `json:"description"`
	City                string `json:"city"`
	Province            string `json:"province"`
	Country             string `json:"country"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	HappenedAt          string `json:"happened_at"`
	EstimatedDeliveryAt string `json:"estimated_delivery_at"`
}
