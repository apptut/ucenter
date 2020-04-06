package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const MsgSuccess = "Success"
const MsgFailed = "Failed"

func JsonpOK(ctx *gin.Context, args ...interface{}) {
	data, msg := makeOkData(args...)
	ctx.JSONP(http.StatusOK, jsonBaseFormat(0, msg, data))
}

func JsonpErr(ctx *gin.Context, args ...interface{}) {
	data, msg, code := makeErrData(args...)
	ctx.JSONP(http.StatusOK, jsonBaseFormat(code, msg, data))
}

func makeOkData(args ...interface{}) (interface{}, string) {
	var data interface{}
	var msg = MsgSuccess

	if len(args) >= 1 {
		data = args[0]
	}

	if len(args) >= 2 {
		msg = args[1].(string)
	}

	return data, msg
}

func makeErrData(args ...interface{}) (interface{}, string, int) {
	var data interface{}
	msg := MsgFailed
	code := 1

	if len(args) >= 1 {
		msg = args[0].(string)
	}

	if len(args) >= 2 {
		data = args[1]
	}

	if len(args) >= 3 {
		code = args[2].(int)
	}

	return data, msg, code
}

// JsonOk(ctx, data, msg)
func JsonOk(ctx *gin.Context, args ...interface{}) {
	data, msg := makeOkData(args...)
	ctx.JSON(http.StatusOK, jsonBaseFormat(0, msg, data))
}

// JsonErr(ctx, msg, data)
func JsonErr(ctx *gin.Context, args ...interface{}) {
	data, msg, code := makeErrData(args...)
	ctx.JSON(http.StatusOK, jsonBaseFormat(code, msg, data))
}

func jsonBaseFormat(code int, msg string, data interface{}) interface{} {
	var baseData = make(map[string]interface{})
	baseData["errno"] = code
	baseData["msg"] = msg
	baseData["data"] = data
	return baseData
}