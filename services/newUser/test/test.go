package main

import (
	"fmt"

	"github.com/koding/kite"
)

func main() {
	k := kite.New("exp2", "1.0.0")

	// Connect to our math kite
	worker := k.NewClient("http://localhost:3635/kite")
	worker.Dial()

	response, _ := worker.Tell("new", "namE", "emaiL", "passworD")
	fmt.Println("result:", response.MustString())
}
