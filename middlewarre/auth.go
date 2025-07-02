package middlewarre

import (
	"fmt"
	"net/http"

	"github.com/ACK-lcn/Blog/apps/token"
	"github.com/ACK-lcn/Blog/apps/user"
	"github.com/ACK-lcn/Blog/exception"
	"github.com/ACK-lcn/Blog/ioc"
	"github.com/ACK-lcn/Blog/response"
	"github.com/gin-gonic/gin"
)

func NewTokenAuther() *TokenAuther {
	return &TokenAuther{
		tk: ioc.Controller().Get(token.AppName).(token.Service),
	}
}

// Middleware for authentication
// Middleware for Token authentication
type TokenAuther struct {
	tk   token.Service
	role user.Role
}

// Gin middleware func(*Context)
func (a *TokenAuther) Auth(c *gin.Context) {
	// 1. Get Token
	at, err := c.Cookie(token.TOKEN_COOKIE_NAME)
	if err != nil {
		if err == http.ErrNoCookie {
			response.Failed(c, token.CookieNotFound)
			return
		}
		response.Failed(c, err)
		return
	}

	// 2. Call the Token module for authentication
	in := token.NewValiateToken(at)
	tk, err := a.tk.ValiateToken(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	// Put the result of authentication: tk into the request context for later business logic use
	if c.Keys == nil {
		c.Keys = map[string]any{}
	}
	c.Keys[token.TOKEN_GIN_KEY_NAME] = tk
}

// Authorization is performed when the user has been authenticated
// Determine the role of the current user
func (a *TokenAuther) Perm(c *gin.Context) {
	tkObj := c.Keys[token.TOKEN_GIN_KEY_NAME]
	if tkObj == nil {
		response.Failed(c, exception.NewPermissionDeny("token not found"))
		return
	}

	tk, ok := tkObj.(*token.Token)
	if !ok {
		response.Failed(c, exception.NewPermissionDeny("token not an *token.Token"))
		return
	}

	fmt.Printf("user %s role %d \n", tk.UserName, tk.Role)

	// If it is Admin, then release it directly
	if tk.Role == user.ROLE_ADMIN {
		return
	}

	if tk.Role != a.role {
		response.Failed(c, exception.NewPermissionDeny("role %d not allow", tk.Role))
		return
	}
}

// Write Gin middleware with parameters
func Required(r user.Role) gin.HandlerFunc {
	a := NewTokenAuther()
	a.role = r
	return a.Perm
}
