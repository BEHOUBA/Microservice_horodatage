package main

import (
	"fmt"
	"strconv"
	"strings"
)

func monthFormatter(month string) string {
	if len(month) < 3 {
		return ""
	}
	firstThreeLetter := month[:3]
	letterSlice := strings.Split(firstThreeLetter, "")
	letterSlice[0] = strings.ToUpper(letterSlice[0])
	result := strings.Join(letterSlice, "")
	return result
}

func humanToUnix(value string) {
	var dateFormated []string

	dateSlice := strings.Split(strings.Replace(value, ",", "", -1), " ")

	for _, str := range dateSlice {
		checkValue, err := strconv.Atoi(str)
		if err != nil {
			dateFormated = append(dateFormated, monthFormatter(str))
		} else {
			if checkValue <= 31 {
				dateFormated = append(dateFormated, str)
			} else {
				dateFormated = append(dateFormated, str)
			}
		}
	}
	fmt.Println(strings.Join(dateFormated, "-"))
}

func main() {
	humanToUnix("Decembre 14, 2015")
}
