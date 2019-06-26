package basics

import "fmt"

func FindMe( a int) {

	switch a {
	case 1 :
		fmt.Print("MONDAY")
	case 2 :
		fmt.Print("TUESDAY")
	case 3 :
		fmt.Print("WEDNESDAY")
	case 4 :
		fmt.Print("THURSDAY")
	case 5 :
		fmt.Print("FRIDAY")
	case 6 :
		fmt.Print("SATURDAY")
	case 7 :
		fmt.Print("SUNDAY")
	default:
		fmt.Println("invalid Date")
	}

}
