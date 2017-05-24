package main

import "fmt"

func main(){
    
    a := []string{"aa","bb","cc","dd"}
    b := a[2:]
    b[0] = "kk"
    c := a[:2]
    c = append(c,"qq")
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

    for k,v := range(a){
        v += "qq"
        fmt.Printf("%d --> %s",k,v)
   }

   fmt.Println(a)
}
