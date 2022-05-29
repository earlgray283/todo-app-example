package lib

import (
	"strconv"
	"testing"
)

func FuzzParseStr(f *testing.F) {
	f.Add("1")
	f.Fuzz(func(t *testing.T, a string) {
		res, err1 := ParseStr[int64](a)
		t.Log("err1", err1)
		expect, err2 := strconv.ParseInt(a, 10, 64)
		t.Log("err2", err1)
		if (err1 != nil && err2 == nil) || (err1 == nil && err2 != nil) {
			t.Fatal(err1, err2)
		}
		if err1 != nil && err2 != nil {
			return
		}
		if res != expect {
			t.Fatal(a, res, expect)
		}
	})
}

func FuzzFormatInteger(f *testing.F) {
	f.Add(1)
	f.Fuzz(func(t *testing.T, a int) {
		res := FormatInteger(a, 10)
		expect := strconv.FormatInt(int64(a), 10)
		if res != expect {
			t.Fatalf("%v|%v|%v", a, res, expect)
		}
	})
}
