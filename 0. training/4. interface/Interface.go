package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return (r.height + r.width) * 2
}

func (r rectangle) volume() float64 {
	return (r.height * r.height * r.width)
}

func main() {
	var calrect shape = rectangle{height: 4, width: 5}
	fmt.Println(calrect.area())
	fmt.Println(calrect.perimeter())

	//interface type assertion --> calling functin that hasnt been declared inside the shape interface
	fmt.Println(calrect.(rectangle).volume())

	//random interface --> can be used to store all kind of datatype
	var randomvalues interface{}

	randomvalues = "string"
	randomvalues = 1
	randomvalues = []string{"amanda", "gabriel"}
	randomvalues = true
	fmt.Println(randomvalues)

	//assertion random interface
	var v interface{}
	v = 20
	v = v.(int) * 9 //--->v need to be declared as int or other so you can do operation on them

	fmt.Println(v)
}
