package main

import "fmt"

func main() {
	var nilai = 10
	if nilai == 10 {
		fmt.Println("lulus dengan sempuran")
	} else if nilai >= 7 {
		fmt.Println("lulus")

	} else if nilai == 6 {
		fmt.Println("sedikit lagi")
	} else {
		fmt.Println("belajar lagi")
	}

}
