package advance

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile1 (){
	data, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("error while reading the file",err)
		return
	}
	fmt.Println("byte data",data)
	fmt.Println("string data:\n",string(data))
}

func ReadFile2() {
	data , err := os.Open("hello.txt")
	if err != nil {
		fmt.Println("error while reading the file",err)
		return
	}

	b := make ([]byte , 10)
	for {
		r, _ := data.Read(b)
		if r == 0{
			break
		}
		fmt.Println("byte data :", r)
		fmt.Println("string data :", string(b))
	}

}
