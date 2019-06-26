package advance

import "fmt"

func ChannelWithOutBuffers () {
	ch := make(chan string)

	go func (msg string){
		ch <- msg
	}("hello")



	recv := <- ch
	fmt.Println("receiver side : ", recv)

}

func ChannelWithBuffers () {
	ch := make(chan string,1)

	func (msg string){
		ch <- msg
	}("hello")
	recv := <- ch
	fmt.Println("receiver side : ", recv)

}