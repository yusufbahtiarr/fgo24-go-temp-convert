package manage

import "fgo-go-temp-convert/models"

func TempConvertAdd(temps *models.Temps, temp models.Temp){
	*temps = append(*temps, temp)
}