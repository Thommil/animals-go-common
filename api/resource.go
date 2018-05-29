package resource

import (
	"github.com/gin-gonic/gin"
)

// Routable must be implemented to add routes to a router
type Routable interface {
	GetGroup() *gin.RouterGroup
}
