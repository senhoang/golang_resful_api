package userapi

import (
	"fmt"
	"net/http"
	"strings"
	"web_kenda_api/pkg/database"
	"web_kenda_api/pkg/middlewares"
	"web_kenda_api/pkg/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	// Check quyền đăng nhập
	_, err := middlewares.VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	user := []models.User{}

	tx := database.GetPostgre().Raw("select * from Users order by Id").Scan(&user)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": tx.Error.Error()})
		return
	}

	fmt.Println(tx)

	if tx.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserByDept(c *gin.Context) {
	fmt.Println(c.Param("deptid"))
	// Check quyền đăng nhập
	_, err := middlewares.VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	user := []models.User{}
	deptid := strings.ToUpper(c.Param("deptid"))
	fmt.Println(deptid)

	tx := database.GetPostgre().Raw("select * from Users where DeptID = '" + deptid + "' order by Id").Scan(&user)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": tx.Error.Error()})
		return
	}

	if tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	_, err := middlewares.VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var user models.User

	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	fmt.Println(user)

	tx := database.GetPostgre().Exec("insert into Users values ('" + user.Id + "','" + user.Name + "','" + user.Email + "','" + user.Password + "','" + user.Deptid + "','" + user.Roles + "')")
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": tx.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inserted"})
}

func UpdateUser(c *gin.Context) {
	_, err := middlewares.VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var user models.User

	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	tx := database.GetPostgre().Exec(
		" update Users set" +
			" Name = '" + user.Name + "'," +
			" Email = '" + user.Email + "'," +
			" Password = '" + user.Password + "'," +
			" Deptid = '" + user.Deptid + "'," +
			" Roles = '" + user.Roles + "'" +
			" where Id = '" + user.Id + "'")

	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": tx.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Edited"})
}

func DeleteUser(c *gin.Context) {
	_, err := middlewares.VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	id := c.Param("id")
	fmt.Println(id)

	tx := database.GetPostgre().Exec("delete from Users where ID = '" + id + "'")
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": tx.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})

}
