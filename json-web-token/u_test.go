package json_web_token

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

var Token string//在login里面create

func TestCreate(t *testing.T) {
	token:=Create("shiina18",1)
	fmt.Println(token)
	//fmt.Println(CheckToken(token))
	//res, _ := base64.StdEncoding.DecodeString("eyJpc3MiOiJyZWRyb2NrIiwiZXhwIjoiMTU4MzA1OTE4MCIsImlhdCI6IjE1ODMwNDgzODAiLCJ1c2VybmFtZSI6InNoaWluYTE4IiwiVWlkIjoxfQ==")
	//fmt.Println(string(res))
}

func Meddle(c *gin.Context) {//中间件
	uid, username, err := CheckToken(Token)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 10001, "err":"token有错"})
		c.Abort()
		return
	}
	c.Set("uid", uid)
	c.Set("username", username)
	c.Next()
	return
}

//更新的函数里面，读取taken的值
//id,_:=c.Get("uid")
//user.Id=id.(int)