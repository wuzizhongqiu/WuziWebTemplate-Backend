package biz

import (
	"context"
	"wuzigoweb/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	Save(context.Context, *model.User) (int64, error)
	Login(context.Context, *model.User) (*model.User, error)
	ListUser(ctx context.Context, id int64) (*model.User, error)
	GetUser(ctx context.Context) (*model.User, error)
	ListUserPage(ctx context.Context, current, pageSize int32) ([]*model.User, int, error)
	//Update(context.Context, *Greeter) (*Greeter, errcode)
	//FindByID(context.Context, int64) (*Greeter, errcode)
	//ListByHello(context.Context, string) ([]*Greeter, errcode)
	//ListAll(context.Context) ([]*Greeter, errcode)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 CreateUser 创建用户操作")
	return uc.repo.Save(ctx, user)
}

func (uc *UserUsecase) UserLogin(ctx context.Context, user *model.User) (*model.User, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 UserLogin 用户登录逻辑")
	return uc.repo.Login(ctx, user)
}

func (uc *UserUsecase) GetUserList(ctx context.Context, id int64) (*model.User, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 GetUserList 获取用户列表")
	return uc.repo.ListUser(ctx, id)
}

func (uc *UserUsecase) GetCurrentUser(ctx context.Context) (*model.User, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 GetCurrentUser 获取当前用户信息")
	return uc.repo.GetUser(ctx)
}

func (uc *UserUsecase) ListUserPage(ctx context.Context, current, pageSize int32) ([]*model.User, int, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 ListUserPage 获取用户分页逻辑")
	return uc.repo.ListUserPage(ctx, current, pageSize)
}
