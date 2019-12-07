package kit

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

// MD5Hash md5
func MD5Hash(b []byte) string {
	h := md5.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// MD5HashString md5
func MD5HashString(s string) string {
	return MD5Hash([]byte(s))
}

// SHA1Hash sha1
func SHA1Hash(b []byte) string {
	h := sha1.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA1HashString sha1
func SHA1HashString(s string) string {
	return SHA1Hash([]byte(s))
}
