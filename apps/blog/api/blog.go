package api

import (
	"github.com/ACK-lcn/Blog/apps/blog"
	"github.com/ACK-lcn/Blog/apps/token"
	"github.com/ACK-lcn/Blog/apps/user"
	"github.com/ACK-lcn/Blog/common"
	"github.com/ACK-lcn/Blog/exception"
	"github.com/ACK-lcn/Blog/middlewarre"
	"github.com/ACK-lcn/Blog/response"
	"github.com/gin-gonic/gin"
)

// Router Register
func (h *apiHandler) Register(r gin.IRouter) {
	v1 := r.Group("v1").Group("blog")
	v1.POST("/", h.CreateBlog)
	v1.GET("/", h.QueryBlog)
	v1.GET("/:id", h.DescribeBlog)

	// Backend management interface requires authentication
	v1.Use(middlewarre.NewTokenAuther().Auth)

	v1.PUT("/:id", middlewarre.Required(user.ROLE_AUTHOR), h.UpdateBlog)
	v1.PATCH("/:id", middlewarre.Required(user.ROLE_AUTHOR), h.PatchBlog)
	v1.DELETE("/:id", middlewarre.Required(user.ROLE_AUTHOR), h.DeleteBlog)
	v1.POST("/:id/audit", middlewarre.Required(user.ROLE_AUDITOR), h.AuditBlog)
}

// CreateBlog
func (h *apiHandler) CreateBlog(c *gin.Context) {
	tkObj := c.Keys[token.TOKEN_GIN_KEY_NAME]
	tk := tkObj.(*token.Token)

	in := blog.NewCreateBlogRequest()
	err := c.BindJSON(in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	in.CreateBy = tk.UserName
	ins, err := h.svc.CreateBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, ins)
}

// Query the article list, list query, there is no need to query the specific content of the article.
func (h *apiHandler) QueryBlog(c *gin.Context) {
	in := blog.NewQueryBlogRequest()
	in.Keywords = c.Query("keywords")
	in.ParsePageSize(c.Query("page_size"))
	in.ParsePageNumber(c.Query("page_number"))

	switch c.Query("status") {
	case "draft":
		in.SetStatus(blog.STATUS_DRAFT)
	case "published":
		in.SetStatus(blog.STATUS_PUBLISHED)
	}

	set, err := h.svc.QueryBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, set)
}

// Details page, try to find more content
func (h *apiHandler) DescribeBlog(c *gin.Context) {
	in := blog.NewDescribeBlogRequest(c.Param("id"))
	ins, err := h.svc.DescribeBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// PUT （Full update）
func (h *apiHandler) UpdateBlog(c *gin.Context) {
	in := blog.NewPutUpdateBlogRequest(c.Param("id"))
	err := c.BindJSON(in.CreateBlogRequest)
	if err != nil {
		response.Failed(c, err)
		return
	}

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

	in.Scope = &common.Scope{
		Username: tk.UserName,
	}

	ins, err := h.svc.UpdateBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// PATCH （incremental update）
func (h *apiHandler) PatchBlog(c *gin.Context) {
	in := blog.NewPatchUpdateBlogRequest(c.Param("id"))
	err := c.BindJSON(in.CreateBlogRequest)
	if err != nil {
		response.Failed(c, err)
		return
	}

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

	in.Scope = &common.Scope{
		Username: tk.UserName,
	}

	ins, err := h.svc.UpdateBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// Delete article
func (h *apiHandler) DeleteBlog(c *gin.Context) {
	in := blog.NewDeleteBlogRequest(c.Param("id"))
	err := h.svc.DeleteBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "ok")
}

// Article review, only articles that pass the review can be seen.
func (h *apiHandler) AuditBlog(c *gin.Context) {
	in := blog.NewAuditBlogRequest(c.Param("id"))
	err := c.BindJSON(in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	ins, err := h.svc.AuditBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, ins)
}
