package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"wuzigoweb/internal/biz"
	"wuzigoweb/internal/data/model"
)

type postRepo struct {
	data *Data
	log  *log.Helper
}

func NewPostRepo(data *Data, logger log.Logger) biz.PostRepo {
	return &postRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *postRepo) Save(ctx context.Context, post *model.Post) (int64, error) {
	u := r.data.query.Post
	err := u.WithContext(ctx).Create(post)
	if err != nil {
		return -1, err
	}
	return post.ID, nil
}
