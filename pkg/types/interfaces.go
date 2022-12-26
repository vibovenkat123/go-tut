package types

import (
	"fmt"
)
type num interface {
	Introduce()	string
}
// create a local float type
type coolFloat float64
type uncoolInt int32
// intro for a float
// specify the receiver (coolFloat) which means no other type can access Introduce
// which also menas that num can only be a coolFloat
func (num coolFloat) Introduce() string {
	return fmt.Sprintf("this is %v, %v doubled is %v", num,num,  num * 2)
}
// interfaces are an easy way to contain methods of a struct
func Interfaces() {
	// if we get no error on the following lines, it means that "number" implements num,
	// which means it implements Introduce()
	var number num
	number = coolFloat(2.56)
	fmt.Println(number.Introduce())
	// you get an error if you change "number" to not be a coolFloat
	// number = uncoolInt(3) // will produce error
}
