package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SliceToMap(t *testing.T) {
	ast := assert.New(t)

	// SliceToMap
	slice := []string{"foo", "bar"}
	rtn := SliceToMap(slice)

	ast.Equal(map[string]interface{}(map[string]interface{}{"bar": struct{}{}, "foo": struct{}{}}), rtn, "slice StringsToInterfaces result")

}
