package lib

import (
	"errors"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func ParseStr[T constraints.Integer | constraints.Float](s string) (T, error) {
	z := zero[T]()
	var t any = z
	switch t.(type) {
	case int, int8, int16, int32, int64:
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return z, err
		}
		return T(num), nil
	case uint, uint8, uint16, uint32, uint64:
		unum, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return z, err
		}
		return T(unum), nil
	case float32, float64:
		fl, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return z, err
		}
		return T(fl), nil
	default:
		return z, errors.New("unreachable")
	}
}

func FormatInteger[T constraints.Integer](num, radix T) string {
	if num == 0 {
		return "0"
	}
	isMinus := num < 0
	if isMinus {
		num = -num
	}
	a := make([]int, 0, 100)
	for num > 0 {
		a = append(a, int(num%radix))
		num /= radix
	}
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
	sb := strings.Builder{}
	if isMinus {
		sb.WriteRune('-')
	}
	for i := 0; i < len(a); i++ {
		sb.WriteRune(rune(a[i]) + '0')
	}
	return sb.String()
}

func zero[T any]() T {
	var t T
	return t
}
