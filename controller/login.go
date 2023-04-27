package controller

import (
	"github.com/TskFok/AdminApi/model"
	"github.com/TskFok/AdminApi/tool"
	"github.com/TskFok/AdminApi/tool/tool-string"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	email := ctx.PostForm("email")

	if !tool_string.IsEmail(email) {
		ctx.JSON(http.StatusNotFound, "邮箱格式错误")
		return
	}
	password := ctx.PostForm("password")

	em := make(map[string]interface{})
	em["email"] = email
	mu := &model.Admin{}
	u := mu.Find(em)

	if u == nil {
		ctx.JSON(http.StatusNotFound, "邮箱不存在或密码错误")
		return
	}

	if u.Status != 1 {
		ctx.JSON(http.StatusNotFound, "用户状态错误")
		return
	}

	nPassword := tool_string.Password(password, u.Salt)
	if nPassword != u.Password {
		ctx.JSON(http.StatusNotFound, "邮箱不存在或密码错误")
		return
	}

	token, err := tool.JwtToken(u.Id, u.Email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "生成token错误")
		return
	}

	ctx.JSON(http.StatusOK, token)
}
