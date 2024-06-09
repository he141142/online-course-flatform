package utils

import (
	"log"
	"os"
	"strconv"
)

func EnvGetIntValue(varName string, fallBack int) int {
	valueStr := os.Getenv(varName)
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Println(err)
		return fallBack
	}
	return valueInt
}