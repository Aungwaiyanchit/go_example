package controllers

import (
	"fmt"
	"net/http"

	"github.com/Aungwaiyanchit/books/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Select("email").Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreatUser(c *gin.Context) {
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var olduser models.User
	if err := models.DB.Where("email = ?", input.Email).First(&olduser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already register"})
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	user := models.User{Email: input.Email, Password: string(bytes)}
	models.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"message": "sigin sucessfully"})
}

func LoginUser(c *gin.Context) {
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var user models.User
	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect password"})
		return
	}
	token, err := generateJWT(user.Email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Println(token)
	c.JSON(http.StatusOK, gin.H{"message": "login success", "token": token, "email": user.Email})
}

func generateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JwtUser{
		RegisteredClaims: jwt.RegisteredClaims{},
		Email: username,
	})
	tokenString, error := token.SignedString([]byte("kitty harein"))
	if error != nil {
		return "", error
	}
	return tokenString, nil
}
