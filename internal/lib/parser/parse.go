package parser

import "strings"

func ParsePassport(passport string) (string, string) {
    parsedPassport := strings.Split(passport, " ")
    
    return parsedPassport[0], parsedPassport[1]
}