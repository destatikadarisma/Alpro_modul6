package main
import "fmt"
func main(){
	var n,i, hasil int
	fmt.Scan(&n)
	hasil = 0
	for i = 1; i <= n; i++{
		hasil = 2*i - 1
		fmt.Println("angka ganjilnya terdiri dari", hasil)
	}

}