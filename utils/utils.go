package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Reads in a ports.json file and returns a map[key, PortData]
func ReadFile() []byte {
	// Open our jsonFile
	jsonFile, err := os.Open("ports.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened ports.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
	// // we initialize our Users array
	// var ports map[string]client.PortData

	// // we unmarshal our byteArray which contains our
	// // jsonFile's content into 'users' which we defined above
	// err = json.Unmarshal(byteValue, &ports)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// return ports
}
