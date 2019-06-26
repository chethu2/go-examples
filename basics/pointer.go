package basics

import "fmt"

func Pointer1() {

	var a * int
	b := 0
	a = &b
	fmt.Println("value *a = b : ", *a)
	fmt.Println("address of b is a :",a)

}

func Pointer2 () {
	a := 5
	increment(a)
	fmt.Println("internal increment of a :",a)
	incrementPointer(&a)
	fmt.Println("internal increment of a :",a)

}

func increment (a int){
	a++
}
func incrementPointer(a* int){
	*a++
}
