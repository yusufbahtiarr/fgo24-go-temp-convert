package services

import (
	"bufio"
	"fgo-go-temp-convert/data"
	"fgo-go-temp-convert/utils"
	"fmt"
	"os"
)

func ListMenu(){
	fmt.Println(`
Data History Convert Temperature Celcius
`)
	for i, convert := range data.Temps {
		fmt.Printf("%d. Celcius ke %s : %v\n", i+1, convert.Name, convert.Result)
	}
	if len(data.Temps) < 1 {
		fmt.Println("Data Masih Kosong")
	}

	fmt.Print("\nPress enter to go back...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	utils.Clear()
	return

}