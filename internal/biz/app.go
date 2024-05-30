package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"wuzigoweb/internal/data/model"
)

type AppRepo interface {
	Save(context.Context, *model.App) error
	Update(context.Context, int64, *model.App) error
	ListApp(ctx context.Context, current, pageSize int32) ([]*model.App, int, error)
	GetByID(context.Context, int64) (*model.App, *model.User, error)
	//Login(context.Context, *model.User) (*model.User, error)
	//ListByHello(context.Context, string) ([]*Greeter, errcode)
	//ListAll(context.Context) ([]*Greeter, errcode)
}

type AppUsecase struct {
	repo AppRepo
	log  *log.Helper
}

func NewAppUsecase(repo AppRepo, logger log.Logger) *AppUsecase {
	return &AppUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AppUsecase) CreatApp(ctx context.Context, app *model.App) error {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 CreatApp 创建应用")
	return uc.repo.Save(ctx, app)
}

func (uc *AppUsecase) ModifyApp(ctx context.Context, id int64, app *model.App) error {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 ModifyApp 修改应用信息")
	return uc.repo.Update(ctx, id, app)
}

func (uc *AppUsecase) ListAppPage(ctx context.Context, current, pageSize int32) ([]*model.App, int, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 ListAppPage 应用分页逻辑")
	return uc.repo.ListApp(ctx, current, pageSize)
}

func (uc *AppUsecase) GetAppById(ctx context.Context, id int64) (*model.App, *model.User, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 GetAppById 查询应用信息")
	return uc.repo.GetByID(ctx, id)
}
