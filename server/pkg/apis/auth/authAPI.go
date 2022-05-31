package auth

import (
	"fmt"
	"net/http"
	"web_kenda_api/pkg/database"
	"web_kenda_api/pkg/middlewares"
	"web_kenda_api/pkg/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	auth := models.Auth{}
	user := models.User{}

	// Gán vào json gửi từ client vào auth
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	println(auth.Username)
	println(auth.Password)

	// Truy vấn kiểm tra đăng nhập với GORM
	tx := database.GetPostgre().Raw("select * from users where Id = '" + auth.Username + "' and Password = '" + auth.Password + "'").Scan(&user)
	// tx := database.GetPostgre().Raw("select * from users where Id = 'senhoang' and Password = '123456'").Scan(&user)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": tx.Error.Error()})
		// c.JSON(http.StatusInternalServerError, gin.H{"message": "login failed"})
		return
	}
	fmt.Print(user)

	// tx.RowsAffected đến số dữ liệu
	if tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "The username or password is incorrect"})
		return
	}

	//GenerateToken
	token, err := middlewares.GenerateToken(auth.Username)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Generate token failed"})
		return
	}

	// roles := []string{}
	// roles = append(roles, "ROLE_ADMINISTRATOR")
	// roles = append(roles, "ROLE_MODERATOR")

	fmt.Println(token)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Info(c *gin.Context) {
	_, err := middlewares.VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		// "roles":        "admin",
		"avatar": "my-avatar.png",
		"name":   "Administrator",
	})
}
