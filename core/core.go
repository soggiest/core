package main

import (
	"encoding/json"
	"fmt"
//	"io"
	"net/http"
	"os"
//	"strings"
)

type Operation struct {
	Name string `json:"name"`
}

func listHandler(writer http.ResponseWriter, response *http.Request) {
	//TODO shouldn't be hardcoded
	type Operations []Operation

	operations := Operations{
		Operation{Name: "/list"},
		Operation{Name: "/quit"},
	}

	json.NewEncoder(writer).Encode(operations)

}

func getPid(writer http.ResponseWriter, response *http.Request) {
	var pid int

	pid = os.Getpid()

	pidMessage := fmt.Sprintf("%d", pid)

	//io.WriteString(writer, pidMessage)

	json.NewEncoder(writer).Encode(pidMessage)

}

func quitCore(writer http.ResponseWriter, response *http.Request) {
	os.Exit(0)

}

func main() {
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/pid", getPid)
	http.HandleFunc("/quit", quitCore)
	http.ListenAndServe(":9993", nil)

	//if strings.Compare(message, "Quit") {
	//	os.Exit()
	//}

}
