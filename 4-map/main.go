package main

import "fmt"

func main() {
	colors := map[string]string{
		"blue":   "Lan",
		"yellow": "Huang",
		"green":  "Lv",
	}

	for k, v := range colors {
		fmt.Println(k, v)
	}
}
