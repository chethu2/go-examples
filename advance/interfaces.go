package advance

import "fmt"

type shape interface {
	getArea() float64
	getPerimeter() float64
}

type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) getArea() float64{
	return r.height*r.width
}

func (r Rectangle) getPerimeter() float64 {
	return 2*(r.height+ r.width)
}
func InterfacesDemo() {

	var shapes shape
	rect := Rectangle{height:4,width:5}
	shapes = Rectangle{width:4,height:5}
	fmt.Println("Rectangle Area : ",rect.getArea())
	fmt.Println("Rectangle perimeter : ",rect.getPerimeter())
	fmt.Println("SHAPE Area : ",shapes.getArea())
	fmt.Println("SHAPE perimeter : ",shapes.getPerimeter())
}

