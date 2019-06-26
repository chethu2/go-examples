package advance

import "fmt"

type profile struct {
	name string
	contactNumber string
	emailId string
	address string
}

func (p * profile ) getName() string{
	return p.name
}

func (p *profile) getContact () map[string]string {
	return map[string]string{"emailId":p.emailId,"contactNumber":p.contactNumber}
}

func (p *profile) getAddress() string {
	return p.address
}

func MethodEx () {
	p := profile{address:"Baker's Street",contactNumber:"XXXXX-XXXXX",emailId:"sherlock@outlook.com",name:"Sherlock holmes"}
	fmt.Println("Name : ", p.getName())
	fmt.Println("Contact :", p.getContact())
	fmt.Println("Address :", p.getAddress())

}

