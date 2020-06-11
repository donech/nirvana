package v1

import (
	"github.com/donech/nirvana/internal/code"

	"github.com/unknwon/com"

	"github.com/donech/tool/xlog"

	"github.com/donech/nirvana/internal/domain/user/repository"
	"github.com/gin-gonic/gin"
)

func NewUserController(userSimpleService *repository.UserRepository) *UserController {
	return &UserController{UserSimpleService: userSimpleService}
}

type UserController struct {
	UserSimpleService *repository.UserRepository
}

// @获取指定ID用户
// @Summary Get a single user
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /v1/user/{id} [get]
func (c UserController) GetUser(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt64()
	user, err := c.UserSimpleService.ItemByID(ctx.Request.Context(), id)
	if err != nil {
		xlog.S(ctx.Request.Context()).Error("some thing wrong", err)
		ResponseJSON(ctx, code.Error, err.Error(), nil)
		return
	}
	ResponseJSON(ctx, code.Success, "", user)
}

// @获取用户列表
// @Summary Get user list
// @Produce  json
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /v1/user [get]
func (c UserController) GetUserList(ctx *gin.Context) {
	cursor := com.StrTo(ctx.Query("cursor")).MustInt64()
	size := com.StrTo(ctx.Query("size")).MustInt64()
	if size == 0 {
		size = 20
	}
	users, err := c.UserSimpleService.ItemsByCursorReverse(ctx.Request.Context(), cursor, size+1)
	if err != nil {
		xlog.S(ctx.Request.Context()).Error("some thing wrong", err)
		ResponseJSON(ctx, code.Error, err.Error(), nil)
		return
	}
	cursor = 0
	hasMore := false
	if int64(len(users)) > size {
		cursor = users[size].ID
		hasMore = true
		users = users[0:size]
	}
	ResponseJSON(ctx, code.Success, "", gin.H{
		"list": users,
		"pager": Pager{
			Cursor:  cursor,
			Size:    size,
			HasMore: hasMore,
		}})
}

type UserForm struct {
	Name  string `json:"name" form:"name" binding:"required"`
	Phone string `json:"phone" form:"phone" binding:"required,email"`
}

// @创建用户
// @Summary create a user
// @Param name body string true "name"
// @Param phone body string true "phone"
// @Success 200 {object} Response
// @Failure 200 {object} Response
// @Router /v1/user [post]
func (c UserController) CreateUser(ctx *gin.Context) {
	var userForm UserForm
	err := ctx.ShouldBind(&userForm)
	if err != nil {
		ResponseJSON(ctx, code.Error, err.Error(), nil)
		return
	}
	data := map[string]interface{}{
		"name":  userForm.Name,
		"phone": userForm.Phone,
	}
	user, err := c.UserSimpleService.Create(ctx.Request.Context(), data)
	if err != nil {
		xlog.S(ctx.Request.Context()).Error("some thing wrong", err)
		ResponseJSON(ctx, code.Error, err.Error(), nil)
		return
	}
	ResponseJSON(ctx, code.Success, "ok", gin.H{"status": "success", "user": user})
}

// @更新用户
// @Summary update a user
// @Param name body string true "name"
// @Param phone body string true "phone"
// @Success 200 {object} Response
// @Failure 200 {object} Response
// @Router /v1/user [put]
func (c UserController) UpdateUser(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt64()
	var userForm UserForm
	err := ctx.ShouldBind(&userForm)
	if err != nil {
		ResponseJSON(ctx, code.Error, err.Error(), nil)
		return
	}
	data := map[string]interface{}{
		"name":  userForm.Name,
		"phone": userForm.Phone,
	}
	err = c.UserSimpleService.Update(ctx.Request.Context(), id, data)
	if err != nil {
		xlog.S(ctx.Request.Context()).Error("some thing wrong", err)
		ResponseJSON(ctx, code.Error, err.Error(), nil)
		return
	}
	ctx.JSON(200, gin.H{"status": "success"})
}

// @更新用户
// @Summary delete a user
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /v1/user/{id} [delete]
func (c UserController) DeleteUser(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt64()
	err := c.UserSimpleService.Delete(ctx.Request.Context(), id)
	if err != nil {
		xlog.S(ctx.Request.Context()).Error("some thing wrong", err)
		ResponseJSON(ctx, code.Error, err.Error(), nil)
		return
	}
	ResponseJSON(ctx, code.Success, "", gin.H{"status": "success"})
}

// @创建用户表
// @Summary init user table
// @Produce  json
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /tool/migration/user [get]
func (c UserController) Migration(ctx *gin.Context) {
	c.UserSimpleService.Migration()
	ResponseJSON(ctx, code.Success, "", gin.H{"status": "success"})
}

func (c UserController) RegisterRoute(root *gin.RouterGroup) {
	r := root.Group("/v1/user")
	r.GET("/", c.GetUserList)
	r.GET("/:id", c.GetUser)
	r.DELETE("/:id", c.DeleteUser)
	r.PUT("/:id", c.UpdateUser)
	r.POST("/", c.CreateUser)
	t := root.Group("/tool")
	t.GET("/migration/user", c.Migration)
}
