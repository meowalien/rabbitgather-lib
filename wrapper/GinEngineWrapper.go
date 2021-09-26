package wrapper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


// GinEngine provide some useful feature for github.com/gin-gonic/gin don't have.
type GinEngine struct {
	*gin.Engine
	// setting AllowOrigins will open the CORSHandler
	AllowOrigins []string
}

func (m *GinEngine) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes  {
	CORSHandler := makeCORSHandler([]string{http.MethodGet}, m.AllowOrigins)
	hs := appendHandlerBefore(CORSHandler, handlers...)
	return  m.RouterGroup.GET(relativePath, hs...)
}
func (m *GinEngine) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	CORSHandler := makeCORSHandler([]string{http.MethodPost}, m.AllowOrigins)
	hs := appendHandlerBefore(CORSHandler, handlers...)
	return m.RouterGroup.POST(relativePath, hs...)
}
func (m *GinEngine) DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	CORSHandler := makeCORSHandler([]string{http.MethodDelete}, m.AllowOrigins)
	hs := appendHandlerBefore(CORSHandler, handlers...)
	return m.RouterGroup.DELETE(relativePath, hs...)
}
func (m *GinEngine) PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	CORSHandler := makeCORSHandler([]string{http.MethodPatch}, m.AllowOrigins)
	hs := appendHandlerBefore(CORSHandler, handlers...)
	return m.RouterGroup.PATCH(relativePath, hs...)
}

func (m *GinEngine) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	CORSHandler := makeCORSHandler([]string{http.MethodOptions}, m.AllowOrigins)
	hs := appendHandlerBefore(CORSHandler, handlers...)
	return m.RouterGroup.OPTIONS(relativePath, hs...)
}
func (m *GinEngine) PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	CORSHandler := makeCORSHandler([]string{http.MethodPut}, m.AllowOrigins)
	hs := appendHandlerBefore(CORSHandler, handlers...)
	return m.RouterGroup.PUT(relativePath, hs...)
}

func (m *GinEngine) Group(relativePath string, handlers ...gin.HandlerFunc) *GinRouterGroup {
	_AllowOrigins := make([]string, len(m.AllowOrigins))
	copy(_AllowOrigins,m.AllowOrigins)

	return &GinRouterGroup{
		AllowOrigins: _AllowOrigins,
		RouterGroup: m.RouterGroup.Group(relativePath, handlers...),
	}
}
