package vi

import "testing"

func TestNumberToWord(t *testing.T) {
	tests := []struct {
		input int
		//output  string
		expected string
	}{
		{
			input:    0,
			expected: "Không",
		},
		{
			input:    1,
			expected: "Một",
		},
		{
			input:    4,
			expected: "Bốn",
		},
		{
			input:    5,
			expected: "Năm",
		},
		{
			input:    10,
			expected: "Mười",
		},
		{
			input:    11,
			expected: "Mười một",
		},
		{
			input:    14,
			expected: "Mười bốn",
		},
		{
			input:    15,
			expected: "Mười lăm",
		},
		{
			input:    20,
			expected: "Hai mươi",
		},
		{
			input:    21,
			expected: "Hai mươi mốt",
		},
		{
			input:    24,
			expected: "Hai mươi tư",
		},
		{
			input:    25,
			expected: "Hai mươi lăm",
		},
		{
			input:    100,
			expected: "Một trăm",
		},
		{
			input:    101,
			expected: "Một trăm linh mốt",
		},
		{
			input:    104,
			expected: "Một trăm linh bốn",
		},
		{
			input:    105,
			expected: "Một trăm linh năm",
		},
		{
			input:    1000,
			expected: "Một nghìn",
		},
		{
			input:    1001,
			expected: "Một nghìn không trăm linh một",
		},
		{
			input:    1004,
			expected: "Một nghìn không trăm linh bốn",
		},
		{
			input:    1005,
			expected: "Một nghìn không trăm linh năm",
		},
		{
			input:    1010,
			expected: "Một nghìn không trăm mười",
		},
		{
			input:    1011,
			expected: "Một nghìn không trăm mười một",
		},
		{
			input:    1014,
			expected: "Một nghìn không trăm mười bốn",
		},
		{
			input:    1015,
			expected: "Một nghìn không trăm mười lăm",
		},
		{
			input:    1020,
			expected: "Một nghìn không trăm hai mươi",
		},
		{
			input:    1021,
			expected: "Một nghìn không trăm hai mươi mốt",
		},
		{
			input:    1024,
			expected: "Một nghìn không trăm hai mươi tư",
		},
		{
			input:    1025,
			expected: "Một nghìn không trăm hai mươi lăm",
		},
		{
			input:    2000000,
			expected: "Hai triệu",
		},
		{
			input:    2000001,
			expected: "Hai triệu không trăm linh một",
		},
		{
			input:    2000004,
			expected: "Hai triệu không trăm linh bốn",
		},
		{
			input:    2000005,
			expected: "Hai triệu không trăm linh năm",
		},
		{
			input:    2000010,
			expected: "Hai triệu không trăm mười",
		},
		{
			input:    2000011,
			expected: "Hai triệu không trăm mười một",
		},
		{
			input:    2000014,
			expected: "Hai triệu không trăm mười bốn",
		},
		{
			input:    2000015,
			expected: "Hai triệu không trăm mười lăm",
		},
		{
			input:    2000110,
			expected: "Hai triệu một trăm mười",
		},
		{
			input:    2000124,
			expected: "Hai triệu một trăm hai mươi tư",
		},
		{
			input:    2000125,
			expected: "Hai triệu một trăm hai mươi lăm",
		},
		{
			input:    2001111,
			expected: "Hai triệu không trăm linh một nghìn một trăm mười một",
		},
		{
			input:    2001125,
			expected: "Hai triệu không trăm linh một nghìn một trăm hai mươi lăm",
		},
		{
			input:    1000000000,
			expected: "Một tỉ",
		},
		{
			input:    1000000111,
			expected: "Một tỉ một trăm mười một",
		},
		{
			input:    1111111111,
			expected: "Một tỉ một trăm mười một triệu một trăm mười một nghìn một trăm mười một",
		},
		{
			input:    1234567891,
			expected: "Một tỉ hai trăm ba mươi tư triệu năm trăm sáu mươi bảy nghìn tám trăm chín mươi mốt",
		},
		{
			input:    1000234567891,
			expected: "Một nghìn tỉ hai trăm ba mươi tư triệu năm trăm sáu mươi bảy nghìn tám trăm chín mươi mốt",
		},
		{
			input:    1234567899876543210,
			expected: "Một tỉ hai trăm ba mươi tư triệu năm trăm sáu mươi bảy nghìn tám trăm chín mươi chín tỉ tám trăm bảy mươi sáu triệu năm trăm bốn mươi ba nghìn hai trăm mười",
		},
	}
	for _, test := range tests {
		output := ToWords(test.input)
		if output != test.expected {
			t.Errorf("ToWords(%d) = %s, want %s", test.input, output, test.expected)
		}
	}
}
