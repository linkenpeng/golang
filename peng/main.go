package main

import (
	"fmt"
	"os"
)

func main() {
	os.RemoveAll("tb")
	fmt.Println("finish.")
}
