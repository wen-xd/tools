package i

import "fmt"

type Chi interface {
	Eat()
}
type Total struct {
	C Chi
}

type Person struct {
}

func (p Person) Eat() {
	fmt.Println("person...")
}
