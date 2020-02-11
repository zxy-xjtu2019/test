package main

import "fmt"

func main(){
	for i, c := range "go0" {
		fmt.Println(i, c)
	}
}
