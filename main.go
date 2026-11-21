package main

import "fmt"
import "net/http"

func main() {
handlerIndex := func(w http.ResponseWriter, r *http.Request) {
fmt. Fprintln(w, "halo nama saya Ridho")
}
http.HandleFunc("/", handlerIndex)

fmt.Println("server started at 43.173.13.18:9021")
http.ListenAndServe(":9021", nil)

}
