package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GenMD5(raw string) string {
	sum := md5.Sum([]byte(raw))
	return hex.EncodeToString(sum[:])
}

func CheckMD5(raw string, target string) bool {
	return GenMD5(raw) == target
}