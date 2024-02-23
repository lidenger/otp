package otp

import (
	"fmt"
	"testing"
)

func TestTOTP(t *testing.T) {
	key := "6KFGGMBBCCRAAY3D"
	code, err := TOTP(key)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("code:%s", code)
}

func TestVerifyTOTP(t *testing.T) {
	key := "6KFGGMBBCCRAAY3D"
	result, err := VerifyTOTP(key, "712512")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result:%v", result)
}

func TestWindows(t *testing.T) {
	tt1 := 1708686600
	tt2 := 1708686630
	tt3 := 1708686660
	fmt.Println(tt1 / 30)
	fmt.Println(tt2 / 30)
	fmt.Println(tt3 / 30)
}
