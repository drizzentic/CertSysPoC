package utils

import (
	"fmt"
	"net/http"
	"strings"
)

type Requests struct {
	Method string
	Params []string
}

var configs = GetConfigs()
var (
	username = configs.Rpcusername
	password = configs.Rpcpassword
)

func RpcCalls(r *Requests, p []string) *http.Response {
	params := strings.Join(p, "\",\"")
	body := strings.NewReader(`{ "method": "` + r.Method + `","params":["` + params + `"]}`)
	fmt.Print(&body)
	req, err := http.NewRequest("POST", "http://127.0.0.1:5000", body)
	if err != nil {
		// handle err
	}

	req.SetBasicAuth(username, password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}

	return resp

}
