package main

import (
	"fmt"

	"github.com/miketmoore/go-package-organization-example/concrete"
)

func main() {
	fmt.Println("main")

	c := concrete.New(1, "Mike")
	fmt.Printf("Name: %s, ID: %d\n", c.Name(), c.Id())
}
