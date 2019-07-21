package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/someJSON")
	//resp, err := http.Get("http://localhost:8080/ping")
	fmt.Println(err)
	fmt.Println(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	//var data interface{}
	//json.Unmarshal(body, &data)
	fmt.Printf("Results: %s\n", body)

}
