package blog

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/ACK-lcn/Blog/common"
)

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

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize:   10,
		PageNumber: 1,
		Usernames:  []string{},
	}
}

type QueryBlogRequest struct {
	// page size
	PageSize int `json:"page_size"`
	// Current page
	PageNumber int `json:"page_number"`
	// "0" DRAFT: indicates draft status, to query all blogs
	// "nil": indicates no such filter condition
	// "1" PUBLISHED: indicates "published status"
	Status *Status `json:"status"`
	// Keyword search based on article title
	Keywords string `json:"keywords"`
	// Query which users' blogs
	Usernames []string `json:"usernames"`
}

func (r *QueryBlogRequest) Offset() int {
	return int(r.PageSize * (r.PageNumber - 1))
}

func (r *QueryBlogRequest) ParsePageSize(ps string) {
	psInt, _ := strconv.ParseInt(ps, 10, 64)
	if psInt != 0 {
		r.PageSize = int(psInt)
	}
}

func (r *QueryBlogRequest) ParsePageNumber(pn string) {
	psInt, _ := strconv.ParseInt(pn, 10, 64)
	if psInt != 0 {
		r.PageNumber = int(psInt)
	}
}

func (r *QueryBlogRequest) SetStatus(s Status) {
	r.Status = &s
}

func NewDescribeBlogRequest(id string) *DescribeBlogRequest {
	return &DescribeBlogRequest{
		BlogId: id,
	}

}

type DescribeBlogRequest struct {
	BlogId string `json:"blog_id"`
}

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

type BlogSet struct {
	// Total number of blogs
	Total int64 `json:"total"`
	// Return a page of data
	Items []*Blog `json:"items"`
}

func (b *BlogSet) String() string {
	dj, _ := json.Marshal(b)
	return string(dj)
}

func (s *BlogSet) Add(Items ...*Blog) {
	s.Items = append(s.Items, Items...)
}

type UpdateBlogStatusRequest struct {
	BlogId int64  `json:"blog_id"`
	Status Status `json:"status"`
}

func NewAuditBlogRequest(id string) *AuditBlogRequest {
	return &AuditBlogRequest{
		BlogId: id,
	}
}

type AuditBlogRequest struct {
	BlogId        string `json:"blog_id"`
	IsAuditedPass bool   `json:"id_audited_pass"`
}

// NewPutUpdateBlogRequest create a put update request.
func NewPutUpdateBlogRequest(id string) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:            id,
		UpdateMode:        UPDATE_MODE_PUT,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

// NewPatchUpdateBlogRequest create a patch update request.
func NewPatchUpdateBlogRequest(id string) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:            id,
		UpdateMode:        UPDATE_MODE_PATCH,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

type UpdateBlogRequest struct {
	// article id
	BlogId string `json:"blog_id"`
	// The scope of the blog is not passed in by the user, but is automatically filled in by the API interface layer
	Scope *common.Scope `json:"scope"`
	// update mode  (Full update/incremental update)
	UpdateMode UpdateMode `json:"update_mode"`
	// User update request, the user only passed a tag
	*CreateBlogRequest
}

func NewDeleteBlogRequest(id string) *DeleteBlogRequest {
	return &DeleteBlogRequest{
		BlogId: id,
	}
}

type DeleteBlogRequest struct {
	// article id
	BlogId string `json:"blog_id"`
}
