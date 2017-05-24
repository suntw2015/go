package main

import "fmt"

type student struct{
    name string
    age int
}

func main(){
    stu := student{"tom",20}
    stu.name= "Tom";
    fmt.Println(stu);
}
