// A Product is an item for sale in a Tiendanube/Nuvemshop's store. It can be either a good or a service.
package products

import (
	"encoding/json"

	"github.com/gbrancods/tiendanube/i18n"
	variants "github.com/gbrancods/tiendanube/products_variants"
)

func defaultProduct() ProductOpts {
	return ProductOpts{}
}

// List of the names of the Product, in every language supported by the store
func SetName(pn i18n.Langs) OptPrd {
	return func(p *ProductOpts) {
		p.Name = pn
	}
}

// List of Product Image objects representing the Product's images
func SetImages(il []Image) OptPrd {
	return func(p *ProductOpts) {
		p.Images = il
	}
}

// String with a valid URL format. Only admits https links
func SetVideoURL(u string) OptPrd {
	return func(p *ProductOpts) {
		p.VideoURL = u
	}
}

// List of Product Variant objects representing the different version of the Product
func SetVariants(vl []variants.Variant) OptPrd {
	return func(p *ProductOpts) {
		p.Variants = vl
	}
}

// TODO - incompatible with return, two treatments for the same field
// List of Category Ids representing the Product's categories
//func SetCategories(c []int) OptPrd {
//	return func(p *ProductOpts) {
//		p.Categories = c
//	}
//}

// Initialize a product instance
func New(opts ...OptPrd) *Product {

	dp := defaultProduct()

	for _, fn := range opts {
		fn(&dp)
	}

	return &Product{
		ProductOpts: dp,
	}
}

// Create a new Product
func (p Product) Create() (r Product, err error) {
	b, err := json.Marshal(p)
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

// Modify an existing Product
func (p Product) Update() (r Product, err error) {
	b, err := json.Marshal(p)
	if err != nil {
		return
	}
	br, err := update(b)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Receive a list of all Products
func GetAll() (r []Product, err error) {
	br, err := getAll()
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)

	return
}

// Receive a single Product
func GetByID(id int) (r Product, err error) {
	br, err := getByID(id)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Returns the first Product found where one of its variants has the given SKU
func GetBySKU(sku string) (r Product, err error) {
	br, err := getBySKU(sku)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Remove a Product
func Delete(id int) (err error) {
	err = delete(id)
	return
}

// Updates Stock or Price of multiple products and variants. 50 different variants
// can be updated at once taking all products into account.
func (s StockPrice) Update() (r StockPriceResponse, err error) {
	b, err := json.Marshal(s)
	if err != nil {
		return
	}
	br, err := stockPrice(b)
	if err != nil {
		return
	}
	json.Unmarshal(br, &r)
	return
}
