package reg

import (
	"github.com/apptut/rsp"
	"github.com/apptut/validator"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"ucenter/app"
	"ucenter/app/models"
)

func SaveData(ctx *gin.Context) *rsp.Error {
	password, err := bcrypt.GenerateFromPassword([]byte(ctx.PostForm("password")), 10)
	model := models.User{
		Username: ctx.PostForm("username"),
		Mobile:   ctx.PostForm("mobile"),
		Password: string(password),
	}
	err = app.Db().Create(&model).Error
	if err != nil {
		return rsp.NewErr(err)
	}

	return nil
}

func CheckUser(ctx *gin.Context) *rsp.Error {
	username := ctx.PostForm("username")
	mobile := ctx.PostForm("mobile")

	var model models.User
	err := app.Db().Find(&model, "username = ? or mobile = ?", username, mobile).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
	}

	return rsp.NewErrMsg("user already exist")
}

func CheckCode(ctx *gin.Context) *rsp.Error {
	mobile := ctx.PostForm("mobile")
	code := ctx.PostForm("code")

	cacheCode, err := app.Redis().Get("reg:" + mobile).Result()
	if err != nil {
		return rsp.NewErr(err)
	}

	if code != cacheCode {
		return rsp.NewErrMsg("验证码不正确！")
	}

	return nil
}

func Valid(ctx *gin.Context) *rsp.Error {
	_, err := validator.New(map[string][]string{
		"username": {ctx.PostForm("username")},
		"mobile":   {ctx.PostForm("mobile")},
		"password": {ctx.PostForm("password")},
		"code":     {ctx.PostForm("code")},
	}, map[string]string{
		"username": "regex:^[\\w_]{6,20}$",
		"mobile":   "mobile",
		"password": "regex:^[\\S]{6,20}$",
		"code":     "regex:^[0-9]{4}$",
	})

	if err != nil {
		return rsp.NewErr(err)
	}

	return nil
}
