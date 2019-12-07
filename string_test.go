package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringsToInterfaces(t *testing.T) {
	ast := assert.New(t)

	slice := []string{"foo", "bar"}
	rtn := StringsToInterfaces(slice)

	ast.Equal([]interface{}([]interface{}{"foo", "bar"}), rtn, "slice StringsToInterfaces result")
}
