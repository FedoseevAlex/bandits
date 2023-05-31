package main

import (
	"fmt"

	"github.com/FedoseevAlex/bandits/internal/version"
)

func main() {
	ver, err := version.GetCurrentVersion()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ver)
}
