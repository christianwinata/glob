package tegj

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GenerateResponse(s string) string {
	var (
		total         float64
		creditssuffix string
	)
	num, material, respprefix, err := Translate(s)
	if err != nil {
		fmt.Println("I have no idea what you are talking about")
		os.Exit(0)
	}

	price, err := RomanToFloat(num)

	if material > 0 {
		total = price * material
		creditssuffix = " Credits"
	} else {
		total = price
	}

	return respprefix + strconv.FormatFloat(total, 'f', 1, 64) + creditssuffix
}

// returns roman number, material price, response prefix, and error
func Translate(s string) (string, float64, string, error) {
	fields := strings.Fields(s)

	var (
		roman      string
		material   float64
		respprefix string
	)
	for _, word := range fields {
		switch strings.ToLower(word) {
		case "glob":
			roman += "I"
			respprefix += word + " "
		case "prok":
			roman += "V"
			respprefix += word + " "
		case "pish":
			roman += "X"
			respprefix += word + " "
		case "tegj":
			roman += "L"
			respprefix += word + " "
		case "silver":
			material = 17 // glob glob Silver = 34c | 34 / 2 = 17
			respprefix += word + " "
		case "gold":
			material = 14450 // glob prok Gold = 57800 | 57800 / 4 = 14450
			respprefix += word + " "
		case "iron":
			material = 195.5 // pish pish Iron = 3910 | 3910 / 20 = 195.5
			respprefix += word + " "
		}
	}
	respprefix += "is "

	if len(roman) == 0 {
		return "", 0, respprefix, fmt.Errorf("no idea")
	}
	return roman, material, respprefix, nil
}

func RomanToFloat(s string) (float64, error) {
	l := len(s)
	var (
		num          float64
		roman        string
		rcount       int
		mappedRoman  float64
		mappedRoman2 float64
		err          error
	)
	for i := 0; i < l; i++ {
		roman = strings.ToUpper(string(s[i]))
		mappedRoman, err = MapRoman(roman)
		if err != nil {
			return 0, err
		}

		if i < l-1 {
			// these roman chars can only repeated 3 times.
			if s[i] == s[i+1] { // compare to next index
				if roman == "I" || roman == "X" || roman == "C" || roman == "M" {
					rcount++
					if rcount > 3 {
						return 0, fmt.Errorf("Invalid roman")
					}
				}
			} else {
				rcount = 0 // reset count
			}
			if mappedRoman2, err = MapRoman(string(s[i+1])); mappedRoman < mappedRoman2 {
				num += mappedRoman2 - mappedRoman
				i++
			} else {
				num += mappedRoman
			}
		} else {
			// these roman chars can only repeated 3 times.
			if roman == "I" || roman == "X" || roman == "C" || roman == "M" {
				rcount++
				if rcount > 3 {
					return 0, fmt.Errorf("Invalid roman")
				}
			} else {
				rcount = 0 // reset count
			}
			num += mappedRoman
		}

	}

	return num, nil
}

func MapRoman(s string) (float64, error) {
	switch strings.ToUpper(s) {
	default:
		return 0, fmt.Errorf("Invalid char on map roman")
	case "I":
		return 1, nil
	case "V":
		return 5, nil
	case "X":
		return 10, nil
	case "L":
		return 50, nil
	case "C":
		return 100, nil
	case "D":
		return 500, nil
	case "M":
		return 1000, nil
	}
}
