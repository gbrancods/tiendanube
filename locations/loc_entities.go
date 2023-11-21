package locations

type OptLoc func(*LocationOpts)

type Location struct {
	LocationOpts
}

type LocationOpts struct {
	ID        string      `json:"id,omitempty"`
	Name      Name        `json:"name,omitempty"`
	StoreID   string      `json:"store_id,omitempty"`
	Priority  int         `json:"priority,omitempty"`
	Address   interface{} `json:"address,omitempty"` // Address
	IsDefault bool        `json:"is_default,omitempty"`
	CreatedAt string      `json:"created_at,omitempty"`
	UpdatedAt string      `json:"updated_at,omitempty"`
}

type Name struct {
	EsAR string `json:"es_AR,omitempty"`
	PtBR string `json:"pt_BR,omitempty"`
	EnUS string `json:"en_US,omitempty"`
}

type Address struct {
	Zipcode        string      `json:"zipcode,omitempty"`
	Street         string      `json:"street,omitempty"`
	Number         string      `json:"number,omitempty"`
	Floor          string      `json:"floor,omitempty"`
	Locality       string      `json:"locality,omitempty"`
	Reference      string      `json:"reference,omitempty"`
	BetweenStreets string      `json:"between_streets,omitempty"`
	City           string      `json:"city,omitempty"`
	Province       interface{} `json:"province,omitempty"`
	Region         interface{} `json:"region,omitempty"`
	Country        interface{} `json:"country,omitempty"`
}

type Province struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type Region struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type Country struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type InventoryLevel struct {
	ID         string `json:"id,omitempty"`
	VariantID  string `json:"variant_id,omitempty"`
	LocationID string `json:"location_id,omitempty"`
	Stock      int    `json:"stock,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
}
