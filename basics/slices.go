package basics

import "fmt"

func Slice1() {
	slice := make ([]int, 0)
	fmt.Println("before append slice --> ",slice)
	slice = append(slice, 0,1,2,3,4,5,6,7,8,9)
	fmt.Println("after append slice --> ",slice)
	fmt.Println("another way to acess full slice  slice[:] -->",slice[:])
	fmt.Println("from index 2 : slice[2] -->",slice[2:])
	fmt.Println("till index 2 -->",slice[:2])
	fmt.Println("from index 2 to index 5 -->",slice[2:5])
}

func Slice2 (){
	slice := make([] int , 3, 10)
	fmt.Println("slice after initialization -->",slice)
	slice[2]= 9
	fmt.Println("slice after assignment --> ",slice)
	slice = append(slice, 1,2,3,4,7,8,9,5,6,7)
	fmt.Println("slice after append --> ",slice)
}

func Slice3 () {
	slice := make ([] string , 3)
	fmt.Println("slice after initialization -->",slice)
	slice[0] = "Monday"
	slice = append(slice, "Tuesday", "Wednesday", "Thursday","Friday","Saturday", "Sunday")
	fmt.Println("slice after append -->",slice)
}
