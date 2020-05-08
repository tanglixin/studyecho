package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main(){
	e := echo.New()

	//注册一个Get请求, 路由地址为: /tizi365  并且绑定一个控制器函数, 这里使用的是闭包函数。
	//e.GET("/tizi365", func(c echo.Context) error {
	//	//控制器函数直接返回一个字符串，http响应状态为http.StatusOK，就是200状态。
	//	return c.String(http.StatusOK, "欢迎访问tizi365.com")
	//})


	//echo 路由
	//定义post请求, url为：/users, 绑定saveUser控制器函数
	//e.POST("/users", saveUser)

	//定义get请求，url模式为：/users/:id  （:id是参数，例如: /users/10, 会匹配这个url模式），绑定getUser控制器函数
	e.GET("/users", getUser)
	e.GET("/user", getList)

	//定义put请求
	//e.PUT("/users/:id", updateUser)

	//定义delete请求
	//e.DELETE("/users/:id", deleteUser)

	e.Start(":8080")
}


// 路由定义：e.GET("/users/:id", getUser)
// getUser控制器函数实现
func getUser(c echo.Context) error {
	// 获取url上的path参数，url模式里面定义了参数:id
	id := c.Param("id")
	//username := c.Param("username")
	//usertype := c.Param("type")

	//获取url上的参数，例如：/users?username=tizi365&type=2
	username := c.QueryParam("username")  //值为："tizi365"
	usertype := c.QueryParam("type")      //值为："2"

	return c.String(http.StatusOK, username)

	//响应一个字符串，这里直接把id以字符串的形式返回给客户端。
	return c.String(http.StatusOK, username)
	return c.String(http.StatusOK, usertype)
	return c.String(http.StatusOK, id)

}





//定义json结构体
type User struct {
	Id       int
	Username string
}


//响应请求
func getList(c echo.Context)error{
	//返回html
	html := "<html><head><title>测试标题</title></head></html>"
	return c.HTML(http.StatusOK,html)


	//返回json
	u := User{1,"userlist"}
	return c.JSON(http.StatusAccepted,u)

	//返回字符串
	return c.String(http.StatusOK,"返回的是一个字符串")
}

