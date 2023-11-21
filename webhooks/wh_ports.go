/*
A Webhook is a tool that allows you to receive a notification for a
certain event. It allows you to register an https URL which will receive
the event data, stored in JSON.

	Category / Events:

	App: uninstalled/suspended/resumed
	Category: created/updated/deleted
	Order: created/updated/paid/packed/fulfilled/cancelled/custom_fields_updated
	Product: created/updated/deleted
	Product Variant: custom_fields_updated
	Domain: updated
	Order Custom Field: created/updated/deleted
	Product Variant Custom Field: created/updated/deleted
*/
package webhooks

import "encoding/json"

func defaultWebhook() WebhookOpts {
	return WebhookOpts{}
}

// The URL where the webhook should send the POST request when the event occurs. Must be HTTPS.
func SetURL(url string) OptWebH {
	return func(w *WebhookOpts) {
		w.URL = url
	}
}

// The event that will trigger the webhook
func SetEvent(event string) OptWebH {
	return func(w *WebhookOpts) {
		w.Event = event
	}
}

// Initialize a webhook instance
func New(opts ...OptWebH) *Webhook {
	dt := defaultWebhook()
	for _, fn := range opts {
		fn(&dt)
	}
	return &Webhook{
		WebhookOpts: dt,
	}
}

// Receive a list of all Webhooks for your application.
func GetAll() (r []Webhook, err error) {
	b, err := getAll()
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Receive a single Webhook
func GetById(webhookID int) (r Webhook, err error) {
	b, err := getByID(webhookID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Create a new Webhook
func (w Webhook) Create() (r Webhook, err error) {
	j, err := json.Marshal(w)
	if err != nil {
		return
	}
	br, err := create(j)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Modify an existing Webhook
func (w Webhook) Update(webhookID int) (r Webhook, err error) {
	j, err := json.Marshal(w)
	if err != nil {
		return
	}
	br, err := update(j, webhookID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Remove a Webhook
func Delete(webhookID int) (err error) {
	err = delete(webhookID)
	return
}
