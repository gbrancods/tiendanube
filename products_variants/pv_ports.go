// Product variants allow you to group a shoe with different sizes
// and colors in the same product. You can only create products with
// up to 1000 variants.
package productvariant

import (
	"encoding/json"
)

func defaultVariant() VariantOpts {
	return VariantOpts{}
}

// The unique numeric identifier for the Product Variant
func SetID(id int) OptVar {
	return func(v *VariantOpts) {
		v.ID = id
	}
}

// The id of the product's image associated with the variant
func SetImageID(imageID int) OptVar {
	return func(v *VariantOpts) {
		v.ImageID = imageID
	}
}

// Price of the Product Variant. null indicates the product will initiate a contact instead of a checkout process
func SetPrice(price string) OptVar {
	return func(v *VariantOpts) {
		v.Price = price
	}
}

// Lower price to display as a sale. The value of price will be displayed crossed out for comparison
func SetPromotionalPrice(promotionalPrice string) OptVar {
	return func(v *VariantOpts) {
		v.PromotionalPrice = promotionalPrice
	}
}

// Specifies whether or not Tiendanube/Nuvemshop tracks the number of items in stock for this product variant.
// Valid values are true if Tiendanube/Nuvemshop tracks the stock, false if it doesn't.
func SetStockManagement(stockManagement bool) OptVar {
	return func(v *VariantOpts) {
		v.StockManagement = stockManagement
	}
}

// Depth of the Product Variant in centimetres
func SetDepth(depth string) OptVar {
	return func(v *VariantOpts) {
		v.Depth = depth
	}
}

// Height of the Product Variant in centimetres
func SetHeight(height string) OptVar {
	return func(v *VariantOpts) {
		v.Height = height
	}
}

// List of the values of the attributes whose values define the variant. E.g.: Large, Medium, etc. It is
// important that the number of values is equal to the number of attributes within the products.
func SetValue(en, es, pt string) OptVar {
	return func(v *VariantOpts) {
		v.Values = Values{
			En: en,
			Es: es,
			Pt: pt,
		}
	}
}

// The id of the product associated with the variant
func SetProductID(productID int) OptVar {
	return func(v *VariantOpts) {
		v.ProductID = productID
	}
}

// Stock of the Product Variant. If stock_management is false the stock will be null
func SetStock(stock int) OptVar {
	return func(v *VariantOpts) {
		v.Stock = stock
	}
}

// Unique identifier of the Product Variant in your store
func SetSKU(sku string) OptVar {
	return func(v *VariantOpts) {
		v.Sku = sku
	}
}

// The Manufacturer Part Number (MPN) of the product
func SetMPN(mpn string) OptVar {
	return func(v *VariantOpts) {
		v.Mpn = mpn
	}
}

// Attribute to set the demographic that the product is designed for. It is optional and only supports
// this values: "newborn", "infant", "toddler", "kids" and "adult".
func SetAgeGroup(ageGroup string) OptVar {
	return func(v *VariantOpts) {
		v.AgeGroup = ageGroup
	}
}

// Attribute to specify the gender your product is designed for. It is optional and only supports
// the values: "female", "male" and "unisex"
func SetGender(gender string) OptVar {
	return func(v *VariantOpts) {
		v.Gender = gender
	}
}

// Weight of the Product Variant in kilograms
func SetWeight(weigth string) OptVar {
	return func(v *VariantOpts) {
		v.Weight = weigth
	}
}

// Width of the Product Variant in centimetres
func SetWidth(width string) OptVar {
	return func(v *VariantOpts) {
		v.Width = width
	}
}

// Cost of getting or producing the product. It is optional and, if present, has to be a number bigger than 0.
func SetCost(cost string) OptVar {
	return func(v *VariantOpts) {
		v.Cost = cost
	}
}

// Initialize a product variant instance
func New(opts ...OptVar) *Variant {

	dp := defaultVariant()

	for _, fn := range opts {
		fn(&dp)
	}

	return &Variant{
		VariantOpts: dp,
	}
}

// Create a new Product Variant
func (v Variant) Create(productID int) (r VariantOpts, err error) {
	b, err := json.Marshal(v)
	if err != nil {
		return
	}
	br, err := create(b, productID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Modify an existing Product Variant
func (v Variant) Update(productID int) (r VariantOpts, err error) {
	b, err := json.Marshal(v)
	if err != nil {
		return
	}
	br, err := update(b, productID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Updates the entire ProductVariant collection owned by a specific Product. Use this endpoint to add, modify or remove
// ProductVariants in one single batch operation.
//
// If the operation is successful, all the variants sent in the request will be the current and only variants for the Product.
//
// Each ProductVariant will be identified by its value combination. If a specified value combination doesn't exist,
// a new ProductVariant will be created, otherwise, the ProductVariant matching that value combination will be updated. Value
// combinations that aren't present in the request body, identify ProductVariants that will be deleted.
func (v Variant) UpdateById(productID, variantID int) (r VariantOpts, err error) {
	b, err := json.Marshal(v)
	if err != nil {
		return
	}
	br, err := updateById(b, productID, variantID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Partially update a ProductVariant collection. This endpoint allows to modify many existing ProductVariants that belong to the given Product.
//
// This endpoint will not add new ProductVariants or remove existing ProductVariants; it will just update their values.
//
// Preconditions
// Request body is an array of ProductVariant resources.
//
// Each ProductVariant in the request body must:
//
// - include the ProductVariant ID, set in the field id
//
// - correspond to an existing ProductVariant with same ID than the one set in the field id
//
// - belong to the Product which ID is set in {product_id} in the URL
//
// - have a unique combination of values among all ProductVariants of the Product (either ProductVariants included in the request or previously persisted).
func (v Variant) Patch(productID int) (r VariantOpts, err error) {
	b, err := json.Marshal(v)
	if err != nil {
		return
	}
	br, err := patch(b, productID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Receive a list of all Product Variants for a given product.
func GetAll(productID int) (r []VariantOpts, err error) {
	b, err := getAll(productID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Receive a single Product Variant
func GetByID(productID, variantID int) (r VariantOpts, err error) {
	b, err := getByID(productID, variantID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Remove a Product Variant
func Delete(productID, variantID int) (err error) {
	err = delete(productID, variantID)
	return
}

// Update the stock for one or all variants of a product by adding, removing or replacing.
func (sv StockVariant) Update(p int) (r VariantOpts, err error) {
	b, err := json.Marshal(sv)
	if err != nil {
		return
	}
	br, err := stockVariant(p, b)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}
