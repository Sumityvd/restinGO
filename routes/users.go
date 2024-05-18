package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)
//  Signup creates a new user
//
//	@Summary		Create new user
//	@Description		Create a new service user
//	@Id			CreateUser
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserExample		true	"User to create"
//	@Success		201	{object}	models.UserCreated 			"User is created successfully"
//	@Failure		500	{object}	models.UserCreatedUnsuccessful 		"User Already existing"
//	@Router			/signup [post]
func Signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
// Login user and get JWT tokens
//
//	@Summary		Login
//	@Description		Login to get JWT token
//	@Id			Login
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			user		body		models.UserExample	true	"Login credentials"
//	@Success		200		{object}	models.UserLogin		"Successful login response"
//	@Failure		401		{object}	models.UserLoginUnsuccessful	"Unsuccessful login response"
//	@Router			/login [post]
func Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
