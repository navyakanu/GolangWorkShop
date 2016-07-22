package main

import (
	"fmt"
	"errors"
)







func main() {

	var variable string="Golang Workshop"
	fmt.Println(variable)
	fmt.Println("Main.go executing\n")
	x:=20
	fmt.Println(x)
	y,err:=ErrorExample()
	fmt.Println("Printing error from function")
	fmt.Println(err)
	fmt.Println(y)

}



func ErrorExample() (int,error) {

if 2 != 1 {
	return 2,errors.New("Not equal- Error ")
} else {
	return 0,nil
}
}
