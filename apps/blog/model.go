package blog

type CreateBlogRequest struct {
	// Title of blog
	Title string `json:"title"`
	// Author of blog
	Author string `json:"author"`
	// CreateBy of blog
	CreateBy string `json:"create_by"`
	// Content of blog
	Content string `json:"content"`
	// Summary of blog
	Summary string `json:"summary"`
	// Tags
	Tags map[string]string `json:"tags" gorm:"serializer:json"`
}

type Blog struct {
	// id
	Id int64 `json:"id"`
	// CreateAtime
	CreatedAt int64 `json:"created_at"`
	// Update time
	UpdatedAt int64 `json:"updated_at"`
	// Publish time
	PublishedAt int64 `json:"published_at"`
	// Doc Status
	Status Status `json:"status"`
	// Audit time
	AuditedAt int64 `json:"audited_at"`
	// IsAudited Pass
	IsAuditedPass bool `json:"is_audited_pass"`
	// Create blog request
	*CreateBlogRequest
}
