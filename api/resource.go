package resource

import (
	"github.com/gin-gonic/gin"
)

// Routable interface should return a gin Router.Group based on
// routing implemented by the underlying value
type Routable interface {
	GetGroup() *gin.RouterGroup
}
