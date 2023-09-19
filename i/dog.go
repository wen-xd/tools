package i

import "fmt"

type Dog struct {
}

func (d Dog) Eat() {
	fmt.Println("dog...")
}
