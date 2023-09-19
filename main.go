package main

import "fmt"

//func main() {
//	x := i.Total{C: i.Dog{}}
//	x.C.Eat()
//}

//func main() {
//	fmt.Println(test1())  // 0 0
//	fmt.Println(test2())  // 3 3
//	fmt.Println(test3())  // 0 4
//	fmt.Println(test4())  // 0 5
//}
//
//func test1() (v int) {
//	defer fmt.Println(v)
//	return v
//}
//func test2() (v int) {
//	defer func() {
//		fmt.Println(v)
//	}()
//	return 3
//}
//func test3() (v int) {
//	defer fmt.Println(v)
//	v = 3
//	return 4
//}
//func test4() (v int) {
//	v = 5
//	/**
//	defer语句中的函数参数在defer语句出现时就会计算并保存，而不是在实际执行时计算。这意味着defer语句中的函数会使用当时的变量值，而不是最终的值，所以需要注意闭包函数中变量的值可能会发生改变。
//	*/
//	defer func(n int) {
//		fmt.Println(n)
//	}(v)
//	// 不会对defer的值产生影响
//	v = 10
//	return 5
//}

func main() {
	i := foo()
	fmt.Println(*i)
	println(max(1, 2, 3))
}

func foo() *int {
	// x 发生了逃逸 本来在函数执行结束后就会被直接销毁
	x := 1
	return &x
}
