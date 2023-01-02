package context;
// necessary packages
import (
    "fmt"
    "net/http"
    "time"
)
func Context() {
    fmt.Println("Context:")
    http.HandleFunc("/example", example)
    http.ListenAndServe(":8070", nil)
}
// request handler
func example(write http.ResponseWriter, req *http.Request) {
    // context variable
    ctx := req.Context()
    fmt.Println("received request")
    // only say it ended at the end
    defer fmt.Println("server handler ended")
    select {
        // dont immediately stop
    case <-time.After(5 * time.Second):
        fmt.Fprintf(write, "example\n")
        // if the context is done ( e.g: <C-c>)
    case <-ctx.Done():
        // print the error
        err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        // write the error on the request
        http.Error(write, err.Error(), internalError)
    }
}
