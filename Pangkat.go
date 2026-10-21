package main 
import "fmt"
func main(){
	var basis, pangkat,hasil, i int
	fmt.Scan(&basis, &pangkat)
	hasil = 0
	hasil = basis
	for i = 1; i < pangkat; i++{
		hasil = hasil*basis
	}
	fmt.Println(hasil)

}
// Program Perkalian
// kamus
//   basis, pangkat, hasil integer
// algoritma
//   input(basis, pangkat)
//   hasil = 0
//   hasil = basis
//   for i <- 1 < pangkat
//     hasil= basis*hasil
//   endfor
//endprogram