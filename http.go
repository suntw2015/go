package main

import(
    "fmt"
    "net/http"
)

type server struct{}

func (s server)ServeHTTP(w http.ResponseWriter,r *http.Request){
    fmt.Fprint(w,r)
}

func main(){
    var s server
    err := http.ListenAndServe("localhost:5000",s)
    if err != nil{
        fmt.Println(err)
    }
}
