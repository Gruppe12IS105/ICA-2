// Copyright 2017 Gruppe 12 IS-105. All rights reserved.
// Kode fra https://github.com/uia-worker/is105-ica02/blob/master/boring/boring.go delvis gjennbrukt her.
package boring

import (
	"fmt"
	"math/rand"
	"time"
)



func Boring01(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}

func Boring10(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
