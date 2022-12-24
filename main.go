package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

var (
	ports = ReadFile()
)

// port data
type PortData struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

func ReadFile() map[string]PortData {
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

	// we initialize our Users array
	var ports map[string]PortData

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &ports)
	if err != nil {
		fmt.Println(err)
	}
	return ports
}

func getPorts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ports)
}

func postPorts(c *gin.Context) {
	var newPort map[string]PortData

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newPort); err != nil {
		return
	}

	// Add the new port to the ports map.
	maps.Copy(ports, newPort)
	c.IndentedJSON(http.StatusCreated, newPort)
}

func main() {
	router := gin.Default()
	router.GET("/ports", getPorts)
	router.POST("/ports", postPorts)

	router.Run("localhost:5080")
}
