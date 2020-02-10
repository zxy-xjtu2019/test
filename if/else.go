package main

import "fmt"

func main(){
	if 7%2 == 0{
		fmt.Println("7 is even")
	}else {
		fmt.Println("7 is odd")
	}
	for i := 1;i<=20; i++{
		if i%2 == 0 {
			fmt.Println("ooo")
		}else {
			fmt.Println("nonono")
		}
	}
}
