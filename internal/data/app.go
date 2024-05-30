package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"wuzigoweb/internal/biz"
	"wuzigoweb/internal/data/model"
)

type appRepo struct {
	data *Data
	log  *log.Helper
}

func NewAppRepo(data *Data, logger log.Logger) biz.AppRepo {
	return &appRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *appRepo) Save(ctx context.Context, app *model.App) error {
	u := r.data.query.App
	return u.WithContext(ctx).Create(app)
}

func (r *appRepo) Update(ctx context.Context, id int64, app *model.App) error {
	u := r.data.query.App
	_, err := u.WithContext(ctx).Where(u.ID.Eq(id)).Updates(app)
	if err != nil {
		return err
	}
	return nil
}

func (r *appRepo) ListApp(ctx context.Context, current, pageSize int32) ([]*model.App, int, error) {
	u := r.data.query.App
	apps, err := u.WithContext(ctx).Offset(int(current - 1)).Limit(int(pageSize)).Find()
	count, err := u.WithContext(ctx).Count()
	if err != nil {
		return nil, 0, err
	}
	return apps, int(count), nil
}

func (r *appRepo) GetByID(ctx context.Context, id int64) (*model.App, *model.User, error) {
	u := r.data.query.App
	app, err := u.WithContext(ctx).Where(u.ID.Eq(id)).First()
	user, err := r.data.query.User.WithContext(ctx).Where(r.data.query.User.ID.Eq(app.UserID)).First()
	if err != nil {
		return nil, nil, err
	}
	return app, user, nil
}
