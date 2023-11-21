// A custom field allows the store owner/merchant to expand their experience and control their own
// business through personalized and unique custom fields for orders.
package ordercustomfields

import (
	"encoding/json"
)

func defaultOrderCustomField() CustomFieldsOpts {
	return CustomFieldsOpts{}
}

// Name of the custom field
func SetName(name string) OptOCF {
	return func(ocf *CustomFieldsOpts) {
		ocf.Name = name
	}
}

// Description of the custom field
func SetDescription(description string) OptOCF {
	return func(ocf *CustomFieldsOpts) {
		ocf.Description = description
	}
}

// Custom field type (text_list, text, numeric, date)
func SetValueType(valueType string) OptOCF {
	return func(ocf *CustomFieldsOpts) {
		ocf.ValueType = valueType
	}
}

// Custom field owner (order)
func SetOwnerResource(ownerResource string) OptOCF {
	return func(ocf *CustomFieldsOpts) {
		ocf.OwnerResource = ownerResource
	}
}

// If set to true, it restricts the association of the custom field by merchants via the administrator panel,
// the merchant can only read the value associated with the custom field. (default value is false)
func SetReadOnly(readOnly bool) OptOCF {
	return func(ocf *CustomFieldsOpts) {
		ocf.ReadOnly = readOnly
	}
}

// A list of all values for a custom field (for value_type text_list only, for other types can be an empty array)
func SetValues(values []string) OptOCF {
	return func(ocf *CustomFieldsOpts) {
		ocf.Values = values
	}
}

// Initialize a order custom field instance
func New(opts ...OptOCF) *CustomFields {
	ocf := defaultOrderCustomField()
	for _, fn := range opts {
		fn(&ocf)
	}
	return &CustomFields{
		CustomFieldsOpts: ocf,
	}
}

// Create a new custom field
func (c CustomFields) Create() (r CustomFields, err error) {
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
func GetAll() (r []CustomFields, err error) {
	br, err := getAll()
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Update the custom field values
func (c CustomFields) Update(customFieldID string) (r CustomFields, err error) {
	b, err := json.Marshal(c)
	if err != nil {
		return
	}
	br, err := update(b, customFieldID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// List custom fields associated with a specific order
func GetByOrder(id string) (r CustomFields, err error) {
	br, err := getByOrder(id)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// List orders associated with a specific custom field
func GetByOwner(id string) (r CustomFields, err error) {
	br, err := getByOwner(id)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Update a value associated with a order on order details.
func (v Values) UpdateValue(orderID int) (r CustomFields, err error) {
	b, err := json.Marshal(v)
	if err != nil {
		return
	}
	err = updateValue(b, orderID)
	return
}

// Delete a custom field
func DeleteProduct(id int) (err error) {
	err = delete(id)
	return
}

// List the data of a given order custom field
func GetCustomFieldByID(id string) (r CustomFields, err error) {
	br, err := getByID(id)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}
