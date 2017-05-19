// Copyright 2017 Gruppe 12 IS-105. All rights reserved.
package main

import (
	"os"
	"./boring"
)

func main() {
	msg := os.Args[1]
	boring.Boring01(msg)

}
