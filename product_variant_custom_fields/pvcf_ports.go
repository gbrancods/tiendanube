// A custom field allows the store owner/merchant to expand their experience and control their own business
// through personalized and unique custom fields for product variants.
package pvcustomfields

import (
	"encoding/json"
)

func defaultVariantCustomFields() ProductVariantCustomFieldsOpts {
	return ProductVariantCustomFieldsOpts{}
}

// Name of the custom field
func SetName(name string) OptVCF {
	return func(v *ProductVariantCustomFieldsOpts) {
		v.Name = name
	}
}

// Description of the custom field
func SetDescription(description string) OptVCF {
	return func(v *ProductVariantCustomFieldsOpts) {
		v.Description = description
	}
}

// Custom field type (text_list, text, numeric, date)
func SetValueType(valueType string) OptVCF {
	return func(v *ProductVariantCustomFieldsOpts) {
		v.ValueType = valueType
	}
}

// If set to true, it restricts the association of the custom field by merchants via the administrator panel,
// the merchant can only read the value associated with the custom field. (default value is false)
func SetReadOnly(readOnly bool) OptVCF {
	return func(v *ProductVariantCustomFieldsOpts) {
		v.ReadOnly = readOnly
	}
}

// A list of all values for a custom field (for value_type text_list only, for other types can be an empty array)
func SetValues(values []string) OptVCF {
	return func(v *ProductVariantCustomFieldsOpts) {
		v.Values = values
	}
}

// Initialize a product variant custom field  instance
func New(opts ...OptVCF) *ProductVariantCustomFields {
	dp := defaultVariantCustomFields()
	for _, fn := range opts {
		fn(&dp)
	}
	return &ProductVariantCustomFields{
		ProductVariantCustomFieldsOpts: dp,
	}
}

// Create a new custom field
func (c ProductVariantCustomFields) Create(pid int) (r ProductVariantCustomFields, err error) {
	b, err := json.Marshal(c)
	if err != nil {
		return
	}
	br, err := create(b)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Return a list of all custom fields from a specific owner resource
func GetAll() (r []ProductVariantCustomFields, err error) {
	b, err := getAll()
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Update the custom field values
func (c ProductVariantCustomFields) Update(id string) (r ProductVariantCustomFields, err error) {
	b, err := json.Marshal(c)
	if err != nil {
		return
	}
	br, err := update(b, id)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// List custom fields associated with a specific product variant
func GetByVariant(id int) (r ProductVariantCustomFields, err error) {
	b, err := getByVariant(id)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// List product variants associated with a specific custom field
func GetOwners(id string) (r ProductVariantCustomFields, err error) {
	b, err := getOwners(id)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Update a value associated with a product variant.
func (c ProductVariantCustomFields) UpdateValue(id string) (r ProductVariantCustomFields, err error) {
	b, err := json.Marshal(c)
	if err != nil {
		return
	}
	br, err := updateValue(b, id)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Delete a custom field
func Delete(id string) (err error) {
	err = delete(id)
	return
}

// List the data of a given product variant custom field
func GetByID(id string) (r ProductVariantCustomFields, err error) {
	b, err := getByID(id)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}
