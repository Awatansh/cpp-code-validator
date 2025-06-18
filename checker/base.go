package main

import (
	"fmt"
)

func main() {
	runner()
	xy := judge()
	if xy {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
