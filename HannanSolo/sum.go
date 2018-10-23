package main

import (
	"fmt"
	"os"
	"strconv"
)

//Takes in 2 number args and adds them together
func main() {
	var var1, var2 int
	var err error

	var1, err = strconv.Atoi(os.Args[1])
	var2, err = strconv.Atoi(os.Args[2])
	if err != nil {
		return
	}

	fmt.Println(var1 + var2)
}
