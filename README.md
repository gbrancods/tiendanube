# Tiendanube Go SDK 
### A complete sdk for the Tiendanube *(beta version)*

### How to use:

```go
// All documentation is imported, hover over imports or functions for more details

//Starting
package tiendanube

import (
	"os"

	"github.com/gbrancods/tiendanube/auth"
)

func Start() {
	auth.Access{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		TokenType:   "bearer",
		UserID:      1234567,
		APIurl:      "api.nuvemshop.com.br",
	}.Start()
}
```
