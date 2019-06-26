package basics

import "fmt"

func Pattern(){
	for i:=1; i<9 ; i++ {
		for j:=0;j<i; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Pattern1(){
	for i:=9; i>0 ; i-- {
		for j:=0;j<i; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Pattern2 () {
	for i := 1; i <= 11 ; i++ {
		if i < 7 {
			for j := 0; j < i; j++ {
				fmt.Print("*")
			}
		} else{
			x := (i - 6) * 2
			for j:=i; j>x;j--{
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
