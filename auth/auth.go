/*
Necessary partner account

https://partners.tiendanube.com/

or

https://partners.nuvemshop.com.br/

See more in the official documentation:

https://dev.tiendanube.com/en/docs/applications/authentication

or

https://dev.nuvemshop.com.br/en/docs/applications/authentication
*/
package auth

import "fmt"

//type OptAuth func(*AuthOpts)

var au *auth

type Access struct {
	//Example: "api.nuvemshop.com.br" or "api.nuvemshop.com"
	APIurl      string
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	UserID      int    `json:"user_id"`
}

type auth struct {
	Auth Access
}

func (a Access) Start() error {

	au = new(auth)

	au.Auth = Access{
		AccessToken: a.AccessToken,
		TokenType:   a.TokenType,
		UserID:      a.UserID,
		APIurl:      a.APIurl,
	}

	return nil
}

func GetAuth() Access {
	return au.Auth
}

func GetToken() (t string) {
	t = fmt.Sprintf("%s %s", au.Auth.TokenType, au.Auth.AccessToken)
	return
}

func GetUserID() (id int) {
	id = au.Auth.UserID
	return
}

func GetAPIurl() (url string) {
	url = au.Auth.APIurl
	return
}
