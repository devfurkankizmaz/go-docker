package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", Hello)
	router.GET("/hello", Greeter)
	router.Run("localhost:8080")

}

func Hello(c *gin.Context) {
	str := fmt.Sprintf("%s%s", "Hello from ", runtime.GOOS)
	c.JSON(http.StatusOK, str)
	fmt.Println(str)
}

func Greeter(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, "Greetings "+name)
}
