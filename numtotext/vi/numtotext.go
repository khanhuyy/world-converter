package vi

import (
	"fmt"
	"strings"
	"unicode"
)

func ToWords(n int) string {

	exception := map[int]string{
		0:  "linh", // ngoai le hang chuc x
		1:  "mốt",  // ngoai le hang chuc x
		4:  "tư",   // ngoai le hang don vi x
		5:  "lăm",  // ngoai le hang don vi x
		10: "mười", // ngoai le hang chuc x
	}

	group := map[int]string{
		1: "mươi",
		2: "trăm",
		3: "nghìn",
		4: "vạn",
		6: "triệu",
		9: "tỉ",
	}

	num := map[string]string{
		"0": "không",
		"1": "một",
		"2": "hai",
		"3": "ba",
		"4": "bốn",
		"5": "năm",
		"6": "sáu",
		"7": "bảy",
		"8": "tám",
		"9": "chín",
	}
	s := fmt.Sprintf("%d", n)
	var lenString = len(s)
	check9 := lenString - (lenString % 9)
	check3 := lenString - (lenString % 3) // position multi tripple (kiem tra bo 3)
	offset := lenString - check3 + 1      // phan le
	oneFlag := false                      // flag to check exceptionOne
	if lenString < 2 {
		result := num[s]
		charResult := []byte(result)
		charResult[0] = byte(unicode.ToUpper(rune(charResult[0])))
		result = string(charResult)
		result = fmt.Sprintf(strings.TrimRight(result, " "))
		result += " đồng"
		return result
	} else {
		result := ""
		var flag int64 // kiem tra so khong
		for pos, char := range s {
			if offset == 1 {
				offset = 3
				check3 -= 3
			} else {
				offset -= 1
			}
			unit := len(s) - pos - 1
			triple := unit % 3
			switch triple {
			case 1:
				if fmt.Sprintf("%c", char) == "0" {
					flag += 1
				} else if fmt.Sprintf("%c", char) == "1" {
					if flag%3 == 1 {
						result += num["0"] + " " + group[2] + " " + exception[10] + " "
					} else {
						result += exception[10] + " "
					}
					flag = 0
				} else {
					if flag%3 == 1 {
						result += num["0"] + " " + group[2] + " " + num[fmt.Sprintf("%c", char)] + " " + group[1] + " "
					} else {
						result += num[fmt.Sprintf("%c", char)] + " " + group[1] + " "
					}
					oneFlag = true
					flag = 0
				}
			case 2:
				if fmt.Sprintf("%c", char) == "0" {
					flag += 1
				} else {
					flag = 0
					result += num[fmt.Sprintf("%c", char)] + " " + group[2] + " "
				}
			default:
				if fmt.Sprintf("%c", char) == "0" {
					flag += 1
				} else if fmt.Sprintf("%c", char) == "1" {
					if flag > 1 {
						result += exception[0] + " " + num["1"] + " "
						oneFlag = false
					} else {
						if oneFlag {
							result += exception[1] + " "
							oneFlag = true
						} else {
							result += num["1"] + " "
						}
					}
					flag = 0
				} else if fmt.Sprintf("%c", char) == "4" {
					if flag > 1 {
						result += exception[0] + " " + exception[4] + " "
					} else {
						if pos != 0 {
							result += exception[4] + " "
						} else {
							result += num["4"] + " "
						}
					}
					flag = 0
				} else if fmt.Sprintf("%c", char) == "5" {
					if flag > 1 {
						result += exception[0] + " " + num["5"] + " "
					} else {
						result += exception[5] + " "
					}
					flag = 0
				} else {
					if flag > 1 {
						result += exception[0] + " " + num[fmt.Sprintf("%c", char)] + " "
					} else {
						result += num[fmt.Sprintf("%c", char)] + " "
					}
					flag = 0
				}
				if check3%9 == 0 && check3 > 0 {
					check9 -= 9
					result += group[9] + " "
				} else {
					if flag < 1 {
						result += group[check3%9] + " "
						flag = 0
					} else if flag/3 < 1 {
						result += group[check3%9] + " "
						flag = 0
					}
				}
			}
		}
		charResult := []byte(result)
		charResult[0] = byte(unicode.ToUpper(rune(charResult[0])))
		result = string(charResult)
		result = fmt.Sprintf(strings.TrimRight(result, " "))
		result += " đồng"
		return result
	}
}
