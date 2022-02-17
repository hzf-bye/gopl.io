package main

import (
	"fmt"
	"net/http"
)
/*
go 基础 http 库，只要是实现了 ServeHTTP(*w* http.ResponseWriter, *req* http.Request) 这个方法就可以给挂到 http 里面。

这里我们声明了一个 Engine 的结构体，里面我们声明了一个 map 类型的路由，这是为了方便挂载不同的方法。
同时因为 func(http.ResponseWriter, *http.Request) 这个方法会被经常用到，所以我们对他进行了重命名。
最关键的就是 ServeHTTP 方法了：
当我们的请求过来时，我们把他的 method 和 path 两个参数，合并起来，当做 key 去 router 里面查找这请求的处理方法，你也可以使用其他的处理方式。
最后就是 run 起来，把自己挂载到 http 里面。
 */

// 从定义方法
type HandleFun func(http.ResponseWriter, *http.Request)

// 核心结构体
type Engine struct {
	router map[string]HandleFun
}

func New() *Engine {
	return &Engine{
		router: make(map[string]HandleFun,0),
	}
}

//ServeHTTP 关键方法 实现 Handler 的接口
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler,ok := e.router[key];ok {
		handler(w, req)
	}else{
		fmt.Fprintf(w,"404 Not Found: %s\n", req.URL)
	}
}

// 发射
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) AddRouter(method string, pattern string, handler HandleFun) {
	key := method + "-" + pattern
	e.router[key] = handler
}

func (e *Engine) GET(pattern string, handler HandleFun) {
	e.AddRouter("GET", pattern, handler)
}
