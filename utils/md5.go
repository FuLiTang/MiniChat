package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Md5Encode 小写
func Md5Encode(d string) string {
	h := md5.New()
	h.Write([]byte(d))
	return hex.EncodeToString(h.Sum(nil))
}

// MD5Encode 大写
func MD5Encode(d string) string {
	return strings.ToUpper(Md5Encode(d))
}

// MakePassword 小写加密
func MakePassword(p, s string) string {
	return Md5Encode(p + s)
}

// ValidPassword 小写解密
func ValidPassword(p, s, password string) bool {
	return Md5Encode(p+s) == password
}
func Salt() string {
	return fmt.Sprintf("%06d", rand.Int31())
}
func SaltTime() string {
	return MD5Encode(fmt.Sprintf("%d", time.Now().Unix()))
}
