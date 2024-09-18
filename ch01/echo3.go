package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	strings.Join(os.Args[1:], " ")
	fmt.Println(time.Since(start))
}
