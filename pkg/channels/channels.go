package channels;
// necessary packages
import "fmt"

func Channels() {
    fmt.Println("Products:")
    // channels send and receive data
    // you can receive multiple values from channels
    // used with go routines
    // sample slice
    slice1 := []int{10, 3, 7, 1, 2, 9} // 3780
    slice2 := []int{6, 7, 2, 3, 1, 9} // 2268
    // make an empty channel
    c := make(chan int)
    // call "product" with goroutines
    go product(slice1, c)
    go product(slice2, c)
    // receive the two products with channels
    prod1, prod2 := <-c, <-c
    fmt.Println(prod1, prod2)
}
// calculate product of an array
func product(nums []int, c chan int) {
    prod := 1
    for _, v := range nums {
        prod *= v
    }
    // put the product in the channel i gave it
    c <- prod
}

