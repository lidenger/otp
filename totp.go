// Package otp_algorithm otp_algorithm
package otp

import (
	"encoding/base32"
	"strconv"
	"time"
)

var (
	T0      int64 = 0                 // 时间起始时间戳，0表示UNIX起始时间即：1970-01-01 00:00:00
	step    int64 = 30                // 步骤：每个窗口30秒
	windows       = [3]int8{0, -1, 1} // 当前窗口，前一个窗口，后一个窗口
	digit         = 6                 // 验证码位数
)

// TOTPWithOptions Time-Based One-Time Password
// 标准：https://datatracker.ietf.org/doc/html/rfc6238
// TOTP = HOTP(K, T)
// k 账号密钥种子
// t 时间增量，一般为当前时间
// window 窗口参数 -1:上一个 ; 0:当前; 1:下一个
// 返回6位
func TOTPWithOptions(k []byte, t int64, window int8) string {
	T := (t - T0) / step
	T += int64(window)
	data := [8]byte{}
	for i := 7; i > 0; i-- {
		data[i] = byte(T)
		T = T >> 8
	}
	code := HOTP(k, data[:])
	codeStr := strconv.FormatInt(int64(code), 10)
	for len(codeStr) != digit {
		codeStr = "0" + codeStr
	}
	return codeStr
}

// TOTP
// k : base32 encode string
// 默认当前窗口
// 返回6位验证码
func TOTP(k string) (string, error) {
	now := time.Now().Unix()
	keyDecode, err := base32.StdEncoding.DecodeString(k)
	if err != nil {
		return "", err
	}
	code := TOTPWithOptions(keyDecode, now, 0)
	return code, nil
}

// VerifyTOTP 验证otp
// 在3个窗口内都有效
// k 密钥种子，base32 encode
// code 需要校验的otp
func VerifyTOTP(k string, code string) (bool, error) {
	for _, window := range windows {
		now := time.Now().Unix()
		keyDecode, err := base32.StdEncoding.DecodeString(k)
		if err != nil {
			return false, err
		}
		c := TOTPWithOptions(keyDecode, now, window)
		if c == code {
			return true, nil
		}
	}
	return false, nil
}
