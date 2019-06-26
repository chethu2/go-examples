package basics

import "fmt"

func Factorial (a int) {
	z := fact(a)
	fmt.Printf("Factorial of %d is %d",a,z)
}
func fact (a int) int{
	if a == 1 {
		return 1
	}else {
		return a * fact(a-1)
	}
}

func RecursionPattern ( n int) {
	if n == 1 {
		fmt.Print("*")
	} else {
		for i := 0; i <= n; i++ {
			fmt.Print("* ")
		}
		fmt.Println()
		RecursionPattern(n - 1)
	}
}
