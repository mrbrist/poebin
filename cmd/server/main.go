package main

import (
	"fmt"
	"log"

	"github.com/mrbrist/poebin/internal/r2"
)

func main() {
	r2, err := r2.Setup()
	if err != nil {
		log.Fatal(err)
	}

	build, err := r2.GetBuild("mvz44hssVFeAwspmxNUpxo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(build)
}
