package main

import (
	"github.com/echo"
	"net/http"
)

//定义一个结构体接收参数
type User struct{
	Name string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}
func main(){
	e := echo.New()
	e.GET("getName",nameHandle)
	e.GET("/getEmail/:email", func(c echo.Context) error {
		email := c.Param("email")
		//如果请求url为: /users/tizi365 则name的值为tizi365
		//Param获取的值也是String类型
		return c.String(http.StatusOK, email)
	})
	e.Start(":8081")
}

func nameHandle(c echo.Context)error  {
	//name := c.FormValue("name")
	name := c.FormValue("name")
	return c.String(http.StatusOK, name)
}
