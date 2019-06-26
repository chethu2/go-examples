package basics

import (
	"fmt"
	"sort"
)

func Sorting() {
	a := [] int{1, 4, 2, 8, 2, 6}
	fmt.Println("before sorting :", a)
	sort.Ints(a)
	fmt.Println("after sorting :", a)

	b := [] string{"z", "g", "a", "t", "m"}
	fmt.Println("before sorting :", b)
	sort.Strings(b)
	fmt.Println("after sorting :",b)
}
