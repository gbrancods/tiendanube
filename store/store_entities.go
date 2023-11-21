package store

import "github.com/gbrancods/tiendanube/i18n"

type Store struct {
	Address             string     `json:"address,omitempty"`
	AdminLanguage       string     `json:"admin_language,omitempty"`
	Blog                string     `json:"blog,omitempty"`
	BusinessID          string     `json:"business_id,omitempty"`
	BusinessName        string     `json:"business_name,omitempty"`
	BusinessAddress     string     `json:"business_address,omitempty"`
	ContactEmail        string     `json:"contact_email,omitempty"`
	Country             string     `json:"country,omitempty"`
	CreatedAt           string     `json:"created_at,omitempty"`
	CustomerAccounts    string     `json:"customer_accounts,omitempty"`
	Description         i18n.Langs `json:"description,omitempty"`
	Domains             []string   `json:"domains,omitempty"`
	Email               string     `json:"email,omitempty"`
	Facebook            string     `json:"facebook,omitempty"`
	GooglePlus          string     `json:"google_plus,omitempty"`
	ID                  int        `json:"id,omitempty"`
	Instagram           string     `json:"instagram,omitempty"`
	Languages           Languages  `json:"languages,omitempty"`
	Logo                string     `json:"logo,omitempty"`
	MainCurrency        string     `json:"main_currency,omitempty"`
	CurrentTheme        string     `json:"current_theme,omitempty"`
	MainLanguage        string     `json:"main_language,omitempty"`
	Name                i18n.Langs `json:"name,omitempty"`
	OriginalDomain      string     `json:"original_domain,omitempty"`
	Phone               string     `json:"phone,omitempty"`
	WhatsappPhoneNumber string     `json:"whatsapp_phone_number,omitempty"`
	Pinterest           string     `json:"pinterest,omitempty"`
	PlanName            string     `json:"plan_name,omitempty"`
	Type                string     `json:"type,omitempty"`
	Twitter             string     `json:"twitter,omitempty"`
	HasMulticd          bool       `json:"has_multicd,omitempty"`
}

type Languages struct {
	En Currency `json:"en,omitempty"`
	Es Currency `json:"es,omitempty"`
	Pt Currency `json:"pt,omitempty"`
}

type Currency struct {
	Currency string `json:"currency,omitempty"`
	Active   bool   `json:"active,omitempty"`
}
