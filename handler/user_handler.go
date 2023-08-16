package handler

import (
	"easylib-go/middleware"
	"easylib-go/model"
	"easylib-go/usecase"
	"easylib-go/utils"
	"errors"
	"fmt"

	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usrUseCase usecase.UserUsecase
}

func (usrHandler UserHandler) GetUserById(ctx *gin.Context) {
	session := sessions.Default(ctx)
	existSession := session.Get("Id")
	id, ok := existSession.(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success":      false,
			"errorMessage": "User session not found or invalid",
		})
		return
	}
	usr, err := usrHandler.usrUseCase.GetUserById(id)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("UserHandler.GetUserById() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("UserHandler.GetUserById() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": err.Error(),
				
			})
			return
		}
		return
	}
	if usr == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usr,
	})
}

func (usrHandler UserHandler) AddUser(ctx *gin.Context) {
	usr := &model.User{}
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = usrHandler.usrUseCase.AddUser(usr)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("UserHandler.InsertUser() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("UserHandler.InsertUser() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving user data",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully added user",
	})
}

func (usrHandler UserHandler) UpdateUser(ctx *gin.Context) {
	usr := &model.User{}
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = usrHandler.usrUseCase.UpdateUser(usr, ctx)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("UserHandler.UpdateUser() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("UserHandler.UpdateUser() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while updating user data",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User updated successfully",
	})
}

func (usrHandler *UserHandler) DeleteUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	existSession := session.Get("Username")
	userName, ok := existSession.(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success":      false,
			"errorMessage": "User session not found or invalid",
		})
		return
	}
	if err := usrHandler.usrUseCase.DeleteUser(userName); err != nil {
		fmt.Printf("UserHandler.DeleteUser(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User deleted successfully",
	})
}

func NewUserHandler(srv *gin.Engine, usrUseCase usecase.UserUsecase) *UserHandler {
	usrHandler := &UserHandler{
		usrUseCase: usrUseCase,
	}

	// route
	srv.POST("/user", usrHandler.AddUser)
	srv.PUT("/user", middleware.RequireToken(), usrHandler.UpdateUser)
	srv.GET("/user", middleware.RequireToken(), usrHandler.GetUserById)
	srv.DELETE("/user", middleware.RequireToken(), usrHandler.DeleteUser)
	return usrHandler
}