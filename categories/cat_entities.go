package categories

import "github.com/gbrancods/tiendanube/i18n"

type OptCat func(*CategoryOpts)

type Category struct {
	CategoryOpts
}

type CategoryOpts struct {
	ID                     int        `json:"id,omitempty"`
	Name                   i18n.Langs `json:"name,omitempty"`
	Description            i18n.Langs `json:"description,omitempty"`
	Handle                 i18n.Langs `json:"handle,omitempty"`
	Parent                 int        `json:"parent,omitempty"`
	GoogleShoppingCategory string     `json:"google_shopping_category,omitempty"`
	Subcategories          []int      `json:"subcategories,omitempty"`
	CreatedAt              string     `json:"created_at,omitempty"`
	UpdatedAt              string     `json:"updated_at,omitempty"`
}
