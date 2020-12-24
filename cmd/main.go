package main

import (
	"fmt"

	"github.com/suradidchao/listenfield/entity"
)

func main() {
	fmt.Println("Hello listen field")
	farm := entity.Farm{FarmName: "jia farm"}
	fmt.Printf("%+v\n", farm)
}
