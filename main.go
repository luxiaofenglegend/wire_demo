package main

import (
	"fmt"
	"wire_demo/demo"
)

func main() {
	b := demo.InitializeBroadCast()

	b.Start()

	inject, err := demo.BuildInjector()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", inject)
}
