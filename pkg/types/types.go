package types

import (
	"fmt"
)
type correct struct {
	name string
	email string
}
type incorrect struct {
	pass string
	age uint
}
func Assertion() {
	// to access an interface's value only if it is a certain type, use type Assertion
	// first declare an interface
	var test interface{} = 42
	// lets check if test is an int
	// syntax: val, ok = interface.(type)
	num, isInt := test.(int)
	fmt.Println(num, isInt)
	// isInt is true because the the test is an integer
	// change test value to a string
	test = "not an integer"
	// try checking if it is an integer again
	num, isInt = test.(int)
	fmt.Println(num, isInt)
	// isInt is false because test is not an int
	// you can do this with other types as well - like structs
	var testStruct interface{} = correct {
		"billy",
		"billy@email.com",
	}
	val, ok := testStruct.(correct)
	fmt.Println(val, ok) // all good
	var incorrectTestStruct interface{} = incorrect {
		"1234",
		32,
	}
	// you can get the type of an interface:
	switch t := incorrectTestStruct.(type) {
		case correct: 
			fmt.Printf("%v is correct!!\n\n", t)
		case incorrect:
			fmt.Printf("%v is wrong!!\n\n", t)
		default:
			fmt.Println("unknown type")
	}
	// if there is no  "ok" value declared: 
	// if it is not having the type specified, it will panic
// 	val = incorrectTestStruct.(correct) //panic 

}
