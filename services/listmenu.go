package services

import (
	"bufio"
	"fgo-go-temp-convert/data"
	"fgo-go-temp-convert/utils"
	"fmt"
	"os"
	"strings"
)

func ListMenu(){
	fmt.Println(`
Data History Convert Temperature Celcius
`)
	for i, convert := range data.Temps {
		fmt.Printf("%d. Celcius ke %s : %v\n", i+1, convert.Name, convert.Result)
	}

	fmt.Print("\nPress enter to go back...")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	utils.Clear()
	return

}