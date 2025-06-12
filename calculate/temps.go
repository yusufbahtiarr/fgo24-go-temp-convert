package calculate

import "strings"

func TempConvert(tempName string, tempValue float64 ) float64 {
	tempName = strings.ToLower(tempName)
	switch (tempName){
	case "kelvin":
		return tempValue + 273
	case "reamur":
		return tempValue * 4 / 5
	case "fahrenheit":
		return tempValue * 9 / 5 + 32
	default:
		return 0
	}
}