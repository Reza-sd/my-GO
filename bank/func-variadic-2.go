package main

import "fmt"

func main() {
	sayHello("mioo")
	sayHello()
	sayHello("Sammy1")
	sayHello("Sammy2", "Jessica", "Drew", "Jamie")
}

func sayHello(names ...string) {
	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}
}
