package unit

import "github.com/gogf/gf/v2/errors/gerror"

func newErr(cmd byte, flag byte) error {
	return gerror.Newf("%x失败:%x", cmd, flag)
}
