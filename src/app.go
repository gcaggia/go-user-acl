package main

import (
	Router "github.com/gcaggia/go-user-acl/src/modules"
)

func main() {
	r := Router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
