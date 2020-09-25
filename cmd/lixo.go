package main

import "fmt"

func main() {
	toChange := "to-change"
	var pointer *string = &toChange
	newPointer :=*pointer
	
	
	teste(toChange)
	teste2(&*pointer)
	teste(newPointer)
	
	fmt.Println(toChange)
	fmt.Println(pointer)

	fmt.Println(&newPointer)
}

func teste(str string){
	fmt.Println(str)
}

func teste2(str *string){
	fmt.Println(str)
	fmt.Println(&str)
}