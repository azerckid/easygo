package main

import "fmt"

func multyply(a, b int) (int, string) {
	return a * b, "hello world"
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	name := "Hello world"
	bunho, kul := multyply(2, 3)
	fmt.Println(name, bunho, kul)
	repeatMe("gogo", "tomson", "happy", "sumok")
}
