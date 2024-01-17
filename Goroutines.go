package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(750 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

//	func alphabets() {
//		for i := 'a'; i <= 'e'; i++ {
//			time.Sleep(1000 * time.Millisecond)
//			fmt.Printf("%c ", i)
//		}
//	}
func main() {
	go numbers()
	//go alphabets()
	//time.Sleep(5100 * time.Millisecond)
	fmt.Println("main terminated")
	time.Sleep(5100 * time.Millisecond)
}
