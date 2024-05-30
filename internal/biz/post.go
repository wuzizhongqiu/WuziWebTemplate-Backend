package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"wuzigoweb/internal/data/model"
)

type PostRepo interface {
	Save(context.Context, *model.Post) (int64, error)
	//Login(context.Context, *model.User) (*model.User, error)
	//Update(context.Context, *Greeter) (*Greeter, errcode)
	//FindByID(context.Context, int64) (*Greeter, errcode)
	//ListByHello(context.Context, string) ([]*Greeter, errcode)
	//ListAll(context.Context) ([]*Greeter, errcode)
}

type PostUsecase struct {
	repo PostRepo
	log  *log.Helper
}

func NewPostUsecase(repo PostRepo, logger log.Logger) *PostUsecase {
	return &PostUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PostUsecase) CreateAPost(ctx context.Context, post *model.Post) (int64, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 CreateAPost 创建帖子操作")
	return uc.repo.Save(ctx, post)
}
