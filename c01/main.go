package main

import _1 "GitS/c01/01"

func main() {
	//ans := _1.Subs("abc")
	//ans1 := _1.Permutation("aaa")
	//fmt.Printf("%#v\n", ans)
	//fmt.Printf("%#v\n", ans1)

	//w := []int{1, 2, 3, 4, 5, 6, 7} // 占背包的容量
	//v := []int{2, 9, 8, 7, 8, 5, 4} // 所对应的价值
	//println(_1.MaxValue(w, v, 15))
	//println(_1.DpWays(w, v, 15))

	//println(_1.Queen(8))
	//println(_1.Queen2(8))

	//println(_1.Way1(8, 4, 5, 7))
	//println(_1.RobotDpWays(8, 4, 5, 7))

	arr := []int{5, 10, 50, 100}
	println(_1.CornWay1(arr, 1000))
	println(_1.CornWay2(arr, 1000))
	println(_1.CornWay3(arr, 1000))
	println(_1.CornWay4(arr, 1000))
}

//func modifySlice(innerSlice []string) {
//	fmt.Printf("%p,%v\n", &innerSlice, innerSlice)
//	// 只改变数据而不改变capital会影响到outerSlice
//	innerSlice[0] = "b"
//	// 改变了capital 所以改变指针 copy到一个新的slice 以下修改将不会影响到外部
//	innerSlice = append(innerSlice, "a")
//	innerSlice[1] = "b"
//	fmt.Printf("%p,%v\n", &innerSlice, innerSlice)
//}
//
//func main() {
//	outerSlice := []string{"a", "a"}
//	modifySlice(outerSlice)
//	fmt.Printf("out:%p,%v", &outerSlice, outerSlice)
//}

//func modifySlice(innerSlice []string) {
//	fmt.Printf("out:%p,%v,len: %d,ca: %d\n", &innerSlice, innerSlice, len(innerSlice), cap(innerSlice))
//	innerSlice = append(innerSlice, "a")
//	innerSlice[0] = "b"
//	innerSlice[1] = "b"
//	fmt.Printf("out:%p,%v,len: %d,ca: %d\n", &innerSlice, innerSlice, len(innerSlice), cap(innerSlice))
//}
//
//func main() {
//	outerSlice := make([]string, 0, 3)
//	fmt.Printf("out:%p,%v,len: %d,ca: %d\n", &outerSlice, outerSlice, len(outerSlice), cap(outerSlice))
//
//	outerSlice = append(outerSlice, "a", "a")
//	modifySlice(outerSlice)
//	fmt.Printf("out:%p,%v,len: %d,ca: %d\n", &outerSlice, outerSlice, len(outerSlice), cap(outerSlice))
//}

//func modifySlice(innerSlice *[]string) {
//	fmt.Printf("in:%p,%v,ca: %d\n", &innerSlice, innerSlice, cap(*innerSlice))
//	*innerSlice = append(*innerSlice, "a")
//	(*innerSlice)[0] = "b"
//	(*innerSlice)[1] = "b"
//	fmt.Printf("in:%p,%v,ca: %d\n", &innerSlice, innerSlice, cap(*innerSlice))
//}
//
//func main() {
//	outerSlice := make([]string, 0, 3)
//	fmt.Printf("out:%p,%v,ca: %d\n", &outerSlice, outerSlice, cap(outerSlice))
//
//	outerSlice = append(outerSlice, "a", "a")
//	modifySlice(&outerSlice)
//	fmt.Printf("out:%p,%v,ca: %d\n", &outerSlice, outerSlice, cap(outerSlice))
//}

// 总结：不改变capital，传递的slice能改变源slice的值
// 		改变了capital的值（用append），会指向新的slice需要返回时需要传递指针类型
// 		用指针传递改变了capital则会用指针指向新的slice底层
