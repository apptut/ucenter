package login

import (
	"github.com/apptut/rsp"
	"github.com/apptut/validator"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/redis.v5"
	"math/rand"
	"strconv"
	"time"
	"ucenter/app"
	"ucenter/app/models"
	"ucenter/app/utils/crypto"
)

func GenToken(model *models.User) (string, *rsp.Error) {
	now := strconv.FormatInt(time.Now().UnixNano(), 10)
	rand.Seed(time.Now().UnixNano())
	randStr := strconv.Itoa(rand.Intn(100000))

	uid := strconv.Itoa(int(model.ID))
	token := crypto.SHA256(uid + randStr + model.Mobile + now)

	err := app.Redis().Set("login:"+token, uid, time.Hour*4).Err()
	lErr := app.Redis().Set("login:uid:" + uid, uid, 0)
	if err != nil && lErr != nil {
		return "", rsp.NewErr(err)
	}

	return token, nil
}

func UserInfoValid(ctx *gin.Context) (*models.User, *rsp.Error) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	var model models.User
	err := app.Db().First(&model, "username = ? OR mobile = ?", username, username).Error
	if err != nil {
		return nil, rsp.NewErr(err)
	}

	if bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(password)) != nil {
		return nil, rsp.NewErrMsg("用户名或密码不正确！")
	}

	// 判断用户是否已经登录
	uid := strconv.Itoa(int(model.ID))
	exists, err := app.Redis().Get("login:uid:" + uid).Result()
	if err != nil && err != redis.Nil {
		return nil, rsp.NewErr(err)
	}

	if exists != ""{
		return nil, rsp.NewErrMsg("用户已登录")
	}

	return &model, nil
}

func ParamsValid(ctx *gin.Context) *rsp.Error {
	_, err := validator.New(map[string][]string{
		"username": {ctx.PostForm("username")},
		"password": {ctx.PostForm("password")},
	}, map[string]string{
		"username": "regex:^\\w_{6,20}$",
		"password": "regex:^[\\S]{6,20}$",
	})

	if err != nil {
		return rsp.NewErr(err)
	}

	return nil
}
