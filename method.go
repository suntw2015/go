package main

import "fmt"

type student struct{
    name string
    age uint8
    score float64
}

func (stu *student) say() string{
    stu.age ++
    content := fmt.Sprintf("I am %s, %d years old,my score is %f",stu.name,stu.age,stu.score)
    return content
}

func main(){
    stu1 := student{"tom",20,80.5}
    fmt.Println(stu1.say())
    fmt.Println(stu1.say())
}
