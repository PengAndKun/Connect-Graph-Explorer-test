package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/tidwall/pretty"
)

type uniswapStruct struct {
	Query string `json:"query"`

	//variables        interface `json:"variables"`
}
type urlHTTP struct {
	Queryaddr string `json:"queryaddr"`
}
type result struct {
	Data interface{} `json:"data"`
}

func main() {

	cs := uniswapStruct{}
	uH := urlHTTP{}
	//var cs []uniswapStruct = make([]uniswapStruct, 0)
	data, err := ioutil.ReadFile("queryGraphQL.json")
	if err != nil {
		panic(err)
	}
	urldata, err := ioutil.ReadFile("urlHttp.json")
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(data))
	// fmt.Println(string(urldata))

	err = json.Unmarshal(data, &cs)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(urldata, &uH)
	if err != nil {
		panic(err)
	}
	// fmt.Println(cs)
	// fmt.Println(uH)
	// fmt.Println(111111)
	// fmt.Println(uH.Queryaddr)

	//json

	response, err := http.Post(
		uH.Queryaddr,
		"application/json",
		strings.NewReader(string(data)),
	)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(body))
	// // b, err := json.Marshal(body)
	// fmt.Println(body)
	// fmt.Println(reflect.TypeOf(body))
	resultDate := result{}
	//fmt.Println(reflect.TypeOf([]byte(string(body))))
	err = json.Unmarshal(body, &resultDate)
	if err != nil {
		panic(err)
	}
	fp, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	a := pretty.Pretty(body)
	//fmt.Println(a)
	_, err = fp.Write(a)
	if err != nil {
		panic(err)
	}

}
