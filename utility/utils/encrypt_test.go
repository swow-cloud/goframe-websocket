package utils_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"

	"goframe-websocket/utility/utils"
)

func Test_Encrypt(t *testing.T) {
	str, _ := utils.Encrypt([]byte(`abcdefg`), []byte(`kfskdfsf`))
	de, _ := utils.Decrypt(str, []byte("kfskdfsf"))
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.Trim(str), de)
	})
}
