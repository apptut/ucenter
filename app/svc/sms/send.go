package sms

import (
	"github.com/apptut/rsp"
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
	"gopkg.in/redis.v5"
	"math/rand"
	"strconv"
	"time"
	"ucenter/app"
	"ucenter/app/utils/crypto"
)

func UpdateSmsCache(ip string, phone string) {
	redisKey := "sms:" + crypto.MD5(ip+phone)
	app.Redis().Set(redisKey, true, time.Minute*1)
}

// 验证发送频率
func SendEnable(ip string, phone string) bool {
	redisKey := "sms:" + crypto.MD5(ip+phone)
	redisErr := app.Redis().Get(redisKey).Err()
	if redisErr != nil {
		return redisErr == redis.Nil
	}

	return false
}

// 发送验证码
func PostCode(mobile string) *rsp.Error {
	code := generateCode()

	// 请填写自己申请的 apiKey
	client := ypclnt.New("apikey")
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = mobile
	param[ypclnt.TEXT] = "【猿谋人】您的验证码是" + code

	r := client.Sms().SingleSend(param)
	if r.Code != ypclnt.SUCC {
		return rsp.NewErrMsg(r.Msg)
	}

	// 保存数据供注册接口验证使用
	app.Redis().Set("reg:"+mobile, code, time.Minute*30)

	return nil
}

func generateCode() string {
	code := ""
	// 以当前时间作为伪随机种子
	rand.Seed(time.Now().Unix())

	// 生成一个 [0,10) 范围的数字，循环四次
	for i := 0; i < 4; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}

	return code
}
