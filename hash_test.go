package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MD5HashString(t *testing.T) {
	// MD5HashString
	ast := assert.New(t)
	rtn := MD5HashString("s")

	ast.Equal("03c7c0ace395d80182db07ae2c30f034", rtn, "s md5 result")
}

func Test_SHA1HashString(t *testing.T) {
	ast := assert.New(t)
	rtn := SHA1HashString("s")

	ast.Equal("a0f1490a20d0211c997b44bc357e1972deab8ae3", rtn, "s sha1 result")
}
