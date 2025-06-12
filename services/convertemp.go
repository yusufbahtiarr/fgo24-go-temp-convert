package services

import (
	"bufio"
	"fgo-go-temp-convert/calculate"
	"fgo-go-temp-convert/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ConvertTemp(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(`Convert Temperature Celcius
1. Kelvin
2. Reamur
3. Fahrenheit

0. Back to Menu

Pilih Menu : `)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Print("Nilai Suhu : ")
	inputValue, _ := reader.ReadString('\n')
	inputValue = strings.TrimSpace(inputValue)
	inputTemp, err := strconv.ParseFloat(inputValue, 64)
	if err != nil {
		fmt.Println("Error: Konversi gagal. Pastikan input adalah angka yang valid.")
		return
	}

	switch (input){
	case "1":
		fmt.Printf("Kelvin : %v", calculate.TempConvert("Kelvin", inputTemp))
	case "2":
		fmt.Printf("Reamur : %v", calculate.TempConvert("Reamur", inputTemp))
	case "3":
		fmt.Printf("Fahrenheit : %v", calculate.TempConvert("Fahrenheit", inputTemp))
	case "0":
		utils.Clear()
		return
	default:
		break
	}
}

func handleConvert(){

}