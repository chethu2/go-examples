package advance

import "fmt"

func Sorting () {
	slice := make ([]int , 5)
	for i := 0; i<5; i++ {
		slice[i] = i+1
	}
	fmt.Println("loop initalization : ", slice)

	slice = append(slice,6,7,8,9,10)
	fmt.Println("after appending: ", slice)

	slice1  := make ([] int , len(slice))
	copy(slice1,slice)
	fmt.Println("copying to slice1: ", slice1)

	slice = append(slice[:4], slice[:5]...)
	fmt.Println("cutting the entry 5: ", slice)

	x := cutByIndex(2,slice)
	fmt.Println("after removing 2nd index from the slice",x)

}

func cutByIndex(i int, slice []int )  []int{
	slice = append(slice[:i],slice[i+1:]...)
	return slice
}
