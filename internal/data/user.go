package data

import (
	"context"
	"fmt"
	"wuzigoweb/internal/biz"
	"wuzigoweb/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, user *model.User) (int64, error) {
	u := r.data.query.User
	// 查看用户名是否被占用
	users, _ := u.WithContext(ctx).Where(u.Account.Eq(user.Account)).First()
	if users != nil {
		return -2, nil
	}
	// 存入数据库
	err := u.WithContext(ctx).Save(user)
	if err != nil {
		return -3, fmt.Errorf("用户注册存入数据库失败")
	}
	return user.ID, err
}

func (r *userRepo) Login(ctx context.Context, user *model.User) (*model.User, error) {
	u := r.data.query.User
	users, err := u.WithContext(ctx).Where(u.Account.Eq(user.Account)).First()
	if err != nil {
		r.log.WithContext(ctx).Debug("查询用户信息失败：", err)
		return nil, err
	}
	if users == nil {
		r.log.WithContext(ctx).Debugf("users 是空")
		return nil, err
	}
	if users.Password != user.Password {
		return nil, err
	}

	return users, err
}

func (r *userRepo) ListUser(ctx context.Context, id int64) (*model.User, error) {
	u := r.data.query.User
	users, err := u.WithContext(ctx).Where(u.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) GetUser(ctx context.Context) (*model.User, error) {
	u := r.data.query.User
	users, err := u.WithContext(ctx).Where(u.ID.Eq(ctx.Value("user_id").(int64))).First()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) ListUserPage(ctx context.Context, current, pageSize int32) ([]*model.User, int, error) {
	u := r.data.query.User
	users, err := u.WithContext(ctx).Offset(int(current - 1)).Limit(int(pageSize)).Find()
	count, err := u.WithContext(ctx).Count()
	if err != nil {
		return nil, 0, err
	}
	return users, int(count), nil
}
