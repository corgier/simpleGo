package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Println(freezingF, fToC(freezingF))
	fmt.Println(boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
