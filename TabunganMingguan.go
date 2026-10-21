package main
import "fmt"
func main(){
	var i, n, tabunganAwal, kenaikan, total, penambahan int
	fmt.Scan(&n, &tabunganAwal,&kenaikan)
	total = tabunganAwal
	penambahan = kenaikan
	for i = 1; i <= n; i++{
		total = total + penambahan
		penambahan = penambahan + kenaikan
	}
	fmt.Println(total)
}