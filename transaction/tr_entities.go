package transaction

import "time"

type OptTra func(*TransactionOpts)

type Transaction struct {
	TransactionOpts
}

type TransactionOpts struct {
	ID                string      `json:"id,omitempty"`
	PaymentProviderID string      `json:"payment_provider_id,omitempty"`
	CreatedAt         string      `json:"created_at,omitempty"`
	Status            string      `json:"status,omitempty"`
	RefundedAmount    interface{} `json:"refunded_amount,omitempty"` // RefundedAmount
	CapturedAmount    interface{} `json:"captured_amount,omitempty"` // CapturedAmount
	PaymentMethod     interface{} `json:"payment_method,omitempty"`  // PaymentMethod
	Info              interface{} `json:"info,omitempty"`            // Info
	FirstEvent        interface{} `json:"first_event,omitempty"`     // FirstEvent
	Events            interface{} `json:"events,omitempty"`          // []Event
	AuthorizedAmount  interface{} `json:"authorized_amount,omitempty"`
	VoidedAmount      interface{} `json:"voided_amount,omitempty"`
	FailureCode       interface{} `json:"failure_code,omitempty"`
}

type FirstEvent struct {
	Amount     Amount    `json:"amount,omitempty"` // Amount
	Type       string    `json:"type,omitempty"`
	Status     string    `json:"status,omitempty"`
	HappenedAt time.Time `json:"happened_at,omitempty"`
}

type CapturedAmount struct {
	Value    string `json:"value,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type RefundedAmount struct {
	Value    string `json:"value,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type PaymentMethod struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

type Info struct {
	Card                  Card         `json:"card,omitempty"`         // Card
	Installments          Installments `json:"installments,omitempty"` // Installments
	ExternalID            string       `json:"external_id,omitempty"`
	ExternalURL           string       `json:"external_url,omitempty"`
	RefundURL             string       `json:"refund_url,omitempty"`
	SupportsPartialRefund bool         `json:"supports_partial_refund,omitempty"`
	IP                    string       `json:"ip,omitempty"`
}

type Installments struct {
	Quantity int    `json:"quantity,omitempty"`
	Interest string `json:"interest,omitempty"`
}

type Card struct {
	Brand           string `json:"brand,omitempty"`
	ExpirationMonth int    `json:"expiration_month,omitempty"`
	ExpirationYear  int    `json:"expiration_year,omitempty"`
	FirstDigits     string `json:"first_digits,omitempty"`
	LastDigits      string `json:"last_digits,omitempty"`
	MaskedNumber    string `json:"masked_number,omitempty"`
	Name            string `json:"name,omitempty"`
}

type Event struct {
	TransactionID string      `json:"transaction_id,omitempty"`
	Amount        Amount      `json:"amount,omitempty"` // Amount
	Type          string      `json:"type,omitempty"`
	Status        string      `json:"status,omitempty"`
	FailureCode   interface{} `json:"failure_code,omitempty"`
	CreatedAt     time.Time   `json:"created_at,omitempty"`
	HappenedAt    time.Time   `json:"happened_at,omitempty"`
	ExpiresAt     interface{} `json:"expires_at,omitempty"`
}

type Amount struct {
	Value    string `json:"value,omitempty"`
	Currency string `json:"currency,omitempty"`
}
