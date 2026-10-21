package main

import "fmt"

func main(){
	var n, i, hasil int
	fmt.Scan(&n)
	hasil = 1
	for i = 1; i <= n; i++{
		hasil = hasil * i
	}
	fmt.Println(hasil)
}

// program faktorial
// kamus
//   n,i,hasil integer
//algoritma
//   input(n)
//   hail = 1
//   for i <- 1 to n do
//      hasil = hasil * 1
//   endfor
//endprogram