// https://tutorialedge.net/golang/parsing-json-with-golang/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Users struct which contains an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name, a type, and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	source := "./users.json"

	jsonFile, err := os.Open(source)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully opened \"%v\"\n", source)

	// defer the closing of jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read opened file as byte array
	byteArray, _ := ioutil.ReadAll(jsonFile)

	// initialize Users array
	var users Users

	// unmarshall byteArray into users
	json.Unmarshal(byteArray, &users)

	// iterate through every user within users and print out the details
	for i := 0; i < len(users.Users); i++ {
		fmt.Printf("User name: %v\n", users.Users[i].Name)
		fmt.Printf("User type: %v\n", users.Users[i].Type)
		fmt.Printf("User age: %v\n", users.Users[i].Age)
		fmt.Printf("Facebook URL: %v\n", users.Users[i].Social.Facebook)
	}
}
