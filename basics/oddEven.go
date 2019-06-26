package basics

import "fmt"

func Odd(a int) {
	for i := 1; i<=a ; i++ {
		if i%2 == 0{
			continue;
		}
		fmt.Println("odd :",i)
	}
}

func Even(a int) {
	for i := 1; i<=a ; i++ {
		if i%2 != 0{
			continue;
		}
		fmt.Println("even :",i)
	}
}
