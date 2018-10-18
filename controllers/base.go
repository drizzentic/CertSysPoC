package controllers

import (
	"certSys/utils"
	"encoding/json"
	"io/ioutil"
)

func processRpcCalls(m string, params []string) Response {
	var address Response

	method := utils.Requests{m, params}
	response := utils.RpcCalls(&method, params)
	body, _ := ioutil.ReadAll(response.Body)

	if err := json.Unmarshal(body, &address); err != nil {

		panic(err)
	}

	defer response.Body.Close()

	return address
}