package blog

import "context"

const (
	AppName = "blog"
)

// Blog module interface
type Service interface {
	// Create blog
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	// Query the article list, list query, there is no need to query the specific content of the article.
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	// Details page, try to find more content
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	// Modify article status
	updateBlogStatus(context.Context, *UpdateBlogStatusRequest) (*Blog, error)
	// Update article
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	//  Delete article
	DeleteBlog(context.Context, *DeleteBlogRequest) error
	// Article review, only articles that pass the review can be seen.
	AuditBlog(context.Context, *AuditBlogRequest) (*Blog, error)
}

type QueryBlogRequest struct {
}

type DescribeBlogRequest struct {
	BlogId int64 `json:"blog_id"`
}

type BlogSet struct {
	// Total number of blogs
	Total int64 `json:"total"`
	// Return a page of data
	Items []*Blog `json:"items"`
}

type UpdateBlogStatusRequest struct {
	BlogId int64  `json:"blog_id"`
	Status Status `json:"status"`
}

type AuditBlogRequest struct {
	BlogId        int64 `json:"blog_id"`
	IsAuditedPass bool  `json:"id_audited_pass"`
}

type UpdateBlogRequest struct {
	// article id
	BlogId int64 `json:"blog_id"`
	// update mode  (Full update/incremental update)
	UpdateMode UpdateMode `json:"update_mode"`
	// User update request, the user only passed a tag
	*CreateBlogRequest
}

type DeleteBlogRequest struct {
	// article id
	BlogId int64 `json:"blog_id"`
}
