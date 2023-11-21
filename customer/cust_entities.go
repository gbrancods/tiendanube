package customer

type OptCus func(*CustomerOpts)

type Customer struct {
	CustomerOpts
}

type CustomerOpts struct {
	CreatedAt          string    `json:"created_at,omitempty"`
	Email              string    `json:"email,omitempty"`
	ID                 int       `json:"id,omitempty"`
	Identification     string    `json:"identification,omitempty"`
	LastOrderID        int       `json:"last_order_id,omitempty"`
	Name               string    `json:"name,omitempty"`
	Note               string    `json:"note,omitempty"`
	TotalSpent         string    `json:"total_spent,omitempty"`
	TotalSpentCurrency string    `json:"total_spent_currency,omitempty"`
	UpdatedAt          string    `json:"updated_at,omitempty"`
	BillingAddress     string    `json:"billing_address,omitempty"`
	BillingCity        string    `json:"billing_city,omitempty"`
	BillingCountry     string    `json:"billing_country,omitempty"`
	BillingFloor       string    `json:"billing_floor,omitempty"`
	BillingLocality    string    `json:"billing_locality,omitempty"`
	BillingNumber      string    `json:"billing_number,omitempty"`
	BillingPhone       string    `json:"billing_phone,omitempty"`
	BillingProvince    string    `json:"billing_province,omitempty"`
	BillingZipcode     string    `json:"billing_zipcode,omitempty"`
	SendEmailInvite    bool      `json:"send_email_invite,omitempty"`
	Password           string    `json:"password,omitempty"`
	Extra              Extra     `json:"extra,omitempty"`
	DefaultAddress     Address   `json:"default_address,omitempty"`
	Addresses          []Address `json:"addresses,omitempty"`
}

type Address struct {
	Address string `json:"address,omitempty"`
	City    string `json:"city,omitempty"`

	//ISO 3166-1 alpha 2 country code.
	//https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	Country  string `json:"country,omitempty"`
	Locality string `json:"locality,omitempty"`
	Number   string `json:"number,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Province string `json:"province,omitempty"`
	Zipcode  string `json:"zipcode,omitempty"`
}

type Extra struct {
	NumberOfChildren string `json:"number_of_children,omitempty"`
	Gender           string `json:"gender,omitempty"`
}
