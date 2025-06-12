package main

import (
	"bufio"
	"fgo-go-temp-convert/services"
	"fgo-go-temp-convert/utils"
	"fmt"
	"os"
	"strings"
)

func main(){
	for {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Println(data.Temps)
	fmt.Print(`
Konversi Suhu dari Celcius:
1. List History Convert
2. Convert

0. Exit

Pilih Menu : `)


input, _ := reader.ReadString('\n')
input = strings.TrimSpace(input)

switch (input){
case "1":
	utils.Clear()
	services.ListMenu()
case "2":
	utils.Clear()
	services.ConvertTemp()
case "0":
	os.Exit(0)
default:
	utils.Clear()
}
}


}