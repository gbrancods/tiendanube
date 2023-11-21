package webhooks

type OptWebH func(*WebhookOpts)

type Webhook struct {
	WebhookOpts
}

type WebhookOpts struct {
	CreatedAt string `json:"created_at,omitempty"`
	Event     string `json:"event,omitempty"`
	ID        int    `json:"id,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	URL       string `json:"url,omitempty"`
}
