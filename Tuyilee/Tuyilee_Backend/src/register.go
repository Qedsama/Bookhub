package main

import (
    "encoding/json"
    "net/http"
    "os"
    "io"
    "github.com/gin-gonic/gin"
	"fmt"
)

func registerHandler(c *gin.Context) {
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

	for _, existingUser := range users {
		if existingUser.Username == user.Username {
			c.JSON(http.StatusConflict, gin.H{"error": "用户名已存在"})
			return
		}
	}

	users = append(users, user)
	file.Seek(0, 0)
	file.Truncate(0)

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Received user data: %+v\n", user)

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}
