package metafields

type OptMet func(*MetafieldOpts)

type Metafield struct {
	MetafieldOpts
}

type MetafieldOpts struct {
	ID            int    `json:"id,omitempty"`
	Namespace     string `json:"namespace,omitempty"`
	Description   string `json:"description,omitempty"`
	Key           string `json:"key,omitempty"`
	Value         string `json:"value,omitempty"`
	OwnerID       string `json:"owner_id,omitempty"`
	OwnerResource string `json:"owner_resource,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	DeletedAt     string `json:"deleted_at,omitempty"`
}

type CouponResponse struct {
	ID         int    `json:"id"`
	Code       string `json:"code"`
	Type       string `json:"type"`
	Value      string `json:"value"`
	Valid      bool   `json:"valid"`
	Used       int    `json:"used"`
	MaxUses    int    `json:"max_uses"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	MinPrice   int    `json:"min_price"`
	Categories []struct {
		ID   int `json:"id"`
		Name struct {
			Es string `json:"es"`
			En string `json:"en"`
		} `json:"name"`
		Description struct {
			Es string `json:"es"`
			En string `json:"en"`
		} `json:"description"`
		Handle struct {
			Es string `json:"es"`
			En string `json:"en"`
		} `json:"handle"`
		Parent        interface{}   `json:"parent"`
		Subcategories []interface{} `json:"subcategories"`
		CreatedAt     string        `json:"created_at"`
		UpdatedAt     string        `json:"updated_at"`
	} `json:"categories"`
}
