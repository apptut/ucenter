package user

import (
	"github.com/apptut/rsp"
	"github.com/gin-gonic/gin"
	"ucenter/app/svc/login"
)

func Login(ctx *gin.Context) {
	// 参数验证
	err := login.ParamsValid(ctx)
	if err != nil {
		rsp.JsonErr(ctx, err)
		return
	}

	// 用户信息验证
	data, err := login.UserInfoValid(ctx)
	if err != nil {
		rsp.JsonErr(ctx, err)
		return
	}

	token, err := login.GenToken(data)
	if err != nil {
		rsp.JsonErr(ctx, err)
		return
	}

	rsp.JsonOk(ctx, gin.H{"t": token})
}
