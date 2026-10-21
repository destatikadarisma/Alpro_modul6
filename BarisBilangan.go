package main
import "fmt"
func main(){
	var a, b int
	var j int
	fmt.Scan(&a,&b)
	for j = a; j <= b; j++{
		fmt.Println("baris bilangannya =", j)
	}
}

//Prigram BarisBilangan
//kamus
//  a,b integer
//  j integer
//algoritma
//  input(a,b)
//  for j <- a to b do
//  output(baris bilangan =, j)
//endfor
//endprogram