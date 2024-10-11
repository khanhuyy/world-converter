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
			expected: "Không đồng",
		},
		{
			input:    1,
			expected: "Một đồng",
		},
		{
			input:    25,
			expected: "Hai mươi lăm đồng",
		},
		{
			input:    1234567891,
			expected: "Một tỉ hai trăm ba mươi tư triệu năm trăm sáu mươi bảy nghìn tám trăm chín mươi mốt đồng",
		},
	}
	for _, test := range tests {
		output := ToWords(test.input)
		if output != test.expected {
			t.Errorf("ToWords(%d) = %s, want %s", test.input, output, test.expected)
		}
	}
}
