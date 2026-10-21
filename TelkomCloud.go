package main
import "fmt"
func main(){
	var n, counter, k, i int
	fmt.Scan(&n)
	k = 1 
	i = n
	for counter = 1; counter <= n; counter++{
		fmt.Println("Request", k, "diperoses dengan prioritas", i)
		k = k + 1
		i = i - 1
	}
}