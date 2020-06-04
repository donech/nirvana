package v1

import (
	"github.com/donech/nirvana/internal/code"

	"github.com/unknwon/com"

	"github.com/donech/tool/xlog"

	"github.com/donech/nirvana/internal/domain/user/service"
	"github.com/gin-gonic/gin"
)

func NewUserController(userSimpleService *service.SimpleService) *UserController {
	return &UserController{UserSimpleService: userSimpleService}
}

type UserController struct {
	UserSimpleService *service.SimpleService
}

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

func (c UserController) GetUserList(ctx *gin.Context) {
	cursor := com.StrTo(ctx.Query("cursor")).MustInt64()
	size := com.StrTo(ctx.Query("size")).MustInt64()
	if size == 0 {
		size = 20
	}
	users, err := c.UserSimpleService.ItemsByCursor(ctx.Request.Context(), cursor, size+1)
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
