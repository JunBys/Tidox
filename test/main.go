package main

import (
	"tidox/api"
)

func main() {
	a := api.NewServer()
	a.Run(":8888")
}
