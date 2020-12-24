package main

import (
	"fmt"

	"github.com/suradidchao/listenfield/repo"
)

func main() {
	fmt.Println("Hello listen field")
	farm := repo.Farm{FarmName: "jia farm"}
	fmt.Printf("%+v\n", farm)
}
