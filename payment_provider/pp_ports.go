// A Payment Provider, shorter name for Payments Services Provider, represents any entity which provides all the necessary
// resources and infrastructure for merchants and consumers to execute Transactions between them. This entities could be any of the following:
//
//	Aggregator
//	Acquirer
//	Gateway
//	Payments companies have many different and sometimes complex features which add value to the purchase experience, mainly providing multiple payments options and simpler checkout flows. They also provide merchants with tools to make better management of their Transactions as well as their incomes.
//
// In our platform, a Payment Provider is created for a specific store.
//
// Note: This endpoint is for the exclusive use of payment apps.
//
// Note: To create a Payments App you need to create an App in the Partners Portal and request our Partner Support Team (partners@nuvemshop.com.br/partners@tiendanube.com)
// to enable your app to access our Payments APIs.
package paymentprovider

import (
	"encoding/json"
)

func defaultPaymentProvider() PaymentProviderOpts {
	return PaymentProviderOpts{}
}

// Name to be displayed to merchants at the store admin tool.
func SetName(name string) OptPayP {
	return func(c *PaymentProviderOpts) {
		c.Name = name
	}
}

// [Optional] Name to be displayed to consumers at the storefront. If not specified, the same value as name is used.
func SetPublicName(publicName string) OptPayP {
	return func(c *PaymentProviderOpts) {
		c.PublicName = publicName
	}
}

// Short paragraph which provides merchants with a description of the Payment Provider.
func SetDescription(description string) OptPayP {
	return func(c *PaymentProviderOpts) {
		c.Description = description
	}
}

// Object containing key:value pair for each version of the logos for the frontend. Only supports HTTPS URLs. See Logos.
func SetLogoURLs(urls LogoURLS) OptPayP {
	return func(c *PaymentProviderOpts) {
		c.LogoUrls = urls
	}
}

// ISO.4217 currency codes supported by the Payment Provider. See Currency Codes.
func SetSupportedCurrencies(supportedCurrencies []string) OptPayP {
	return func(c *PaymentProviderOpts) {
		c.SupportedCurrencies = supportedCurrencies
	}
}

// List of available payment methods for each payment method type. See Payment Methods.
func SetSupportedPaymentMethods(p []SupportedPaymentMethods) OptPayP {
	return func(c *PaymentProviderOpts) {
		c.SupportedPaymentMethods = p
	}
}

// HTTPS URL of the JavaScript file to be included in the checkout frontend. See Checkout.
func SetCheckoutJsURL(checkoutJsURL string) OptPayP {
	return func(c *PaymentProviderOpts) {
		c.CheckoutJsURL = checkoutJsURL
	}
}

// Object containing the available payment options for the checkout frontend. See Checkout Payment Options.
func SetCheckoutPaymentOptions(p []CheckoutPaymentOption) OptPayP {
	return func(c *PaymentProviderOpts) {
		c.CheckoutPaymentOptions = p
	}
}

// [Optional] HTTPS URL of the Payment Provider configuration UI.
func SetConfigurationURL(configurationURL string) OptPayP {
	return func(c *PaymentProviderOpts) {
		c.ConfigurationURL = configurationURL
	}
}

// [Optional] Payment Provider support site HTTPS URL.
func SetSupportURL(supportURL string) OptPayP {
	return func(c *PaymentProviderOpts) {
		c.SupportURL = supportURL
	}
}

// [Optional] List of rates definitions for merchants by payment method type. See Rates.
func SetRates(rates []Rate) OptPayP {
	return func(c *PaymentProviderOpts) {
		c.Rates = rates
	}
}

// [Optional] HTTPS URL of the Payment Provider's rate information site.
func SetRatesURL(ratesURL string) OptPayP {
	return func(c *PaymentProviderOpts) {
		c.RatesURL = ratesURL
	}
}

// [Optional] List of features offered by the Payment Provider. See Features.
func SetFeatures(features []string) OptPayP {
	return func(c *PaymentProviderOpts) {
		c.Features = features
	}
}

// [Optional] Indicates whether the Payment Provider is activated or deactivated in the store. Defaults to true.
func SetEnabled(enabled bool) OptPayP {
	return func(c *PaymentProviderOpts) {
		c.Enabled = enabled
	}
}

// [Optional] Object containing the authentication method type and the store credentials to use on payment processing requests. See Authentication.
func SetAuthentication(a Authentication) OptPayP {
	return func(c *PaymentProviderOpts) {
		c.Authentication = a
	}
}

// Initialize a payment provider instance
func New(opts ...OptPayP) *PaymentProvider {
	dc := defaultPaymentProvider()
	for _, fn := range opts {
		fn(&dc)
	}
	return &PaymentProvider{
		PaymentProviderOpts: dc,
	}
}

// Create a Payment Provider for a given store.
func (p PaymentProvider) Create() (pr PaymentProvider, err error) {
	b, err := json.Marshal(p)
	if err != nil {
		return
	}
	br, err := create(b)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &pr)
	return
}

// Update a Payment Provider for a given store. This is especially useful to update the installments specs.
func (p PaymentProvider) Update(paymentProviderID string) (pr PaymentProvider, err error) {
	b, err := json.Marshal(p)
	if err != nil {
		return
	}
	br, err := update(b, paymentProviderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &pr)
	return
}

// Get all Payment Providers for a given store.
func GetAll() (pl []PaymentProvider, err error) {
	br, err := getAll()
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &pl)
	return
}

// Get a specific Payment Provider for a given store.
func GetByID(paymentProviderID string) (pl []PaymentProvider, err error) {
	br, err := getByID(paymentProviderID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &pl)
	return
}

// Delete a Payment Provider for a given store.
func Delete(paymentProviderID string) (err error) {
	err = delete(paymentProviderID)
	return
}
