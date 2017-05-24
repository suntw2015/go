package main

import "fmt"

func add(x,y int) int{
	return x+y
}

func func2(a,b string)(string,string){
	return b,a
}

func func3(x,y int)(a int){
	a = x + y
	return
}

func main(){
    var i,j int = 1,3;
    k := 34;
    fmt.Println(i,j,k)
	fmt.Println(add(1,4))
	fmt.Println(func2("aa","bb"))
	fmt.Println(func3(3,6))
}
