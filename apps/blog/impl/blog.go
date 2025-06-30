package impl

import (
	"context"

	"github.com/ACK-lcn/Blog/apps/blog"
)

// Create Blog
func (i *blogServiceImpl) CreateBlog(ctx context.Context, in *blog.CreateBlogRequest) (*blog.Blog, error) {
	ins := blog.NewBlog(in)
	if err := i.db.WithContext(ctx).Create(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

// Query Blog
func (i *blogServiceImpl) QueryBlog(ctx context.Context, in *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	query := i.db.WithContext(ctx).Model(&blog.Blog{})
	set := blog.NewBlogSet()
	if in.Status != nil {
		query = query.Where("Status = ?", *in.Status)
	}
	if in.Keywords != "" {
		query = query.Where("title LIKE ?", "%"+in.Keywords+"%")
	}
	if len(in.Usernames) > 0 {
		query = query.Where("create_by In ?", in.Usernames)
	}

	// Query total
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	// Query a page of data
	err = query.Order("create_at DESC").Offset(in.Offset()).Limit(in.PageSize).Find(&set.Items).Error
	if err != nil {
		return nil, err
	}

	return set, nil
}

// Describe Blog
func (i *blogServiceImpl) DescribeBlog(ctx context.Context, in *blog.DescribeBlogRequest) (*blog.Blog, error) {
	return nil, nil
}

// Update Blog Status
func (i *blogServiceImpl) updateBlogStatus(ctx context.Context, in *blog.UpdateBlogStatusRequest) (*blog.Blog, error) {
	return nil, nil
}

// Update Blog
func (i *blogServiceImpl) UpdateBlog(ctx context.Context, in *blog.UpdateBlogRequest) (*blog.Blog, error) {
	return nil, nil
}

// Delete Blog
func (i *blogServiceImpl) DeleteBlog(ctx context.Context, in *blog.DeleteBlogRequest) error {
	return nil
}

// Audit Blog
func (i *blogServiceImpl) AuditBlog(ctx context.Context, in *blog.AuditBlogRequest) (*blog.Blog, error) {
	return nil, nil
}
