package api

import (
	"net/http"

	"github.com/ACK-lcn/Blog/apps/token"
	"github.com/ACK-lcn/Blog/exception"
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

// Logout Handler function.
func (h *TokenApiHandler) Logout(c *gin.Context) {
	// Create and parse logout request parameters.
	inReq := token.NewLogoutRequest()
	if err := c.BindJSON(inReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}

	// Call the Service layer to handle the logout logic.
	err := h.svc.Logout(c.Request.Context(), inReq)
	if err != nil {
		// Handle specific errors.
		if exception.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Token not found or already invalid"})
			return
		}
	}

	// Return success response（200）.
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
