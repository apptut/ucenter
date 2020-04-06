package user

import (
	"github.com/apptut/rsp"
	"github.com/gin-gonic/gin"
	"ucenter/app"
	"ucenter/app/models"
)

func Profile(ctx *gin.Context) {
	uid := ctx.GetString("uid")
	var user models.User
	err := app.Db().First(&user, uid).Error
	if err != nil {
		rsp.JsonErr(ctx, "get user info error")
		return
	}

	rsp.JsonOk(ctx, gin.H{
		"username":   user.Username,
		"mobile":     user.Mobile,
		"created_at": user.CreatedAt.Unix(),
	})
}
