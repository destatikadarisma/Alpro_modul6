package main
import "fmt"
func main(){
	var j, v1, v2, hasil int
	fmt.Scan(&v1, &v2)
	hasil = 0
	for j = 1; j <= v2; j++{
		hasil = hasil + v1
	}
	fmt.Println(hasil)

}

//Program Perkalian
// kamus 
//   j, v1, v2, hasil integer
// algoritma
//   input(v1, v2)
//   hasil = 0
//   for j <- for v1 to v2 do
//     hasil = hasil + 1
//   endfor
//endprogram 