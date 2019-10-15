package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Any uint8 = iota
	GET
	POST
	DELETE
	PATCH
	PUT
	OPTIONS
	HEAD
	Static
	StaticFile
	StaticFS
)

type LeafRouter struct {
	leafType        uint8
	rawRelativePath string
	option          interface{}
	middleware      []gin.HandlerFunc
	afterHandlers   []gin.HandlerFunc
}

func (l *LeafRouter) Use(handlers ...gin.HandlerFunc) *LeafRouter {
	l.middleware = append(l.middleware, handlers...)
	return l
}

func (l *LeafRouter) Build(g gin.IRouter) {

	switch l.leafType {
	case GET:
		g.GET(l.rawRelativePath, append(l.middleware, l.afterHandlers...)...)
	case POST:
		g.POST(l.rawRelativePath, append(l.middleware, l.afterHandlers...)...)
	case DELETE:
		g.DELETE(l.rawRelativePath, append(l.middleware, l.afterHandlers...)...)
	case PATCH:
		g.PATCH(l.rawRelativePath, append(l.middleware, l.afterHandlers...)...)
	case PUT:
		g.PUT(l.rawRelativePath, append(l.middleware, l.afterHandlers...)...)
	case OPTIONS:
		g.OPTIONS(l.rawRelativePath, append(l.middleware, l.afterHandlers...)...)
	case HEAD:
		g.HEAD(l.rawRelativePath, append(l.middleware, l.afterHandlers...)...)
	case Static:
		g.Group("/", l.middleware...).Static(l.rawRelativePath, l.option.(string))
	case StaticFile:
		g.Group("/", l.middleware...).StaticFile(l.rawRelativePath, l.option.(string))
	case StaticFS:
		g.Group("/", l.middleware...).StaticFS(l.rawRelativePath, l.option.(http.FileSystem))
	default:
		panic(fmt.Errorf("bad leaf type %v", l.leafType))
	}
}

func NewLeafRouter(leafType uint8, path string, handlers ...gin.HandlerFunc) *LeafRouter {
	return &LeafRouter{
		leafType:        leafType,
		rawRelativePath: path,
		afterHandlers:   handlers,
	}
}

func NewLeafRouterOp(leafType uint8, path string, option interface{}, handlers ...gin.HandlerFunc) *LeafRouter {
	switch leafType {
	case Static:
		if _, ok := option.(string); !ok {
			panic(fmt.Errorf("must serve as router.Static(path, root string)"))
		}
	case StaticFile:
		if _, ok := option.(string); !ok {
			panic(fmt.Errorf("must serve as router.Static(path, filepath string)"))
		}
	case StaticFS:
		if _, ok := option.(http.FileSystem); !ok {
			panic(fmt.Errorf("must serve as router.Static(path, filesystem http.FileSystem)"))
		}
	}
	return &LeafRouter{
		leafType:        leafType,
		rawRelativePath: path,
		option:          option,
		middleware:      handlers,
	}
}

type Router struct {
	handlers           []gin.HandlerFunc
	SubRouter          map[string]*Router
	SameLevelSubRouter map[string]*Router
	Leafs              map[string]*LeafRouter
}

func NewRouterGroup(handlers ...gin.HandlerFunc) *Router {
	return &Router{
		handlers: handlers,
	}
}

func (c *Router) Build(g gin.IRouter) {
	g.Use(c.handlers...)
	for k, sr := range c.SubRouter {
		sr.Build(g.Group(k))
	}
	for _, sr := range c.SameLevelSubRouter {
		sr.Build(g.Group(""))
	}
	for _, leaf := range c.Leafs {
		leaf.Build(g)
	}
}

func (c *Router) Group(relativePath string, handlers ...gin.HandlerFunc) (r *Router) {
	if c.SubRouter == nil {
		c.SubRouter = make(map[string]*Router)
	} else if _, ok := c.SubRouter[relativePath]; ok {
		panic(fmt.Errorf("must conflict path %v", relativePath))
	}
	r = NewRouterGroup(handlers...)
	c.SubRouter[relativePath] = r
	return
}

func (c *Router) Extend(name string, handlers ...gin.HandlerFunc) (r *Router) {
	if c.SameLevelSubRouter == nil {
		c.SameLevelSubRouter = make(map[string]*Router)
	} else if _, ok := c.SameLevelSubRouter[name]; ok {
		panic(fmt.Errorf("must conflict name %v", name))
	}
	r = NewRouterGroup(handlers...)
	c.SameLevelSubRouter[name] = r
	return
}

func (c *Router) Use(handlers ...gin.HandlerFunc) *Router {
	c.handlers = append(c.handlers, handlers...)
	return c
}

func (c *Router) leaf(leafType uint8, relativePath string, handlers ...gin.HandlerFunc) (r *LeafRouter) {
	if c.Leafs == nil {
		c.Leafs = make(map[string]*LeafRouter)
	} else if _, ok := c.Leafs[relativePath+string(leafType)]; ok {
		panic(fmt.Errorf("must conflict path %v", relativePath+string(leafType)))
	}
	r = NewLeafRouter(leafType, relativePath, handlers...)
	c.Leafs[relativePath+string(leafType)] = r
	return
}

func (c *Router) leafOp(leafType uint8, relativePath string, op interface{}, handlers ...gin.HandlerFunc) (r *LeafRouter) {
	if c.Leafs == nil {
		c.Leafs = make(map[string]*LeafRouter)
	} else if _, ok := c.Leafs[relativePath]; ok {
		panic(fmt.Errorf("must conflict path %v", relativePath))
	}
	r = NewLeafRouterOp(leafType, relativePath, op, handlers...)
	c.Leafs[relativePath] = r
	return
}

func (c *Router) GET(relativePath string, handlers ...gin.HandlerFunc) *LeafRouter {
	return c.leaf(GET, relativePath, handlers...)
}

func (c *Router) POST(relativePath string, handlers ...gin.HandlerFunc) *LeafRouter {
	return c.leaf(POST, relativePath, handlers...)
}

func (c *Router) DELETE(relativePath string, handlers ...gin.HandlerFunc) *LeafRouter {
	return c.leaf(DELETE, relativePath, handlers...)
}

func (c *Router) PATCH(relativePath string, handlers ...gin.HandlerFunc) *LeafRouter {
	return c.leaf(PATCH, relativePath, handlers...)
}

func (c *Router) PUT(relativePath string, handlers ...gin.HandlerFunc) *LeafRouter {
	return c.leaf(PUT, relativePath, handlers...)
}

func (c *Router) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) *LeafRouter {
	return c.leaf(OPTIONS, relativePath, handlers...)
}

func (c *Router) HEAD(relativePath string, handlers ...gin.HandlerFunc) *LeafRouter {
	return c.leaf(HEAD, relativePath, handlers...)
}

func (c *Router) StaticFile(relativePath string, filepath string, handlers ...gin.HandlerFunc) *LeafRouter {
	return c.leafOp(HEAD, relativePath, filepath, handlers...)
}

func (c *Router) Static(relativePath string, root string, handlers ...gin.HandlerFunc) *LeafRouter {
	return c.leafOp(HEAD, relativePath, root, handlers...)
}

func (c *Router) StaticFS(relativePath string, fs http.FileSystem, handlers ...gin.HandlerFunc) *LeafRouter {
	return c.leafOp(HEAD, relativePath, fs, handlers...)
}
