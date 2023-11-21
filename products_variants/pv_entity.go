package productvariant

type OptVar func(*VariantOpts)

type Variant struct {
	VariantOpts
}

type VariantOpts struct {
	ID               int         `json:"id,omitempty"`
	ImageID          int         `json:"image_id,omitempty"`
	PromotionalPrice string      `json:"promotional_price,omitempty"`
	CreatedAt        string      `json:"created_at,omitempty"`
	Depth            string      `json:"depth,omitempty"`
	Height           string      `json:"height,omitempty"`
	Values           interface{} `json:"values,omitempty"` // Values
	Price            string      `json:"price,omitempty"`
	ProductID        int         `json:"product_id,omitempty"`
	StockManagement  bool        `json:"stock_management,omitempty"`
	Stock            int         `json:"stock,omitempty"`
	Sku              string      `json:"sku,omitempty"`
	Mpn              string      `json:"mpn,omitempty"`
	AgeGroup         string      `json:"age_group,omitempty"`
	Gender           string      `json:"gender,omitempty"`
	UpdatedAt        string      `json:"updated_at,omitempty"`
	Weight           string      `json:"weight,omitempty"`
	Width            string      `json:"width,omitempty"`
	Cost             string      `json:"cost,omitempty"`
}

type Values struct {
	En string `json:"en,omitempty"`
	Es string `json:"es,omitempty"`
	Pt string `json:"pt,omitempty"`
}

type StockVariant struct {
	Action string `json:"action,omitempty"`
	Value  int    `json:"value,omitempty"`
}
