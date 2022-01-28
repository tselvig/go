package main

// Tanner Selvig
// Assignment 1.5 - Reads a JSON file 
// Run this program with the following command: go run read-json1.go FileToRead.json NewFileName

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type TransactionType struct {
	TxHash string
	TxIn  int
	TxOut int
}

func main() {
	// Stores the two command line arguments
	var fileNames = os.Args[1:]

	// Takes the first argument and reads the file with that name
	data, err := ioutil.ReadFile(fileNames[0])
	_ = err

	// Transforms data into a usable format
	var tt TransactionType
	err = json.Unmarshal(data, &tt)
	_ = err
	buf, err := json.MarshalIndent(tt, "", "\t")
	_ = err

	// Prints the data
	fmt.Printf("Data: %s\n", buf)

	// Reads the second command line argument and outputs a file with that name, and the data read in
	err2 := ioutil.WriteFile(fileNames[1] + ".json", buf, 0644)
	if err2 != nil {fmt.Printf("Error writing file")}

}