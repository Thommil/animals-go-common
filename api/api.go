package api

import (
	"github.com/gin-gonic/gin"
)

// Resource defines an API endpoint and data
type Resource struct {
	Engine *gin.Engine
}

// IRoutable must be implemented to add routes to a router
type IRoutable interface {
	ApplyRoutes() *Resource
}
