package rrf

import (
	"fmt"
	"net/http"
	"strings"
)

type Context struct{
	req         *http.Request
	w           http.ResponseWriter
	queryParam  map[string]string//GET请求的获取参数
	formParam   map[string]string//POST请求的获取参数
}

func (c *Context) String(s string){//输出格式string
	_,_ = c.w.Write([]byte(s))
}

func NewContext(rw http.ResponseWriter, r *http.Request) (ctx Context){
	ctx = Context{
		req:       r,
		w:         rw,
	}
	ctx.queryParam = parseQuery(r.RequestURI)//GET请求的获取参数
	ctx.formParam = parseForm(r)//POST请求的获取参数
	Group = make([]Middleware,0,0)
	return
}

func (c *Context) Query(key string) string{
	v := c.queryParam[key]
	return v
}

func parseQuery(uri string) (res map[string]string){//GET请求的获取参数，从URL中截取
	res = make(map[string]string)
	uris := strings.Split(uri,"?")//第一次截取
	if len(uris) == 1 {
		return
	}
	param := uris[len(uris)-1]
	pair := strings.Split(param,"&")//第二次截取
	for _,kv := range pair {
		kvPair := strings.Split(kv,"=")//第三次截取
		if len(kvPair) != 2 {
			fmt.Println(kvPair)
			panic("request error")
		}
		res[kvPair[0]] = kvPair[1]
	}
	return
}