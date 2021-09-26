package wrapper

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var defaultAllowOrigins = []string{"http://localhost"}


func makeCORSHandler(allowMethods []string ,allowOrigins []string)gin.HandlerFunc{
	corsConfig := cors.DefaultConfig()
	if allowOrigins == nil || len(allowOrigins) == 0{
		corsConfig.AllowOrigins = defaultAllowOrigins
	}else{
		corsConfig.AllowOrigins = allowOrigins
	}
	corsConfig.AllowMethods = allowMethods
	return cors.New(corsConfig)
}




type GinRouterGroup struct {
	*gin.RouterGroup
	AllowOrigins []string
}


func appendHandlerBefore(handler gin.HandlerFunc,handlers ...gin.HandlerFunc) []gin.HandlerFunc{
	return append([]gin.HandlerFunc{handler}, handlers...)
}

func (m *GinRouterGroup) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes  {
	CORSHandler := makeCORSHandler([]string{http.MethodGet}, m.AllowOrigins)
	hs := appendHandlerBefore(CORSHandler, handlers...)
	return  m.RouterGroup.GET(relativePath, hs...)
}
func (m *GinRouterGroup) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	CORSHandler := makeCORSHandler([]string{http.MethodPost}, m.AllowOrigins)
	hs := appendHandlerBefore(CORSHandler, handlers...)
	return m.RouterGroup.POST(relativePath, hs...)
}
func (m *GinRouterGroup) DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	CORSHandler := makeCORSHandler([]string{http.MethodDelete}, m.AllowOrigins)
	hs := appendHandlerBefore(CORSHandler, handlers...)
	return m.RouterGroup.DELETE(relativePath, hs...)
}
func (m *GinRouterGroup) PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	CORSHandler := makeCORSHandler([]string{http.MethodPatch}, m.AllowOrigins)
	hs := appendHandlerBefore(CORSHandler, handlers...)
	return m.RouterGroup.PATCH(relativePath, hs...)
}

func (m *GinRouterGroup) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	CORSHandler := makeCORSHandler([]string{http.MethodOptions}, m.AllowOrigins)
	hs := appendHandlerBefore(CORSHandler, handlers...)
	return m.RouterGroup.OPTIONS(relativePath, hs...)
}
func (m *GinRouterGroup) PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	CORSHandler := makeCORSHandler([]string{http.MethodPut}, m.AllowOrigins)
	hs := appendHandlerBefore(CORSHandler, handlers...)
	return m.RouterGroup.PUT(relativePath, hs...)
}

func (m *GinRouterGroup) HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	CORSHandler := makeCORSHandler([]string{"HEAD"}, m.AllowOrigins)
	hs := appendHandlerBefore(CORSHandler, handlers...)
	return m.RouterGroup.HEAD(relativePath, hs...)
}

func (m *GinRouterGroup) Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	CORSHandler := makeCORSHandler([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}, m.AllowOrigins)
	hs := appendHandlerBefore(CORSHandler, handlers...)
	return m.RouterGroup.Any(relativePath, hs...)
}


func (m *GinRouterGroup) Group(relativePath string, handlers ...gin.HandlerFunc) *GinRouterGroup {
	_AllowOrigins := make([]string, len(m.AllowOrigins))
	copy(_AllowOrigins,m.AllowOrigins)

	return &GinRouterGroup{
		AllowOrigins: _AllowOrigins,
		RouterGroup: m.RouterGroup.Group(relativePath, handlers...),
	}
}
