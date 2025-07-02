package impl

import (
	"context"

	"github.com/ACK-lcn/Blog/apps/blog"
	"github.com/ACK-lcn/Blog/common"
)

func (i *blogServiceImpl) update(ctx context.Context, scope *common.Scope, ins *blog.Blog) error {
	exec := i.db.WithContext(ctx).Where("id = ?", ins.Id)

	if scope != nil {
		if scope.Username != "" {
			exec = exec.Where("create_by = ?", scope.Username)
		}
	}

	exec = exec.Updates(ins)
	return exec.Error
}
