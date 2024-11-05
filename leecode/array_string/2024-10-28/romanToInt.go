package main

import "fmt"

func main() {

	res := romanToInt("MCMXCIV")

	fmt.Println(res)
}

func romanToInt(s string) int {
	// 本质是一个加法
	// 从高位开始加
	// I II III IV V VI VII VII VIIII X
	// XI XII XII XIV XV XVI XVII XVIII XVIIII XX
	// XXI XII XIII ...
	// XLI XLII XLIII XLIV XLV ...
	// XCI XCII  ... XCIV XCV XCVI..
	// C CI CII ... CIV CV CVI
	// CC
	// CDI CDII CDIV CDV
	// ...
	// CMI CMII CMIII  CMIV CMV
	// MMI
	// M
	var res int = 0
	romanMap := make(map[string]int)
	romanMap["I"] = 1
	romanMap["V"] = 5
	romanMap["X"] = 10
	romanMap["L"] = 50
	romanMap["C"] = 100
	romanMap["D"] = 500
	romanMap["M"] = 1000
	continueFlag := false
	for k, v := range s {
		if continueFlag {
			continueFlag = false
			continue
		}
		str := string(v)
		var plusOneStr string
		switch str {
		case "I", "V":
			if k != len(s)-1 {
				plusOneStr = string(s[k+1])
			}
			if k != len(s)-1 && (plusOneStr == "V" || plusOneStr == "X") {
				continueFlag = true
				res += romanMap[plusOneStr] - romanMap[str]
			} else {
				res += romanMap[str]
			}
		case "X", "L":
			if k != len(s)-1 {
				plusOneStr = string(s[k+1])
			}
			if k != len(s)-1 && (plusOneStr == "L" || plusOneStr == "C") {
				continueFlag = true
				res += romanMap[plusOneStr] - romanMap[str]
			} else {
				res += romanMap[str]
			}
		case "C", "D":
			if k != len(s)-1 {
				plusOneStr = string(s[k+1])
			}
			if k != len(s)-1 && (plusOneStr == "D" || plusOneStr == "M") {
				continueFlag = true
				res += romanMap[plusOneStr] - romanMap[str]
			} else {
				res += romanMap[str]
			}
		case "M":
			res += romanMap[string(v)]
		}
	}
	return res
}

var romanIntValueMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToIntV2(s string) int {
	res := 0
	for i := 0; i < len(s)-1; i++ {
		cur, next := romanIntValueMap[s[i]], romanIntValueMap[s[i+1]]
		if cur < next {
			res -= cur
		} else {
			res += cur
		}
	}
	return res + romanIntValueMap[s[len(s)-1]]
}
