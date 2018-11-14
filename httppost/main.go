package main

import (
	"net/http"
	//"strings"
	"fmt"
	"io/ioutil"
	"bytes"
	"encoding/json"
)

type Masterendpoints struct {
	ClusterInstanceId string `json:"ClusterInstanceId"`
	ClusterName       string `json:"ClusterName"`
	Status            string `json:"Status"`
	Description       string `json:"Description"`
	ClusterCIDR       string `json:"ClusterCIDR"`
	ServiceCIDR       string `json:"ServiceCIDR"`
	MasterEndpoint    string `json:"MasterEndpoint"`
	CreateTime        string `json:"CreateTime"`
	UpdateTime        string `json:"UpdateTime"`
	K8sVersion        string `json:"K8sVersion"`
	CaCrt             string `json:"CaCrt"`
	ClusterId         string `json:"ClusterId"`
	AdminPasswd       string `json:"AdminPasswd"`
}
type TCEResponse struct {
	EndpointsDetail []Masterendpoints `json:"masterEndpoints"`
	TotalNum        int               `json:"totalNum"`
	RequestId       string            `json:"RequestId"`
}

type TCEServerResp struct {
	Response TCEResponse `json:"Response"`
}

//
func main() {
	clusterid := "cls-1w7zlz1x"
	json_string := fmt.Sprintf(`{"Version":"","RequestId":"","Action":"qcloud.container.getMasterEndpoints","Uin":"",
                      "AppId":1255000059,"SubAccountUin":"12345","ClusterId":"%s","Region":"shanghai","Offest":0,"Limit":10}`, clusterid)
	jsonStr := []byte(json_string)
	url := "http://shanghai.operation.ccs.tencentyun.com/cis/api"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	respObj := TCEServerResp{}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &respObj)
	fmt.Println(respObj)

}
