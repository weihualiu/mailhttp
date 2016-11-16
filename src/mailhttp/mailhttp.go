
package main

import (
	"fmt"
	"time"
)

import _ "services"

func main() {
	fmt.Println("test")
	for{
		time.Sleep(time.Second * 10)
	}
}
