package gee

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Router %4s - %s", method, pattern)
	//1.拆分Url
	parts := parsePattern(pattern)
	//2.拼接处理器的key
	key := method + "-" + pattern
	//3.根据请求方式获取前缀树树根，不存在则构建
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	//4.将路由作为一个结点插入到前缀树中
	r.roots[method].insert(pattern, parts, 0)
	//5.保存处理器映射
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	//1.分解url
	searchParts := parsePattern(path)
	//2.待封装的参数映射
	params := make(map[string]string)
	//3.获取该请求方法下的根节点
	root, ok := r.roots[method]
	//4.不存在则直接返回
	if !ok {
		return nil, nil
	}
	//5.根据url搜索前缀树
	n := root.search(searchParts, 0)
	//6.返回的节点不为空
	if n != nil {
		//7.分解该节点存在的路由
		parts := parsePattern(n.pattern)
		//8.遍历路由各部分
		for index, part := range parts {
			//9.对于动态路由，将当前部分作为参数封装到map中
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			//10.对于通配符，将剩余部分作为整体，并插入分隔符/
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		//将当前路由的控制器加入中间件执行队列中
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(context *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
