package rrf

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type H map[string]interface{}

func (c *Context) JSON(obj interface{}){//输出格式json
	b, err := json.Marshal(obj)
	if err != nil {
		log.Printf("error:", err)
	}
	_, _ = c.w.Write(b)
}

func (c *Context) PostForm(key string) string{
	v := c.formParam[key]
	return v
}

func parseForm ( r *http.Request) (res map[string]string){//POST请求的获取参数，从BODY中截取
	res = make(map[string]string)
	_ = r.ParseForm()
	for k, v := range r.Form {
		res[k] = strings.Join(v, "")
	}
	return
}