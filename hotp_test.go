package otp

import (
	"testing"
)

func TestGenerateDigest(t *testing.T) {
	key := []byte("123abcdddk323")
	data := []byte("123abc")
	digest := generateDigest(key, data)
	t.Logf("%v", digest)
}

func TestTruncation(t *testing.T) {
	digested := [20]byte{0x1f, 0x86, 0x98, 0x69, 0x0e, 0x02, 0xca, 0x16, 0x61, 0x85, 0x50, 0xef, 0x7f, 0x19, 0xda, 0x8e, 0x94, 0x5b, 0x55, 0x5a}
	trunc := truncation(digested)
	code := trunc % 1000000
	if code != 872921 {
		t.Fatalf("code应该为872921")
	}
}
