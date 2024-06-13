package auth

import "fmt"

func something() {
	fmt.Println("Is this something that can be accessed outside of this file?")
}

func SomethingElse() {
	something()
}
