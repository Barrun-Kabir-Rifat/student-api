package main
import "fmt"

func add(a,b int)int {
	return a-b+b*b;
}
func sub(a,b int)int {
	return a-b-a
}
func main(){
	fmt.Println("Here we go , this is the win win situation")

	fmt.Println(add(10,20))
	fmt.Println(sub(10,30))

}