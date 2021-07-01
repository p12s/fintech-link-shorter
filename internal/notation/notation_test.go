package notation

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

// NotationDict - правила конвертации числа из 10-чной системы в 62-чную
// приведем к стандартному виду систем счисления - сначала цифры, потом символы
var NotationDict = map[int64]rune{
	0: '0', 1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7', 8: '8', 9: '9', 10: 'a',
	11: 'b', 12: 'c', 13: 'd', 14: 'e', 15: 'f', 16: 'g', 17: 'h', 18: 'i', 19: 'j', 20: 'k',
	21: 'l', 22: 'm', 23: 'n', 24: 'o', 25: 'p', 26: 'q', 27: 'r', 28: 's', 29: 't', 30: 'u',
	31: 'v', 32: 'w', 33: 'x', 34: 'y', 35: 'z', 36: 'A', 37: 'B', 38: 'C', 39: 'D', 40: 'E',
	41: 'F', 42: 'G', 43: 'H', 44: 'I', 45: 'J', 46: 'K', 47: 'L', 48: 'M', 49: 'N', 50: 'O',
	51: 'P', 52: 'Q', 53: 'R', 54: 'S', 55: 'T', 56: 'U', 57: 'V', 58: 'W', 59: 'X', 60: 'Y',
	61: 'Z',
}

// метод конвертации стандартный, тестируем скорее не его реализацию,
// а корректность его использования в методе "Convert"
func TestNotation_ConvertEveryNumberToRune(t *testing.T) {
	notation := NewConvert()
	for number, rune := range NotationDict {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, notation.Convert(number), string(rune))
		})
	}
}

// Выборочная проверка конвертации
func TestNotation_Convert(t *testing.T) {
	tests := []struct {
		name        string
		inputNumber int64
		expectedStr string
	}{
		{
			name:        "61=>Z",
			inputNumber: 61,
			expectedStr: "Z",
		},
		{
			name:        "62=>10",
			inputNumber: 62,
			expectedStr: "10",
		},
		{
			name:        "63=>11",
			inputNumber: 63,
			expectedStr: "11",
		},
		{
			name:        "64=>12",
			inputNumber: 64,
			expectedStr: "12",
		},
		{
			name:        "65=>13",
			inputNumber: 65,
			expectedStr: "13",
		},
		{
			name:        "85=>1n",
			inputNumber: 85,
			expectedStr: "1n",
		},
		{
			name:        "122=>1Y",
			inputNumber: 122,
			expectedStr: "1Y",
		},
		{
			name:        "123=>1Z",
			inputNumber: 123,
			expectedStr: "1Z",
		},
		{
			name:        "124=>20",
			inputNumber: 124,
			expectedStr: "20",
		},
		{
			name:        "125=>21",
			inputNumber: 125,
			expectedStr: "21",
		},
	}

	notation := NewConvert()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, notation.Convert(test.inputNumber), test.expectedStr)
		})
	}
}
