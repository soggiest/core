package main

import (
	"net/http"
	"encoding/json"
)

type Operation struct {
	Name	string	`json:"name"`
}

func listHandler(writer http.ResponseWriter, response *http.Request) {
	//TODO shouldn't be hardcoded
	type Operations []Operation

	operations := Operations {
		Operation{ Name: "/list" },
		Operation{ Name: "/quit" },
	}

	json.NewEncoder(writer).Encode(operations)

}

func main() {
	http.HandleFunc("/list", listHandler)
	http.ListenAndServe(":9993", nil)
}
