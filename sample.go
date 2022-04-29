package main
import "fmt"

func main() {
	adder(add)

}

//so here we can pass a function too as parameter for another function. So for adder fn we pass add fn as parameter.
//Just like how we actually specify data type of parameter in fn header, we specify that we are passing a fn
// that takes in 2 int parameters and rweturns one int parameter

type myFuncDataType func(int,int)int
func adder(f myFuncDataType ) {
	val := f(3,5)
	fmt.Println(val)
}

func add(i int,j int) int {
	return i+j
}