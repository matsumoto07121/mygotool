package mygotool
//package main

import (
	"fmt"
	"errors"
)

type Greet interface {
	Hello(string)(string,error)
	Thankyou(string)(string,error)
}

type Japanese struct {}
type English struct {}

func (m Japanese) Hello(name string) (string, error) {
	if name == "" { return "", errors.New("Name naiyo") }

	return fmt.Sprintf("Konichiwa %s!", name), nil
}
func (m English) Hello(name string) (string, error) {
	if name == "" { return "", errors.New("Name is empty") }
	return fmt.Sprintf("Hello %s!", name), nil
}
func (m Japanese) Thankyou(name string) (string , error){
	if name == "" { return "", errors.New("Name naiyo") }
	return fmt.Sprintf("Arigato %s!", name), nil
}
func (m English) Thankyou(name string) (string, error)  {
	if name == "" { return "", errors.New("Name is empty") }
	return fmt.Sprintf("Thank you, %s!", name), nil
}

/*
func main(){
	a := []Greet{ Japanese{}, English{} }
	for _, v := range a{
		fmt.Println(v.Hello("Mr.A"))
	}
}*/


// Hello returns a greeting for the named person.
// right way to use method?? (todo1
func Hello(name string, speaktype int) (string, error) {
	switch(speaktype){
	case 1:
		return Japanese.Hello(Japanese{}, name)
	case 2:
		return English.Hello(English{}, name)
	default:

	}
	return "..." + name, errors.New("No speaker type")
}


// how to test?? todo2
/*
func main() {
	fmt.Println(Hello("hogehoge", 1))
	fmt.Println(Hello("hogehoge", 2))

}
*/
