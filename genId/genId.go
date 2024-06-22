package genId

import (
	"crypto/md5"
	"encoding/hex"

	uuid "github.com/gofrs/uuid"
	"github.com/jaevor/go-nanoid"
)

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

// 32位guid
func GUID() string {
	u2, _ := uuid.NewV4()
	var s = hex.EncodeToString(u2.Bytes())
	return s
}

// 36位guidv f5394eef-e576-4709-9e4b-a7c231bd34a4
func UUID() string {
	u2, _ := uuid.NewV4()
	return u2.String()

}

// 雪花id 定制字符串
func Nanoid(len int) string {
	//str := "0123456789abcdefghijklmnopqrstuvwsyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str := "0123456789abcdefghijklmnopqrstuvwsyz"
	nid, _ := nanoid.CustomASCII(str, len)
	return nid()
}

// 雪花id 全字符
func NanoidFull(len int) string {
	str := "0123456789abcdefghijklmnopqrstuvwsyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//str := "0123456789abcdefghijklmnopqrstuvwsyz"
	nid, _ := nanoid.CustomASCII(str, len)
	return nid()
}
