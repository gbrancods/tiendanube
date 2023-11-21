// The Store resource contains general settings and information about a Tiendanube/Nuvemshop's store.
package store

import (
	"encoding/json"
)

// Receive a single Store.
func GetInfo() (r Store, err error) {
	b, err := getInfo()
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}
