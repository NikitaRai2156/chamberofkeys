package httpapi

import (
	ck "the_chamber_of_keys/pkg/chamber_of_keys"

	"github.com/gin-gonic/gin"
)

// NewRouter: creates and configures a Gin engine with routes backed by the ChamberOfKeys implementation
func NewRouter(chamber ck.ChamberOfKeys) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	h := &Handler{ChamberOfKeys: chamber}

	r.GET("/string/:key", h.GetString)
	r.POST("/string", h.InsertString)
	r.PUT("/string/:key", h.UpdateString)

	r.POST("/list/:key/items", h.PushItem)
	r.DELETE("/list/:key/items", h.PopItem)

	r.DELETE("/:key", h.Remove)

	return r
}
