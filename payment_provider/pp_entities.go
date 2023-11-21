package paymentprovider

import "time"

type OptPayP func(*PaymentProviderOpts)

type PaymentProvider struct {
	PaymentProviderOpts
}

type PaymentProviderOpts struct {
	Name                    string                    `json:"name,omitempty"`
	PublicName              string                    `json:"public_name,omitempty"`
	Description             string                    `json:"description,omitempty"`
	LogoUrls                LogoURLS                  `json:"logo_urls,omitempty"`
	ConfigurationURL        string                    `json:"configuration_url,omitempty"`
	SupportURL              string                    `json:"support_url,omitempty"`
	RatesURL                string                    `json:"rates_url,omitempty"`
	CheckoutJsURL           string                    `json:"checkout_js_url,omitempty"`
	SupportedCurrencies     []string                  `json:"supported_currencies,omitempty"`
	SupportedPaymentMethods []SupportedPaymentMethods `json:"supported_payment_methods,omitempty"`
	Rates                   []Rate                    `json:"rates,omitempty"`
	CheckoutPaymentOptions  []CheckoutPaymentOption   `json:"checkout_payment_options,omitempty"`
	Features                []string                  `json:"features,omitempty"`
	Enabled                 bool                      `json:"enabled,omitempty"`
	Authentication          Authentication            `json:"authentication,omitempty"`
}

type Authentication struct {
	Type            string    `json:"type,omitempty"`
	AccessToken     string    `json:"access_token,omitempty"`
	ClientID        string    `json:"client_id,omitempty"`
	ClientSecret    string    `json:"client_secret,omitempty"`
	ExpiresAt       time.Time `json:"expires_at,omitempty"`
	RefreshToken    string    `json:"refresh_token,omitempty"`
	RefreshTokenURL string    `json:"refresh_token_url,omitempty"`
}

type LogoURLS struct {
	S400X120 string `json:"400x120,omitempty"`
	S160X100 string `json:"160x100,omitempty"`
}

type SupportedPaymentMethods struct {
	PaymentMethodType string      `json:"payment_method_type,omitempty"`
	PaymentMethods    []string    `json:"payment_methods,omitempty"`
	Installments      Installment `json:"installments,omitempty"`
}

type Installment struct {
	MinInstallmentValue []MinInstallmentValue `json:"min_installment_value,omitempty"`
	Specification       []Specification       `json:"specification,omitempty"`
}

type MinInstallmentValue struct {
	Currency string `json:"currency,omitempty"`
	Value    string `json:"value,omitempty"`
}

type Specification struct {
	Installments int      `json:"installments,omitempty"`
	InterestRate string   `json:"interest_rate,omitempty"`
	AppliesTo    []string `json:"applies_to,omitempty"`
}

type CheckoutPaymentOption struct {
	ID                          string   `json:"id,omitempty"`
	Name                        string   `json:"name,omitempty"`
	Description                 string   `json:"description,omitempty"`
	LogoURL                     string   `json:"logo_url,omitempty"`
	SupportedBillingCountries   []string `json:"supported_billing_countries,omitempty"`
	SupportedPaymentMethodTypes []string `json:"supported_payment_method_types,omitempty"`
	IntegrationType             string   `json:"integration_type,omitempty"`
}

type Rate struct {
	PaymentMethodType string            `json:"payment_method_type,omitempty"`
	RatesDefinition   []RatesDefinition `json:"rates_definition,omitempty"`
}

type RatesDefinition struct {
	PercentFee          string  `json:"percent_fee,omitempty"`
	FlatFee             FlatFee `json:"flat_fee,omitempty"`
	PlusTax             bool    `json:"plus_tax,omitempty"`
	DaysToWithdrawMoney int     `json:"days_to_withdraw_money,omitempty"`
}

type FlatFee struct {
	Value    string `json:"value,omitempty"`
	Currency string `json:"currency,omitempty"`
}
