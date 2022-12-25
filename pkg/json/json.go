package json

import (
	"encoding/json"
	"fmt"
)

func Json() {
	// encode bob
	// user struct
	type User struct {
		Name      string
		Bio       string
		Email     string
		Followers int64
	}
	// create a user with the type User
	u := User{
		"Bob",
		"I am bob",
		"bob@bob.com",
		7890,
	}
	// print bob
	fmt.Printf("\nBob: %v\n", u)
	// marshall bob
	marshalled, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	// print marshalled bob
	fmt.Printf("\nMarshalled Bob: %v\n", marshalled)
	// lets take the byte value, and decode it to get bob again!
	// create placeholder
	var decoded User
	// decode marshalled bob into decoded variable
	err = json.Unmarshal(marshalled, &decoded)
	if err != nil {
		panic(err)
	}
	// print decoded bob
	fmt.Printf("\nBob decoded: %v\n", decoded)
	// lets say we didn't know bob was a user
	var decodedUnknown interface{}
	err = json.Unmarshal(marshalled, &decodedUnknown)
	if err != nil {
		panic(err)
	}
	// the output is a map
	fmt.Println("raw decoded:", decodedUnknown)
	// lets get the actual values
	data := decodedUnknown.(map[string]interface{})
	// lets print the data
	for i, j := range data {
		fmt.Println(i, ":", j)
	}
	// bob is back!!
	// lets say we only care about bob's name and email
	// we can receive only the data we need
	type Login struct {
		Name  string
		Email string
	}
	// create a placeholder
	var loginData Login
	err = json.Unmarshal(marshalled, &loginData)
	if err != nil {
		panic(err)
	}
	// print the loginData
	fmt.Println("Login:", loginData)
	// it shows only the name and email!
}
