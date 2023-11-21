// The Script resource allows the app to register custom Javascript to be run in the
// store or in the checkout page. Also, if the user deletes your app, then he will not
// have to edit his theme to remove your JavaScript. When an app is uninstalled from a
// store, all of the scripts the app created are automatically removed along with it.
package script

import (
	"encoding/json"
)

func defaultScript() ScriptOpts {
	return ScriptOpts{}
}

// Specifies the location of the Script. Must be HTTPS.
func SetSrc(src string) OptScr {
	return func(v *ScriptOpts) {
		v.Src = src
	}
}

// DOM event which triggers the loading of the script. Valid values are onfirstinteraction (default) or onload
func SetEvent(event string) OptScr {
	return func(v *ScriptOpts) {
		v.Event = event
	}
}

// A comma-separated list of places where the javascript will run. Valid values are store (default) or checkout
func SetWhere(where string) OptScr {
	return func(v *ScriptOpts) {
		v.Where = where
	}
}

// Initialize a script  instance
func New(opts ...OptScr) *Script {
	dp := defaultScript()
	for _, fn := range opts {
		fn(&dp)
	}
	return &Script{
		ScriptOpts: dp,
	}
}

// Receive a list of all Scripts.
func GetAll() (r []Script, err error) {
	b, err := getAll()
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Receive a single Script
func GetByID(scriptID int) (r Script, err error) {
	b, err := getByID(scriptID)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}

// Create a new Script.
func (s Script) Create() (r Script, err error) {
	b, err := json.Marshal(s)
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

// Modify an existing Script
func (s Script) Update(scriptID int) (r Script, err error) {
	b, err := json.Marshal(s)
	if err != nil {
		return
	}
	br, err := update(b, scriptID)
	if err != nil {
		return
	}
	err = json.Unmarshal(br, &r)
	return
}

// Remove a Script
func Delete(scriptID int) (err error) {
	err = delete(scriptID)
	return
}
