package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	num, err := strconv.ParseInt("1000000", 10, 64)

	if err == nil {
		fmt.Println(num)
	} else {
		fmt.Println(err.Error())
	}

	val := strconv.FormatInt(1000000, 10) // 10 decimal 8 octal 16 hexadecimal
	fmt.Println(val)

}
