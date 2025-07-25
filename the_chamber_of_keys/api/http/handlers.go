package httpapi

import (
	"net/http"
	ck "the_chamber_of_keys/pkg/chamber_of_keys"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	param_key      = "key"
	param_position = "position"
	param_front    = "front"
	param_back     = "back"
	defaultTTL     = 15 * time.Minute
)

type Handler struct {
	ChamberOfKeys ck.ChamberOfKeys
}

func NewHandler(c ck.ChamberOfKeys) *Handler {
	return &Handler{
		ChamberOfKeys: c,
	}
}

// GetString(): Handles GET /string/:key - returns the string value for a given key
func (h *Handler) GetString(c *gin.Context) {

	key := c.Param(param_key)

	val, err := h.ChamberOfKeys.GetString(key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"value": val})

}

// InsertString(): Handles POST /string - inserts a new string value with optional TTL
func (h *Handler) InsertString(c *gin.Context) {

	var req struct {
		Key   string        `json:"key" binding:"required"`
		Value string        `json:"value" binding:"required"`
		TTL   time.Duration `json:"ttl"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// set default ttl if not provided
	if req.TTL == 0 {
		req.TTL = defaultTTL
	}

	errInsert := h.ChamberOfKeys.InsertString(req.Key, req.Value, req.TTL)
	if errInsert != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errInsert.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "string inserted"})

}

// UpdateString(): Handles PUT /string/:key — updates the value of an existing key
func (h *Handler) UpdateString(c *gin.Context) {

	key := c.Param(param_key)

	var req struct {
		Value string `json:"value" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	errUpdate := h.ChamberOfKeys.UpdateString(key, req.Value)
	if errUpdate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errUpdate.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "string updated"})
}

// PushItem(): Handles POST /list/:key/items?position=front|back — pushes an item to front or back of list
func (h *Handler) PushItem(c *gin.Context) {

	key := c.Param(param_key)
	position := c.Query(param_position)

	if position == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "position not specified"})
		return
	}

	var req struct {
		Value string        `json:"value" binding:"required"`
		TTL   time.Duration `json:"ttl"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if req.TTL == 0 {
		req.TTL = defaultTTL
	}

	var errPush error

	switch position {
	case param_front:
		errPush = h.ChamberOfKeys.PushFront(key, req.Value, req.TTL)
	case param_back:
		errPush = h.ChamberOfKeys.PushBack(key, req.Value, req.TTL)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid position"})
		return
	}

	if errPush != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errPush.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item pushed"})

}

// PopItem(): Handles DELETE /list/:key/items?position=front|back — pops an item from front or back of list
func (h *Handler) PopItem(c *gin.Context) {

	key := c.Param(param_key)
	position := c.Query(param_position)

	if position == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "position not specified"})
		return
	}

	var val string
	var errPop error

	switch position {
	case param_front:
		val, errPop = h.ChamberOfKeys.PopFront(key)
	case param_back:
		val, errPop = h.ChamberOfKeys.PopBack(key)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid position"})
		return
	}

	if errPop != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errPop.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"value": val})
}

// Remove(): Handles DELETE /:key — removes the key from store
func (h *Handler) Remove(c *gin.Context) {

	key := c.Param(param_key)

	if err := h.ChamberOfKeys.Remove(key); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item removed"})

}
