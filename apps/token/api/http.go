package api

import (
	"net/http"

	"github.com/ACK-lcn/Blog/apps/token"
	"github.com/gin-gonic/gin"
)

type TokenApiHandler struct {
	svc token.Service
}

func NewTokenApiHandler(tokenServiceImpl token.Service) *TokenApiHandler {
	return &TokenApiHandler{
		svc: tokenServiceImpl,
	}
}

// Router Register
func (h *TokenApiHandler) Register(r gin.IRouter) {
	v1 := r.Group("v1")
	v1.POST("/api/blog/v1/tokens/", h.Login)
	v1.DELETE("/api/blog/v1/tokens/", h.Logout)
}

// Login Handler Function. (Handler only handles HTTP requests and responses.）
func (h *TokenApiHandler) Login(c *gin.Context) {
	// Retrieve user request parameters and return them in JSON format.
	inReq := token.NewLoginRequest()
	err := c.BindJSON(inReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ins, err := h.svc.Login(c.Request.Context(), inReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Return the response in JSON format.
	c.JSON(http.StatusOK, ins)
}

func (h *TokenApiHandler) Logout(*gin.Context) {}
