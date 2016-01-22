package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"os"
)

type Operation struct {
	Name	string	`json:"name"`
}

func main() {
	bindHandlers()
	printWelcomeMessage()
	startServer()
}

func bindHandlers() {
	http.HandleFunc("/list", listHttpHandler)
	http.HandleFunc("/quit", quitHttpHandler)
}

func printWelcomeMessage() {
	fmt.Println("Now listening on :9993")
}

func startServer() {
	http.ListenAndServe(":9993", nil)
}

func listHttpHandler(writer http.ResponseWriter, response *http.Request) {
	json.NewEncoder(writer).Encode(operationList())
}

func operationList() ([]Operation) {
	type Operations []Operation

	operations := Operations {
		Operation{ Name: "/list" },
		Operation{ Name: "/quit" },
	}

	return operations
}

//TODO figure out how to get the writer to flush before the aplication shuts down
func quitHttpHandler(writer http.ResponseWriter, response *http.Request) {
	defer shutdown()
	sendStatusOK(writer)
	sendByeMessageToClient(writer)
}

func sendStatusOK(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusOK)
}

func sendByeMessageToClient(writer http.ResponseWriter) {
	writer.Write([]byte("bye"))
}

func sendByeMessageToConsole() {
	fmt.Println("See you next time");
}

func shutdown() {
	os.Exit(0)	// 0 == everything is ok
}

