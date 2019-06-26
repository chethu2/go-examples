package advance

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Write1()  {

	data := [] byte("Hi, Hello \nGood Morning!!")

	err := ioutil.WriteFile("hello1.txt", data, 0666)
	if err != nil {
		fmt.Println("error while writing into file ", err)
	}
	fmt.Println("done")
}

func Write2 () {
	f , err := os.OpenFile("hello.txt",os.O_APPEND|os.O_WRONLY,0777)
	defer f.Close()
	if err != nil {
		fmt.Println("error while opening file ",err)
	}
	b := [] byte ("\nHi Hello !!")
	f.Write(b)
	fmt.Println("File descriptor stats",f)
}
