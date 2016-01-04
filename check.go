package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func isJSON(s string) bool {
	var js map[string]interface{}
	fmt.Print(json.Unmarshal([]byte(s), &js))

	return true

}

func isJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil

}

func main() {

	buffer := new(bytes.Buffer)

	resp, err := http.Get("http://127.0.0.1:9993/list")
	if err != nil {
	}
	defer resp.Body.Close()
	buffer.ReadFrom(resp.Body)
	//str := buffer.String()

	newstr = "{\"name\":\"/list\"}/" //,{\"name":"/quit\"}"
	test := isJSONString(newstr)
	fmt.Print(str)
	if test == false {
		fmt.Print(test)
	}
	//body, err := ioutil.ReadAll(resp.Body)

	//isjson := isJSON()

	//Output: true
}
