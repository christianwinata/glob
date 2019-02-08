package tests

import (
	"fmt"
	"glob/tegj"
	"testing"
)

func TestTranslate(t *testing.T) {
	s := "test glob silver"
	r, m, resp, err := tegj.Translate(s)
	if r != "I" || m != 17 || resp != "glob silver is " || err != nil {
		t.Fail()
	}

	s = "invalid test"
	_, _, _, err = tegj.Translate(s)
	if err == nil {
		t.Fail()
	}
}

func TestRomanToFloat(t *testing.T) {
	s := "MCMIII"
	f, err := tegj.RomanToFloat(s)
	if f != 1903 || err != nil {
		t.Fail()
	}

	// invalid
	s = "MCMIIII"
	_, err = tegj.RomanToFloat(s)
	if err == nil {
		t.Fail()
	}
}

func TestFullResponse(t *testing.T) {
	s := "how much is pish tegj glob glob"
	resp := tegj.GenerateResponse(s)
	if resp != "pish tegj glob glob is 42.0" {
		t.Fail()
	}
	fmt.Println(s)
	fmt.Println(resp)

	s = "how many Credits is glob prok Silver"
	resp = tegj.GenerateResponse(s)
	if resp != "glob prok Silver is 68.0 Credits" {
		t.Fail()
	}
	fmt.Println(s)
	fmt.Println(resp)

	s = "how many Credits is glob prok Gold"
	resp = tegj.GenerateResponse(s)
	if resp != "glob prok Gold is 57800.0 Credits" {
		t.Fail()
	}
	fmt.Println(s)
	fmt.Println(resp)

	s = "how many Credits is glob prok Iron"
	resp = tegj.GenerateResponse(s)
	if resp != "glob prok Iron is 782.0 Credits" {
		t.Fail()
	}
	fmt.Println(s)
	fmt.Println(resp)

	s = "how much wood could a woodchuck chuck if a woodchuck could chuck wood"
	resp = tegj.GenerateResponse(s)
	if resp != "I have no idea what you are talking about" {
		t.Fail()
	}
	fmt.Println(s)
	fmt.Println(resp)
}
