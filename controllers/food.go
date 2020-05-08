package main

import "net/http"

func getUser(c echo.Context) error {
	// 获取url上的path参数，url模式里面定义了参数:id
	id := c.Param("id")

	//响应一个字符串，这里直接把id以字符串的形式返回给客户端。
	return c.String(http.StatusOK, id)
}