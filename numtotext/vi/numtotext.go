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

	num := map[rune]string{
		'0': "không",
		'1': "một",
		'2': "hai",
		'3': "ba",
		'4': "bốn",
		'5': "năm",
		'6': "sáu",
		'7': "bảy",
		'8': "tám",
		'9': "chín",
	}
	s := fmt.Sprintf("%d", n)
	var result string
	var lenString = len(s) // flag to check exceptionOne
	if lenString < 2 {
		result = num[rune(s[0])]
	} else { // len s = 5
		check9 := lenString - (lenString % 9) // 5
		check3 := lenString - (lenString % 3) // 2
		offset := lenString - check3 + 1      // 4
		oneFlag := false

		var zeroFlag int64 // continuous zero
		for pos, char := range s {
			if offset == 1 {
				offset = 3
				check3 -= 3
			} else {
				offset -= 1
			}
			unit := len(s) - pos - 1
			set3 := unit % 3
			switch set3 {
			case 1:
				if char == '0' {
					zeroFlag += 1
				} else if char == '1' {
					if zeroFlag%3 == 1 {
						result += num['0'] + " " + group[2] + " " + exception[10] + " "
					} else {
						result += exception[10] + " "
					}
					zeroFlag = 0
				} else {
					if zeroFlag%3 == 1 {
						result += num['0'] + " " + group[2] + " " + num[char] + " " + group[1] + " "
					} else {
						result += num[char] + " " + group[1] + " "
					}
					oneFlag = true
					zeroFlag = 0
				}
			case 2:
				if char == '0' {
					zeroFlag += 1
				} else {
					zeroFlag = 0
					result += num[char] + " " + group[2] + " "
				}
			default:
				if char == '0' {
					zeroFlag += 1
				} else if char == '1' {
					if zeroFlag > 1 {
						result += exception[0] + " " + num['1'] + " "
						oneFlag = false
					} else {
						if oneFlag {
							result += exception[1] + " "
							oneFlag = true
						} else {
							result += num['1'] + " "
						}
					}
					zeroFlag = 0
				} else if char == '4' {
					if zeroFlag > 1 {
						result += exception[0] + " " + exception[4] + " "
					} else {
						if pos == 0 {
							println(pos)
							result += exception[4] + " "
						} else {
							result += num['4'] + " "
						}
					}
					zeroFlag = 0
				} else if char == '5' {
					if zeroFlag > 1 {
						result += exception[0] + " " + num['5'] + " "
					} else {
						result += exception[5] + " "
					}
					zeroFlag = 0
				} else {
					if zeroFlag > 1 {
						result += exception[0] + " " + num[char] + " "
					} else {
						result += num[char] + " "
					}
					zeroFlag = 0
				}
				if check3%9 == 0 && check3 > 0 {
					check9 -= 9
					result += group[9] + " "
				} else {
					if zeroFlag < 1 {
						result += group[check3%9] + " "
						zeroFlag = 0
					} else if zeroFlag/3 < 1 {
						result += group[check3%9] + " "
						zeroFlag = 0
					}
				}
			}
		}
	}
	charResult := []byte(result)
	charResult[0] = byte(unicode.ToUpper(rune(charResult[0])))
	result = string(charResult)
	result = fmt.Sprintf(strings.TrimRight(result, " "))
	return result
}
