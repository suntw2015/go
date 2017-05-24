package main

import "fmt"

func main(){
    for i := 1;i<10;i++{
        fmt.Println(i)
    }

    sum := 10
    for sum<100{
        sum += 20
    }

    fmt.Println(sum)
}
