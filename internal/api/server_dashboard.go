package api

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

// dashboardHTML is a static usage dashboard. It holds no secrets: the page asks for
// the management key and reads quotas through /v8/management with it.
//
//go:embed dashboard.html
var dashboardHTML []byte

// serveDashboard serves GET /dashboard.
func (s *Server) serveDashboard(c *gin.Context) {
	c.Header("Cache-Control", "no-store")
	c.Data(http.StatusOK, "text/html; charset=utf-8", dashboardHTML)
}
