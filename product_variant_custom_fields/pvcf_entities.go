package pvcustomfields

type OptVCF func(*ProductVariantCustomFieldsOpts)

type ProductVariantCustomFields struct {
	ProductVariantCustomFieldsOpts
}

type ProductVariantCustomFieldsOpts struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	ValueType     string      `json:"value_type"`
	ReadOnly      bool        `json:"read_only"`
	OwnerResource string      `json:"owner_resource"`
	Values        interface{} `json:"values"`
}

type Value struct {
	Value   string `json:"value"`
	Created bool   `json:"created"`
}
