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

func RpcCalls(r *Requests, p []string, data string, vout int) *http.Response {
	params := strings.Join(p, "\",\"")
	var formattedIntParams string
	var body *strings.Reader
	//TODO: Handle int and string params intelligently
	if r.Method == "generate" {
		formattedIntParams = strings.Join(p, ",")
		body = strings.NewReader(`{ "method": "` + r.Method + `","params":[` + formattedIntParams + `]}`)
	} else if r.Method == "createrawtransaction" {
		raw := fmt.Sprintf(`{ "method": "%s","params":["[{\"txid\":\"%s\",\"vout\":%d}]","{\"data\":\"%s\"}"]}`, r.Method, params, vout, data)
		body = strings.NewReader(raw)
	} else {
		body = strings.NewReader(`{ "method": "` + r.Method + `","params":["` + params + `"]}`)
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:19001", body)
	if err != nil {
		// handle err
		//fmt.Println(err)
	}

	req.SetBasicAuth(username, password)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		// handle err
		//fmt.Println(err)
	}

	return resp

}

func HttpCalls(r bool, file string) {

	resp, err := http.Get("http://127.0.0.1:5000/api/v0/add?arg=" + file + "&hash=sha2-256")
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	fmt.Print(resp.StatusCode)

}
