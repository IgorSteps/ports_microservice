package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Reads in a ports.json file and returns slice of bytes
func ReadFile() []byte {
	// Open our jsonFile
	jsonFile, err := os.Open("ports.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println("Successfully Opened ports.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}
