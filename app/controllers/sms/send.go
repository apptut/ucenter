package sms

import (
	"github.com/gin-gonic/gin"
	"ucenter/app/svc/sms"
)
import "github.com/apptut/rsp"
import "github.com/apptut/validator"

func Send(ctx *gin.Context) {
	// 获取用户ip
	ip := ctx.ClientIP()

	// 传入参数验证
	mobile, err := validSendParams(ctx)
	if err != nil {
		rsp.JsonErr(ctx, err)
		return
	}

	// 获取频率验证
	if !sms.SendEnable(ip, mobile) {
		rsp.JsonErr(ctx, "发送频率过快")
		return
	}

	// 发送验证码
	err = sms.PostCode(mobile)
	if err != nil {
		rsp.JsonErr(ctx, err)
		return
	}

	// 更新频控
	sms.UpdateSmsCache(ip, mobile)

	rsp.JsonOk(ctx)
}

func validSendParams(ctx *gin.Context) (string, *rsp.Error) {
	// 获取 GET 请求参数
	mobile := ctx.Query("mobile")

	// 验证手机号是否合法
	_, err := validator.New(map[string][]string{
		"mobile": {mobile},
	}, map[string]string{
		"mobile": "mobile",
	}, map[string]string{
		"mobile": "手机号码格式不正确！",
	})

	if err != nil {
		return "", rsp.NewErr(err)
	}

	return mobile, nil
}
