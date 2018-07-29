package main

import (
	"fmt"

	"github.com/gen0cide/privcheck"
)

func main() {
	if privcheck.IsAdmin() {
		fmt.Println("WE ARE ADMIN")
	} else {
		fmt.Println("WE ARE *NOT* ADMIN")
	}
}
