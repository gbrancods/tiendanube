// A Category lets the store owner group his/her products to make the store easier to browse.
package categories

import (
	"encoding/json"

	"github.com/gbrancods/tiendanube/i18n"
)

func defaultCategory() CategoryOpts {
	return CategoryOpts{}
}

// List of the names of the Category, in every language supported by the store
func SetName(cn i18n.Langs) OptCat {
	return func(c *CategoryOpts) {
		c.Name = cn
	}
}

// Show Categories with a given parent category
func SetParent(parentID int) OptCat {
	return func(c *CategoryOpts) {
		c.Parent = parentID
	}
}

// List of the descriptions of the Category, as HTML, in every language supported by the store
func SetDescription(description i18n.Langs) OptCat {
	return func(c *CategoryOpts) {
		c.Description = description
	}
}

// Show Categories with a given URL
func SetHandle(handle i18n.Langs) OptCat {
	return func(c *CategoryOpts) {
		c.Handle = handle
	}
}

// Attributes used to categorize an item. This category is selected from the Google’s taxonomy
// See more:
// https://www.google.com/basepages/producttype/taxonomy.en-US.txt
// https://www.google.com/basepages/producttype/taxonomy.es-ES.txt
// https://www.google.com/basepages/producttype/taxonomy.pt-BR.txt
func SetGoogleShoppingCategory(category string) OptCat {
	return func(c *CategoryOpts) {
		c.GoogleShoppingCategory = category
	}
}

// Initialize a categorie instance
func New(opts ...OptCat) *Category {
	dc := defaultCategory()
	for _, fn := range opts {
		fn(&dc)
	}
	return &Category{
		CategoryOpts: dc,
	}
}

// Create a new Category
func (c Category) Create() (r Category, err error) {
	j, err := json.Marshal(c)
	if err != nil {
		return
	}
	br, err := create(j)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Modify an existing Category
func (c Category) Update(categoryID int) (r Category, err error) {
	j, err := json.Marshal(c)
	if err != nil {
		return
	}
	br, err := update(j, categoryID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Receive a list of all Categories
func GetAll() (r []Category, err error) {
	b, err := getAll()
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Receive a single Category
func GetById(id int) (r Category, err error) {
	b, err := getByID(id)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Remove a Category
func Delete(id int) (err error) {
	err = delete(id)
	return
}
