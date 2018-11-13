package main

import (
	"net/http"
	//"strings"
	"fmt"
	"io/ioutil"
	"bytes"
	"encoding/json"
)

type ServerResp struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//
func main(){
	jsonStr :=[]byte(`{"Name":"dl-test","TenantId":"b8a0074b63c34b8297cffb13d3a13dd2","EnvId":"4ce63d9c-8465-4756-9296-e8c7a6c2e727","Type":"static","ServerIP":"10.0.90.48","ServerPath":"/test","Reclaiming":"Delete","Capacity":"1G","AccessModes":"ReadWriteMany","Driver":"nfs"}`)
	url:= "http://localhost:8080/whale/v1/pvc"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	respObj := ServerResp{}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body,&respObj)
	fmt.Println(respObj.Code)

}
