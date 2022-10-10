package main

import "fmt"

func main (){
	var buah =[3]string{"apel","jeruk","anggur"}
	buah [1]="mangga"
	fmt.Println("jumlah Element :",len(buah))
	fmt.Println("isi Element :",buah)
}