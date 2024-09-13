package controllers

import (
	"net/http"
	"strconv"

	"github.com/Ferdinand-work/go-crud/models"
	"github.com/Ferdinand-work/go-crud/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	msg, err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": msg})
}
func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Params.ByName("name")
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Params.ByName("name")
	err := uc.UserService.DeleteUser(&username)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (uc *UserController) GetByAge(ctx *gin.Context) {
	age, err := strconv.ParseInt(ctx.Param("age"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	users, err := uc.UserService.GetByAge(age)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) AddFriends(ctx *gin.Context) {
	username := ctx.Param("name")
	var friends models.FriendsReqBody
	if err := ctx.ShouldBindJSON(&friends); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resFriends, err := uc.UserService.AddFriends(&friends.Friends, username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	ctx.JSON(http.StatusOK, *resFriends)
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/get/:name", uc.GetUser)
	userRoute.GET("/getall", uc.GetAll)
	userRoute.PATCH("/update", uc.UpdateUser)
	userRoute.DELETE("/delete/:name", uc.DeleteUser)
	userRoute.GET("/getByAge/:age", uc.GetByAge)
	userRoute.POST("/addFriends/:name", uc.AddFriends)

}
