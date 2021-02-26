package main

import (
	"fmt"

	"github.com/DataHenHQ/useragent"
)

func main() {
	fmt.Println("Trying out random user agents:")
	for i := 0; i < 20; i++ {
		ua, err := useragent.Desktop()
		if err != nil {
			panic(err)
		}
		fmt.Println(ua)
	}

}
