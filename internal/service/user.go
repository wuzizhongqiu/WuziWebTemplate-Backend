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
	fmt.Println("我是 service 层")
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
