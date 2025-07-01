package impl

import (
	"context"
	"fmt"
	"time"

	"dario.cat/mergo"
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
	query := i.db.WithContext(ctx).Model(&blog.Blog{})
	ins := blog.NewBlog(blog.NewCreateBlogRequest())

	// SELECT * FROM `blog` WHERE id = '66'
	err := query.Where("id = ?", in.BlogId).Find(ins).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Update Blog (Includes: full update and incremental update.)
func (i *blogServiceImpl) UpdateBlog(ctx context.Context, in *blog.UpdateBlogRequest) (*blog.Blog, error) {
	// Query the objects that need to be updated
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(in.BlogId))
	if err != nil {
		return nil, err
	}

	// update Model
	switch in.UpdateMode {
	case blog.UPDATE_MODE_PUT:
		// PUT （Full update）
		ins.CreateBlogRequest = in.CreateBlogRequest
	case blog.UPDATE_MODE_PATCH:
		// PATCH （incremental update）

		// Use “dario.cat/mergo” golang library, mergo function, to merge structures and mappings.
		// Mergo source code address: "https://github.com/darccio/mergo"
		if err := mergo.Merge(ins.CreateBlogRequest, in.CreateBlogRequest, mergo.WithOverride); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown update mode: %d", in.UpdateMode)
	}

	// Update data
	// One: Update Time
	ins.UpdatedAt = time.Now().Unix()

	// Two:Update data
	// SQL Snytax: UPDATE `blog` SET `created_at`=1696809853,`updated_at`=1696809853,`status`='1',`title`='blog Web Service Api2',`content`='Go',`tags`='{"分类":"Go"}' WHERE id = '666' AND `id` = 666
	if err = i.db.WithContext(ctx).Where("id = ?", in.BlogId).Updates(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

// Update Blog Status
func (i *blogServiceImpl) updateBlogStatus(ctx context.Context, in *blog.UpdateBlogStatusRequest) (*blog.Blog, error) {
	return nil, nil
}

// Delete Blog (SQL Snytax: DELETE FROM `blog` WHERE id = '666')
func (i *blogServiceImpl) DeleteBlog(ctx context.Context, in *blog.DeleteBlogRequest) error {
	return i.db.WithContext(ctx).
		Model(&blog.Blog{}).
		Where("id =?", in.BlogId).
		Delete(&blog.Blog{}).
		Error
}

// Audit Blog
func (i *blogServiceImpl) AuditBlog(ctx context.Context, in *blog.AuditBlogRequest) (*blog.Blog, error) {
	return nil, nil
}
