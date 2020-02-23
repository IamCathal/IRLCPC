package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func generateAddress(name string, number int) string {
	// fmt.Println(name, number)
	matchString := regexp.MustCompile(`[^A-Za-z]`)
	nameArr := matchString.Split(name, 6)
	resEmail := ""

	for i := 0; i < len(nameArr); i++ {
		if i > 0 {
			resEmail += "."
			resEmail += nameArr[i]
		} else {
			resEmail += nameArr[i]
		}
	}
	if number > 1 {
		resEmail += strconv.Itoa(number)
	}
	resEmail += "@ucc.ie"
	return resEmail
}

func question10() {
	numNames := 6
	names := make([]string, numNames)
	emails := make([]string, numNames)

	names[0] = "Bastien Pietropaoli"
	names[1] = "Milan De Cauwer"
	names[2] = "Raffaele Baldassini"
	names[3] = "Eduardo Vyhmeister"
	names[4] = "Iarfhlaith O’Sullivan"
	names[5] = "Milan De Cauwer"

	var seenNames = make(map[string]int)

	for i := 0; i < len(names); i++ {
		seenNames[strings.ToLower(names[i])]++
		emails[i] = generateAddress(strings.ToLower(names[i]), seenNames[strings.ToLower(names[i])])
	}

	fmt.Println(emails)
}
