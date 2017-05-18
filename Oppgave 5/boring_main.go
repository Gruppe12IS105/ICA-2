// Copyright 2017 Gruppe 12 IS-105. All rights reserved.
package main

import (
	"os"

	"./boring"
	//"fmt"
)

func main() {
	msg := os.Args[1]
	boring.Boring01(msg)

	/*
	  c := make(chan string) //oppretter channel c i denne filen

	  go boring.Boring10("Index nr: ", c) //starter func Boring10 i den andre koden
	  for i := 0; i < 1; i++{             //denne for-løkken kjører parallellt med Boring10
	    fmt.Printf("Test er: %q\n", <-c)
	  }
	  fmt.Println("Neiass, dette gidder jeg ikke")
	*/
}
