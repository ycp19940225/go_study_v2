package main

func main() {

}

func intToRoman(num int) string {
	// 取倍数 对 1000 100 10 取倍数和求于
	//
	var res = ""

	// 得到各个位置的数字
	num4 := int(num / 1000)
	num3 := int((num - num4*1000) / 100)
	num2 := int((num - num4*1000 - num3*100) / 10)
	num1 := num - num4*1000 - num3*100 - num2*10

	// 对各个位置进行`解码`
	if num4 > 0 {
		for i := 0; i < num4; i++ {
			res += "M"
		}
	}
	if num3 > 0 {
		if num3 < 4 {
			for i := 0; i < num3; i++ {
				res += "C"
			}
		} else if num3 == 4 {
			res += "CD"
		} else if num3 == 9 {
			res += "CM"
		} else {
			res += "D"
			for i := 0; i < num3-5; i++ {
				res += "C"
			}
		}
	}

	if num2 > 0 {
		if num2 < 4 {
			for i := 0; i < num2; i++ {
				res += "X"
			}
		} else if num2 == 4 {
			res += "XL"
		} else if num2 == 9 {
			res += "XC"
		} else {
			res += "L"
			for i := 0; i < num2-5; i++ {
				res += "X"
			}
		}
	}
	if num1 > 0 {
		if num1 < 4 {
			for i := 0; i < num1; i++ {
				res += "I"
			}
		} else if num1 == 4 {
			res += "IV"
		} else if num1 == 9 {
			res += "IX"
		} else {
			res += "V"
			for i := 0; i < num1-5; i++ {
				res += "I"
			}
		}
	}

	return res
}
