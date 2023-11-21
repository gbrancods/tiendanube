package productimages

type OptPImg func(*ProductImageOpts)

type ProductImage struct {
	ProductImageOpts
}

type ProductImageOpts struct {
	ID         int    `json:"id,omitempty"`
	Src        string `json:"src,omitempty"`
	Filename   string `json:"filename,omitempty"`
	Attachment string `json:"attachment,omitempty"`
	Position   int    `json:"position,omitempty"`
	ProductID  int    `json:"product_id,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
}
