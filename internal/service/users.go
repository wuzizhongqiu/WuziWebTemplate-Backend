package service

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"wuzigoweb/api/http/errcode"
	pb "wuzigoweb/api/http/user"
	"wuzigoweb/internal/biz"
	"wuzigoweb/internal/data/model"
)

type UserService struct {
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	// 参数校验
	if req.Password != req.CheckPassword {
		return nil, fmt.Errorf("两次输入的密码不同\n")
	}
	// 做参数转换
	id, err := s.uc.CreateUser(ctx, &model.User{
		Account:  req.Account,
		Password: req.Password,
	})
	if id == -2 {
		return nil, errcode.ErrorInvalidParam("用户存在，无法注册")
	}
	if err != nil {
		return nil, err
	}
	return &pb.RegisterReply{Id: id}, nil
}

type JWTClaims struct {
	UserId   int64
	UserRole int32
	jwt.RegisteredClaims
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	// 参数校验：暂无

	// 参数转换
	user, err := s.uc.UserLogin(ctx, &model.User{
		Account:  req.Account,
		Password: req.Password,
	})
	if user == nil {
		fmt.Println("用户不存在")
		return nil, errcode.ErrorTodoNotFound("用户不存在")
	}

	// 登录成功
	var token *jwt.Token

	claims := JWTClaims{
		UserId:   user.ID,
		UserRole: user.UserRole,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "戊子仲秋",
			// ExpiresAt: time.Now().Add(), // 设置过期时间
		},
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	SignedToken, err := token.SignedString([]byte("template"))
	if err != nil {
		panic(err)
	}

	return &pb.LoginReply{Id: user.ID, Token: SignedToken}, nil
}

func (s *UserService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	user, err := s.uc.GetUserList(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.ListReply{
		Username:  user.Username,
		AvatarUrl: user.AvatarURL,
		UserInfo:  user.UserInfo,
	}, nil
}

func (s *UserService) GetCurrentUser(ctx context.Context, req *pb.GetCurrentUserRequest) (*pb.GetCurrentUserReply, error) {
	user, err := s.uc.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetCurrentUserReply{
		Id:        user.ID,
		Account:   user.Account,
		Password:  user.Password,
		Username:  user.Username,
		AvatarUrl: user.AvatarURL,
		Gender:    user.Gender,
		UserInfo:  user.UserInfo,
		UserRole:  user.UserRole,
	}, nil
}

func (s *UserService) ListUserByPage(ctx context.Context, req *pb.ListUserByPageRequest) (*pb.ListUserByPageReply, error) {
	user, total, err := s.uc.ListUserPage(ctx, req.Current, req.PageSize)
	if err != nil {
		return nil, err
	}
	userList := make([]*pb.UserInfo, req.PageSize)
	for i := range userList {
		if i > len(user)-1 {
			break
		}
		userList[i] = &pb.UserInfo{
			Account:   user[i].Account,
			Password:  user[i].Password,
			Username:  user[i].Username,
			AvatarUrl: user[i].AvatarURL,
			Gender:    user[i].Gender,
			UserInfo:  user[i].UserInfo,
			UserRole:  user[i].UserRole,
		}
	}
	return &pb.ListUserByPageReply{
		UserList: userList,
		Total:    int32(total),
	}, nil
}
