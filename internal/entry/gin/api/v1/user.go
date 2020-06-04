package v1

import (
	"math/rand"
	"strconv"

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
	r.GET("/:id", c.GetUser)
	r.DELETE("/:id", c.DeleteUser)
	r.PUT("/:id", c.UpdateUser)
	r.POST("/", c.CreateUser)
	t := root.Group("/tool")
	t.GET("/migration/user", c.Migration)
}
