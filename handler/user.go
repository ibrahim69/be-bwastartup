package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct registeruserinput
	// struct di atas kita passing sebagai parameter service
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessages := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errMessages)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token, err := h.jwtService.GenerateToken()

	formatter := user.FormatUser(newUser, "")

	response := helper.APIResponse("Account has ben registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	// user memasukkan input (email & password)
	// input ditangkap handler
	// mapping dari input user ke input struct
	// input struct passing ke service
	// di service mencari dg bantuan repository user dengan email x
	// mencocokkan password

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessages := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errMessages)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedinUser, err := h.userService.LoginInput(input)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentokentoken")

	response := helper.APIResponse("Successfully logged in", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
