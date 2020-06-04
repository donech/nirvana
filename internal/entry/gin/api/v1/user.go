package v1

import (
	"math/rand"
	"strconv"

	"github.com/unknwon/com"

	"github.com/donech/core/xlog"

	"github.com/prometheus/common/log"

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
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Fatalln("id must convert failed", err.Error())
	}
	user, err := c.UserSimpleService.ItemByID(ctx.Request.Context(), id)
	if err != nil {
		xlog.Ctx(ctx.Request.Context()).Error("some thing wrong", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"user": user})
}

func (c UserController) GetUserList(ctx *gin.Context) {
	cursor := com.StrTo(ctx.Query("cursor")).MustInt64()
	size := com.StrTo(ctx.Query("size")).MustInt64()
	if size == 0 {
		size = 20
	}
	users, err := c.UserSimpleService.ItemsByCursor(ctx.Request.Context(), cursor, size+1)
	if err != nil {
		xlog.Ctx(ctx.Request.Context()).Error("some thing wrong", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	cursor = 0
	hasMore := false
	if int64(len(users)) > size {
		cursor = users[size].ID
		hasMore = true
		users = users[0:size]
	}
	ctx.JSON(200, gin.H{
		"list": users,
		"pager": Pager{
			Cursor:  cursor,
			Size:    size,
			HasMore: hasMore,
		}})
}

func (c UserController) CreateUser(ctx *gin.Context) {
	data := map[string]interface{}{
		"name":  "solar" + strconv.Itoa(rand.Int()),
		"phone": "18001023261",
	}
	user, err := c.UserSimpleService.Create(ctx.Request.Context(), data)
	if err != nil {
		xlog.Ctx(ctx.Request.Context()).Error("some thing wrong", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"status": "success", "user": user})
}

func (c UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Fatalln("id must convert failed", err.Error())
	}
	data := map[string]interface{}{
		"name":  "solar" + strconv.Itoa(rand.Int()),
		"phone": "18001023261",
	}
	err = c.UserSimpleService.Update(ctx.Request.Context(), id, data)
	if err != nil {
		xlog.Ctx(ctx.Request.Context()).Error("some thing wrong", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"status": "success"})
}

func (c UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Fatalln("id must convert failed", err.Error())
	}
	err = c.UserSimpleService.Delete(ctx.Request.Context(), id)
	if err != nil {
		xlog.Ctx(ctx.Request.Context()).Error("some thing wrong", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"status": "success"})
}

func (c UserController) Migration(ctx *gin.Context) {
	c.UserSimpleService.Migration()
	ctx.JSON(200, gin.H{"status": "success"})
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
