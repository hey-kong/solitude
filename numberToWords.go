package main

var lowNumber = map[int]string{
	0: "",
	1: "One",
	2: "Two",
	3: "Three",
	4: "Four",
	5: "Five",
	6: "Six",
	7: "Seven",
	8: "Eight",
	9: "Nine",
}

var midNumber = map[int]string{
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
}

var highNumber = map[int]string{
	2: "Twenty",
	3: "Thirty",
	4: "Forty",
	5: "Fifty",
	6: "Sixty",
	7: "Seventy",
	8: "Eighty",
	9: "Ninety",
}

// Leetcode 273. (hard)
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	billion := num / 1000000000
	million := (num - billion*1000000000) / 1000000
	thousand := (num - billion*1000000000 - million*1000000) / 1000
	rest := num - billion*1000000000 - million*1000000 - thousand*1000

	res := ""
	if billion != 0 {
		res += subNumberToWords(billion) + " Billion "
	}
	if million != 0 {
		res += subNumberToWords(million) + " Million "
	}
	if thousand != 0 {
		res += subNumberToWords(thousand) + " Thousand "
	}
	if rest != 0 {
		res += subNumberToWords(rest) + " "
	}
	return res[:len(res)-1]
}

func subNumberToWords(num int) string {
	res := ""
	if num >= 100 {
		res += lowNumber[num/100] + " Hundred "
		num -= num / 100 * 100
	}
	if num >= 20 {
		res += highNumber[num/10] + " "
		num -= num / 10 * 10
	}
	if num >= 10 {
		res += midNumber[num] + " "
	} else if num > 0 {
		res += lowNumber[num] + " "
	}
	return res[:len(res)-1]
}
