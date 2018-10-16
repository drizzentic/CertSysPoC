package utils
import (
	"net/http"
	"strings"
)

type Requests struct {
	Method string
}

func RpcCalls(r *Requests) *http.Response {

	body := strings.NewReader(`{ "method": "` + r.Method + `"}`)
	req, err := http.NewRequest("POST", "http://127.0.0.1:5000", body)
	if err != nil {
		// handle err
	}
	req.SetBasicAuth("rpcusername", "rpcpassword")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}

	return resp
}
