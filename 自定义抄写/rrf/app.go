package rrf

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

type handlerMap map[string]Handler

type Handler func(*Context)

type App struct {
	router map[string]handlerMap
}

func Default() *App{//头
	return &App{
		router: make(map[string]handlerMap),
	}
}

func (a *App) GET(uri string,handle Handler){//方法
	a.handle("GET",uri,handle)
}

func (a *App) POST(uri string,handle Handler){//方法
	a.handle("POST",uri,handle)
}

func (a *App) handle(method,uri string,handler Handler){
	handlers,ok := a.router[method]
	if !ok{
		m :=make(handlerMap)
		a.router[method] = m
		handlers = m
	}
	_,ok = handlers[uri]
	if ok {
		panic("same route")
	}
	handlers[uri] = handler
}

func (a *App) Run(port int){//尾
	ports := strconv.FormatInt(int64(port),10)
	http.Handle("/",Chain(a,Group))
	if err := http.ListenAndServe(":"+ports,nil); err != nil {
		log.Fatal(err.Error)
	}
}

func (a *App) ServeHTTP(w http.ResponseWriter, req *http.Request){
	httpMethod := req.Method
	uri :=req.RequestURI
	uris := strings.Split(uri,"?")
	if len(uris) < 1{
		return
	}
	handlers,ok := a.router[httpMethod]
	if !ok {
		log.Println("may by a hacker:",req.RemoteAddr)
		return
	}
	h,ok := handlers[uris[0]]
	if !ok {
		Handler404(w,req)
		return
	}
	c := NewContext(w,req)
	h(&c)
}

func Handler404(w http.ResponseWriter, req *http.Request){
	_, _ = w.Write([]byte("404 not found"))
}