package main

import (
	"diplom/internal/sms"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	rawData, err := os.ReadFile("../sms.data")
	if err != nil {
		log.Fatal("Что-то опять сломалось", err)
	}
	strData := strings.Split(string(rawData), "\n")
	for i := range strData {
		data := sms.FromSTR(strData[i])
		fmt.Println(data)
	}

}
