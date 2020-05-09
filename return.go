package main

import (
	"github.com/echo"
	"net/http"
)

//定义结构体
type user struct {
	Name string `json:"name"`
	Age string `json:"age"`
}

func main(){
	e := echo.New()
	e.POST("returnJson", testHandle)
	e.Start(":8082")
}

func testHandle(c echo.Context)error  {
	//name := c.Param("name")
	//age := c.Param("age")
	u := user{
		Name: "name",
		Age: "12"}

	return c.JSON(http.StatusOK,u)
}


