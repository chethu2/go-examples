package advance

import "fmt"

func Map () {
	var a map[string] int
	a = map[string]int{"a":1,"b":2}
	fmt.Println("value of the map after initialization ",a)
	fmt.Println("value of a[\"a\"] is : ",a["a"])
	b, isPresent := a["b"]
	fmt.Printf("value of b : %d in the map is and isBPresent : %t \n",b,isPresent)
	delete(a,"a")
	fmt.Println("value of the map after deletion ",a)
}
