package main

import (
	"io/ioutil"
	"github.com/astaxie/beego/logs"
	"bytes"
	"text/template"
	"fmt"
	"os"
)

var DIR_TEMP string = "temp"
var FILE_PREFIX_YAML string = "yaml"

func newTempFile(content string) (string, error) {
	file, err := ioutil.TempFile(DIR_TEMP, FILE_PREFIX_YAML)

	if err != nil {
		logs.Error("create yaml file failed!")
		//panic(err)
		return "",err
	}

	defer file.Close()

	path := file.Name()
	file.WriteString(content)

	return path , nil
}

func parseConfFile(tpl string, obj interface{}) (string, error) {
	var buf bytes.Buffer

	t, err := template.ParseFiles(tpl)
	if err != nil {
		logs.Error(err.Error())
		return string(buf.Bytes()), err
	}
	if err := t.Execute(&buf, obj); err != nil {
		return string(buf.Bytes()), err
	}

	return string(buf.Bytes()), nil
}

//func kubectlCreateViaYaml(yaml string) (out string, err error) {
//	file := newTempFile()
//	defer os.Remove(file.Name())
//	defer file.Close()
//
//	file.WriteString(yaml)
//
//	return
//}
//10.0.91.20  10.0.91.21 10.0.91.22
type ConsuleIpTable struct {
	Ip1 string
	Ip2 string
	Ip3 string
	Env string
}

func main() {
	dd := ConsuleIpTable{Ip1:"10.0.91.20",Ip2:"10.0.91.20",Ip3:"10.0.91.20",Env:"dl"}
	content, err := parseConfFile("/root/work/src/test/consultemfile/pv.yaml", dd)
	if err != nil {
		logs.Error(err.Error())
	}
	fmt.Println("####",content)
	path ,err := newTempFile(content)
	if err != nil {
		logs.Error(err.Error())
		//return "", err
	}
	fmt.Println("##path##",path)

	defer os.RemoveAll(path)

}
