package basics

import "fmt"

func Array1() {
	var oneToTen[10] int
	for i:= 1; i<=10 ; i++ {
		oneToTen[i-1] = i ;
	}
	fmt.Println("array elements:",oneToTen)
}

func Array2() {
	var days [7] string
	days[0] = "MONDAY"
	days[1] = "TUESDAY"
	days[2] = "WEDNESDAY"
	days[3] = "THURSDAY"
	days[4] = "FRIDAY"
	days[5] = "SATURDAY"
	days[6] = "SUNDAY"

	for i := 0 ; i<7 ; i++ {
		fmt.Println("Day " ,i+1, ":",days[i])
	}
}

func Array3 () {
	var matrix[2][3] int
	for i :=0; i<2; i++ {
		for j:=0;j<3 ;j++{
			matrix[i][j] = 1
		}
		fmt.Println()
	}
	fmt.Println(len(matrix))
	fmt.Println(matrix)
}

func Array4 (){
	var row, col int
	fmt.Print("Please enter the value of rows and columns")
	fmt.Scan(&row,&col)
	matrix := make([][] int, row)
	for i:= 0; i<row ;i++{
		matrix[i] = make([]int , col)
	}
	for i:=0; i<row;i++{
		for j:=0;j<col;j++{
			if i == j{
				matrix[i][j] = 1
			}else {
				matrix[i][j] = 0
			}

		}
	}
	fmt.Println(len(matrix))
	fmt.Println(matrix)
}