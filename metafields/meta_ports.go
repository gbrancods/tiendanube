/*
The metafields are Namespaced Key - Value store for Apps.
The metafields can only be associated with the following entities:

	Product
	Product_Variant
	Category
	Page
	Order
	Customer

To do that you need to set the owner_resource to one of the above, an example would be owner_resource='Product'.
An use example would be use it in an app for Bookstores that require associate a book (product) with it's author and genre to provide an advanced search of books by this fields.
*/
package metafields

import (
	"encoding/json"
)

func defaultMetafields() MetafieldOpts {
	return MetafieldOpts{}
}

// String that identifies the metafield in some namespace. It can be any string that starts with a letter followed only by: a-z A-Z 0-9 or _.
func SetKey(key string) OptMet {
	return func(m *MetafieldOpts) {
		m.Key = key
	}
}

// Metafield's value (string).
func SetValue(value string) OptMet {
	return func(m *MetafieldOpts) {
		m.Value = value
	}
}

// The namespace where the metafield makes sense. It can be any string that starts with a letter followed only by: a-z A-Z 0-9 or _.
func SetNamespace(namespace string) OptMet {
	return func(m *MetafieldOpts) {
		m.Namespace = namespace
	}
}

// String explaining the metafield's meaning (optional).
func SetDescription(description string) OptMet {
	return func(m *MetafieldOpts) {
		m.Description = description
	}
}

// Entity id to which is associated the metaField.
func SetOwnerID(ownerID string) OptMet {
	return func(m *MetafieldOpts) {
		m.OwnerID = ownerID
	}
}

// The value is mandatory and can be only the allowed values mentioned at the package documentation.
func SetOwnerResource(ownerResource string) OptMet {
	return func(m *MetafieldOpts) {
		m.OwnerResource = ownerResource
	}
}

// Initialize a metafield instance
func New(opts ...OptMet) *Metafield {
	dm := defaultMetafields()
	for _, fn := range opts {
		fn(&dm)
	}
	return &Metafield{
		MetafieldOpts: dm,
	}
}

// Create a new metafield
func (m Metafield) Create() (r Metafield, err error) {
	b, err := json.Marshal(m)
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

// Modify an existing metafield. You can update the metafield value or/and the description associated.
func (m Metafield) Update(metafieldID int) (r Metafield, err error) {
	b, err := json.Marshal(m)
	if err != nil {
		return
	}
	br, err := update(b, metafieldID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Examples are: products, product_variants, categories and pages
func GetByOwnerResource(ownerResource string) (r []Metafield, err error) {
	b, err := getByOwnerResource(ownerResource)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Receive a single metafield
func GetByID(metafieldID int) (r Metafield, err error) {
	b, err := getByID(metafieldID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Delete an existing metafield
func Delete(metafieldID int) (err error) {
	err = delete(metafieldID)
	return
}
