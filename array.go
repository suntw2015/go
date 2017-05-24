package main

import "fmt"

func main(){
    var a[2] string
    b := [...]int{1,2,5,1,9,54}

    a[0] = "hello"
    a[1] = "world"
    fmt.Println(a,b)
}
