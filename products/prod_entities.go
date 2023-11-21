package products

import (
	"github.com/gbrancods/tiendanube/categories"
	"github.com/gbrancods/tiendanube/i18n"
	productvariant "github.com/gbrancods/tiendanube/products_variants"
)

type OptPrd func(*ProductOpts)

type Product struct {
	ProductOpts
}

type ProductOpts struct {
	ID             int                       `json:"id,omitempty"`
	Name           i18n.Langs                `json:"name,omitempty"`
	Variants       []productvariant.Variant  `json:"variants,omitempty"`
	Categories     []categories.CategoryOpts `json:"categories,omitempty"`
	Attributes     []i18n.Langs              `json:"attributes,omitempty"`
	Description    i18n.Langs                `json:"description,omitempty"`
	Handle         i18n.Langs                `json:"handle,omitempty"`
	Images         []Image                   `json:"images,omitempty"`
	SeoTitle       i18n.Langs                `json:"seo_title,omitempty"`
	SeoDescription i18n.Langs                `json:"seo_description,omitempty"`
	CreatedAt      string                    `json:"created_at,omitempty"`
	VideoURL       string                    `json:"video_url,omitempty"`
	Brand          string                    `json:"brand,omitempty"`
	Published      bool                      `json:"published,omitempty"`
	FreeShipping   bool                      `json:"free_shipping,omitempty"`
	UpdatedAt      string                    `json:"updated_at,omitempty"`
}

type Image struct {
	Src string `json:"src"`
}

type StockPrice []struct {
	ID       int         `json:"id"`
	Variants interface{} `json:"variants"` // []Variant
}

type StockPriceResponse []struct {
	ID       int         `json:"id"`
	Variants interface{} `json:"variants"` // []Variant
}
