package user

import (
	"github.com/apptut/rsp"
	"github.com/gin-gonic/gin"
	"ucenter/app/svc/reg"
)

func Reg(ctx *gin.Context) {
	// 参数验证
	err := reg.Valid(ctx)
	if err != nil {
		rsp.JsonErr(ctx, err)
		return
	}

	// 验证码验证
	err = reg.CheckCode(ctx)
	if err != nil {
		rsp.JsonErr(ctx, err)
		return
	}

	// 验证用户信息
	err = reg.CheckUser(ctx)
	if err != nil {
		rsp.JsonErr(ctx, err)
		return
	}

	// 保存数据
	err = reg.SaveData(ctx)
	if err != nil {
		rsp.JsonErr(ctx, err)
		return
	}

	rsp.JsonOk(ctx)
}
