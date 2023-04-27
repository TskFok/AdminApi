package controller

import (
	"github.com/TskFok/AdminApi/model"
	tool_string "github.com/TskFok/AdminApi/tool/tool-string"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserList(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	size := ctx.DefaultQuery("size", "10")

	p, err := strconv.Atoi(page)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "页码错误")

		return
	}

	s, err := strconv.Atoi(size)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "页数错误")

		return
	}

	u := &model.User{}

	ctx.JSON(http.StatusOK, u.List(p, s))
}

func AddUser(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	if !tool_string.IsEmail(email) {
		ctx.JSON(http.StatusBadRequest, "邮箱格式错误")
		return
	}
	err := tool_string.CheckPasswordLever(password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	salt := tool_string.UUID()
	nPassword := tool_string.Password(password, salt)

	user := &model.User{}
	user.Email = email
	user.Salt = salt
	user.Password = nPassword
	user.Status = 1

	uid := user.Add(user)

	if uid == 0 {
		ctx.JSON(http.StatusBadRequest, "添加用户失败")
		return
	}

	ctx.JSON(http.StatusOK, "添加成功")
}

func UpdateStatus(ctx *gin.Context) {
	id := ctx.PostForm("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, "id不存在")
		return
	}

	i, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "获取id异常")
		return
	}

	user := &model.User{}
	user.Id = uint32(i)
	user.Find(user)

	condition := make(map[string]interface{})
	condition["status"] = 1 - user.Status

	success := user.Update(user, condition)

	if !success {
		ctx.JSON(http.StatusBadRequest, "更新失败")
		return
	}

	ctx.JSON(http.StatusOK, "更新成功")
}
