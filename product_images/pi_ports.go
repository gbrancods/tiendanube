// Product images could well be the single most important design aspect of
// any store. Without the ability to touch, hold, smell, taste or otherwise
// handle the products they are interested in, potential customers have only images to interact with.
//
// The product images have the following restrictions:
//
//	Must weight less than 10MB
//	Must be in one of the following formats: .gif, .jpg, .png
package productimages

import (
	"encoding/json"
)

func defaultProductImage() ProductImageOpts {
	return ProductImageOpts{}
}

// URL of the product image
func SetSrc(imageURL string) OptPImg {
	return func(c *ProductImageOpts) {
		c.Src = imageURL
	}
}

// Must be image encoded in a Base64 String.
func SetAttachment(base64 string) OptPImg {
	return func(c *ProductImageOpts) {
		c.Attachment = base64
	}
}

// A filename string (ex: mewtwo.gif)
func SetFilename(filename string) OptPImg {
	return func(c *ProductImageOpts) {
		c.Filename = filename
	}
}

// A number indicating in which position of the image list, the new image should be placed
func SetPosition(position int) OptPImg {
	return func(c *ProductImageOpts) {
		c.Position = position
	}
}

// Initialize a Product Image instance
func New(opts ...OptPImg) *ProductImage {
	pi := defaultProductImage()
	for _, fn := range opts {
		fn(&pi)
	}
	return &ProductImage{
		ProductImageOpts: pi,
	}
}

// Receive a list of all Product Images for a given product.
func GetByProduct(productID int) (r ProductImage, err error) {
	b, err := getByProduct(productID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Receive a single Product Image
func GetByID(productID, imageID int) (r ProductImage, err error) {
	b, err := getByID(productID, imageID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Create a new Product Image
func (p ProductImage) Create(productID int) (r ProductImage, err error) {
	b, err := json.Marshal(p)
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

// Modify an existing Product Image
func (p ProductImage) Update(productID, imageID int) (r ProductImage, err error) {
	b, err := json.Marshal(p)
	if err != nil {
		return
	}
	br, err := update(b, productID, imageID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Remove a Product Image
func Delete(productID, imageID int) (err error) {
	err = delete(productID, imageID)
	return
}
