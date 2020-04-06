package middleware

import (
	"github.com/apptut/rsp"
	"github.com/gin-gonic/gin"
	"strings"
	"ucenter/app"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authVal := ctx.GetHeader("Authorization")
		uid, err := checkAuth(authVal)
		if err != nil {
			rsp.JsonErr(ctx, err)
			ctx.Abort()
			return
		}
		ctx.Set("uid", uid)

		ctx.Next()
	}
}

func checkAuth(val string) (string, *rsp.Error) {
	if len(val) <= 0 {
		return "", rsp.NewErrMsg("auth params now allowed")
	}

	arr := strings.Split(val, " ")
	if len(arr) != 2 {
		return "", rsp.NewErrMsg("auth params error")
	}

	token := strings.TrimSpace(arr[1])
	uid, err := app.Redis().Get("login:" + token).Result()
	if err != nil {
		return "", rsp.NewErrMsg("user auth failed")
	}

	return uid, nil
}
