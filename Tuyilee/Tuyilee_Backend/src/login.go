package main

import (
    "encoding/json"
    "net/http"
    "os"
    "io"
    "github.com/gin-gonic/gin"
	"fmt"
)

func loginHandler(c *gin.Context) {
	fmt.Println("Received a login request")
	var user User

	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := os.OpenFile("../user_data.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	var users []User
	decoder = json.NewDecoder(file)
	if err := decoder.Decode(&users); err != nil && err != io.EOF {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var foundUser User
	userExists := false

	for _, existingUser := range users {
		if existingUser.Username == user.Username {
			userExists = true
			foundUser = existingUser
			break
		}
	}

	if !userExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户名不存在"})
		return
	}

	if user.Password != foundUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}