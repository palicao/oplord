package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func test(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var i interface{}
	err = json.Unmarshal(body, &i)
	if err != nil {
		panic(err)
	}
	log.Printf("%s %s\n", req.URL, i)
}

func main() {
	http.HandleFunc("/hook", test)
	http.HandleFunc("/delete-hook", test)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
