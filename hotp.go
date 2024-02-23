package otp

import (
	"crypto/hmac"
	"crypto/sha1"
)

// HOTP HMAC-Based One-Time Password
// 标准：https://datatracker.ietf.org/doc/html/rfc4226
// 默认算法为HMAC SHA1，默认6位
// HOTP(K,C) = Truncate(HMAC-SHA-1(K,C))
// k 密钥种子(shared secret between client and server)，一个账号一个，全局唯一
// c 64位的增量因子(8-byte counter value, the moving factor)，可以为时间戳，也可以传入TOTP计算的时间
// 返回
func HOTP(k, c []byte) int32 {
	// 1. 根据k和c生成hmac sha1摘要
	digested := generateDigest(k, c)
	// 2. 根据offset截取摘要的4个字节
	trunc := truncation(digested)
	// 3. 和10^6求余得到6位code码
	code := trunc % 1000000
	return code
}

// Step 1: Generate an HMAC-SHA-1 value Let HS = HMAC-SHA-1(K,C)  // HS is a 20-byte string
func generateDigest(key, data []byte) [20]byte {
	h := hmac.New(sha1.New, key)
	h.Write(data)
	digested := h.Sum(nil)
	return [20]byte(digested[:])
}

// Step 2: Generate a 4-byte string (Dynamic Truncation)
func truncation(digested [20]byte) int32 {
	offset := digested[19] & 0xf
	trunc := int32(digested[offset]) & 0x7f << 24
	trunc |= int32(digested[offset+1]) & 0xff << 16
	trunc |= int32(digested[offset+2]) & 0xff << 8
	trunc |= int32(digested[offset+3]) & 0xff
	return trunc
}
