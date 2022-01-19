package main

import "fmt"
import "rsc.io/quote"

func  hello()string {
	 return quote.Hello()

}
func  glass()string {
	return quote.Glass()

}
func gone()string {
	return quote.Go()

}
func  opt()string {
	return quote.Opt()

}
func main(){
	fmt.Println(hello())
	fmt.Println(glass())
	fmt.Println(gone())
	fmt.Println(opt())
}
