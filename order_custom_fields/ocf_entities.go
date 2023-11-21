package ordercustomfields

type OptOCF func(*CustomFieldsOpts)

type CustomFields struct {
	CustomFieldsOpts
}

type CustomFieldsOpts struct {
	ID            string   `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	OwnerResource string   `json:"owner_resource,omitempty"`
	ValueType     string   `json:"value_type,omitempty"`
	Source        string   `json:"source,omitempty"`
	Description   string   `json:"description,omitempty"`
	ReadOnly      bool     `json:"read_only,omitempty"`
	CreatedAt     string   `json:"created_at,omitempty"`
	UpdatedAt     string   `json:"updated_at,omitempty"`
	Values        []string `json:"values,omitempty"`
}

type Values []struct {
	CustomfieldID string `json:"customfield_id"`
	Value         string `json:"value"`
}
