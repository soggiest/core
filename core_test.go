package main

import "testing"

//totally trivial, I know
func TestOperationList(t *testing.T) {
	var operations []Operation = operationList()
	if operationLength := len(operations); operationLength != 3 {
		t.Error("Expected 2 operations in list but got ", operationLength)
	}
}
