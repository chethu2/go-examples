package basics

import "fmt"

func VariadicSum(){
	sum(1,2,3)
	n := [] int{1,2,3}
	sum(n...)
}
func sum (nums ...int){
	total := 0
	for _, n := range nums{
		total = +n
	}
	fmt.Printf("total : %d\n", total)
}
