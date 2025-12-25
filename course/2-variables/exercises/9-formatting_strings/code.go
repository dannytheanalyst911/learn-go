package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s! Your open rate is %.1f%%.", name, openRate)

	// don't edit below this line

	fmt.Println(msg)
}
