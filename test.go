package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "133"
	from, err := strconv.Atoi(strings.Split(s, "-")[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(from)
	fmt.Println(strings.Split(s, "-"))
}
