package main

import "fmt"

// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Print("Angka",
// 			i)
// 	}
// }

// func main(){
// var i=0
// for i<10{
// 	fmt.Println("Angka",i)
// 	i++
// }
// }

func main() {
	var i = 0
	for {
		fmt.Println("Angka", i)
		i++
		if i == 10 {
			break
		}
	}
}
