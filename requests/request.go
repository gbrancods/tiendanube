package requests

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gbrancods/tiendanube/auth"
)

type Request struct {
	Destiny string
	Method  string
	Body    io.Reader
}

func (r Request) Do() (tr []byte, err error) {
	url := fmt.Sprintf("https://%s/v1/%d/%s/", auth.GetAPIurl(), auth.GetUserID(), r.Destiny)

	client := &http.Client{}
	req, _ := http.NewRequest(r.Method, url, r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authentication", auth.GetToken())

	res, err := client.Do(req)
	if err != nil {
		return
	}

	tr, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}

	if res.StatusCode > 299 {
		err = fmt.Errorf("status code %d\nresponse:\n%s", res.StatusCode, string(tr))
		return
	}

	return
}
