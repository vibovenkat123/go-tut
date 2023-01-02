package routines;
import (
    "fmt"
    "time"
)
func test(x string) {
    for i := 0; i < 3; i ++ {
        // need to sleep so that main function doesnt end to early
        time.Sleep(200 * time.Millisecond)
        fmt.Println(x) 
    }
}
func Routines() {
    // goroutines run all at once (concurrently)
    go test("routine")
    test("normal")
}

