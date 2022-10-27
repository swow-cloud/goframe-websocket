package utils

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
)

var Charset = new(charset)

type charset struct{}

func (util *charset) GetStack(err error) []string {
	stackList := gstr.Split(gerror.Stack(err), "\n")
	for i := 0; i < len(stackList); i++ {
		stackList[i] = gstr.Replace(stackList[i], "\t", "--> ")
	}

	return stackList
}
