package services

import (
	"bufio"
	"fgo-go-temp-convert/calculate"
	"fgo-go-temp-convert/data"
	"fgo-go-temp-convert/manage"
	"fgo-go-temp-convert/models"
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
	inputTemp, err := strconv.ParseFloat(inputValue, 32)
	if err != nil {
		fmt.Println("Error: Konversi gagal. Pastikan input adalah angka yang valid.")
		return
	}

	switch (input){
	case "1":
		temp := "Kelvin"
		result := calculate.TempConvert(temp, inputTemp)
		convertTemp := models.Temp{Name: temp, Result: float32(result)}
		fmt.Printf("Hasil Konversi Suhu dari Celcius ke %s : %v\n",temp, result)
		manage.TempConvertAdd(&data.Temps, convertTemp)
		fmt.Print("\nPress enter to go back...")
		reader.ReadString('\n')
		utils.Clear()
		return
		
	case "2":
		temp := "Reamur"
		result := calculate.TempConvert(temp, inputTemp)
		convertTemp := models.Temp{Name: temp, Result: float32(result)}
		fmt.Printf("Hasil Konversi Suhu dari Celcius ke %s : %v\n",temp, result)
		manage.TempConvertAdd(&data.Temps, convertTemp)
		fmt.Print("\nPress enter to go back...")
		reader.ReadString('\n')
		utils.Clear()
		return
	case "3":
		temp := "Fahrenheit"
		result := calculate.TempConvert(temp, inputTemp)
		convertTemp := models.Temp{Name: temp, Result: float32(result)}
		fmt.Printf("Hasil Konversi Suhu dari Celcius ke %s : %v\n",temp, result)
		manage.TempConvertAdd(&data.Temps, convertTemp)
		fmt.Print("\nPress enter to go back...")
		reader.ReadString('\n')
		utils.Clear()
		return
	case "0":
		utils.Clear()
		return
	default:
		break
	}
}

// func handleConvert(temp string, result string){
// 	fmt.Printf("Hasil Konversi Suhu dari Celcius ke %s : %v\n", temp, result)
// 		manage.TempConvertAdd(&data.Temps, models.Temp{
// 			Name: temp,
// 			Result: float32(result),
// 		})
// 		fmt.Print("\nPress enter to go back...")
// 		reader.ReadString('\n')
// 		utils.Clear()
// }