// A Location of the store where physical goods reside.
package locations

import (
	"encoding/json"
)

func defaultLocation() LocationOpts {
	return LocationOpts{}
}

// Localized Location Name (Required)
func SetName(name Name) OptLoc {
	return func(l *LocationOpts) {
		l.Name = name
	}
}

// The Address of the Location (Required)
func SetAddress(address Address) OptLoc {
	return func(l *LocationOpts) {
		l.Address = address
	}
}

// "true" if it is the default Location, "false" otherwise
func SetIsDefault(isDefault bool) OptLoc {
	return func(l *LocationOpts) {
		l.IsDefault = isDefault
	}
}

//TODO
// "true" if this Location allows pickup, "false" otherwise
//func SetAllowsPickup(allowsPickup bool) OptLoc {
//	return func(l *LocationOpts) {
//		l. = isDefault
//	}
//}

// Priority of the location to assign stock during checkout. The lower the value the higher the priority
func SetPriority(isDefault bool) OptLoc {
	return func(l *LocationOpts) {
		l.IsDefault = isDefault
	}
}

func New(opts ...OptLoc) *Location {
	dc := defaultLocation()
	for _, fn := range opts {
		fn(&dc)
	}
	return &Location{
		LocationOpts: dc,
	}
}

func (l Location) Create() (r Location, err error) {
	b, err := json.Marshal(l)
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

func (l Location) Update(locationID string) (r Location, err error) {
	b, err := json.Marshal(l)
	if err != nil {
		return
	}
	br, err := update(b, locationID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

func GetAll() (l []Location, err error) {
	b, err := getAll()
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &l)
	return
}

func GetByID(locationID string) (l Location, err error) {
	b, err := getByID(locationID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &l)
	return
}

func Delete(locationID string) (err error) {
	err = delete(locationID)
	return
}

func GetInventoryLevels(locationID string) (r InventoryLevel, err error) {
	b, err := getInventoryLevels(locationID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}
