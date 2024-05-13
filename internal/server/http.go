package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/golang-jwt/jwt/v5"
	v2 "wuzigoweb/api/http/post"
	v1 "wuzigoweb/api/http/user"

	"wuzigoweb/internal/conf"
	"wuzigoweb/internal/service"
)

// MiddlewareTest 自定义中间件示例（定制化中间件的实现）
func MiddlewareTest(opts ...string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// 执行 Handler 之前做点事情
				fmt.Println("MiddlewareTest 执行 Handler 前", tr)
				defer func() {
					fmt.Println("MiddlewareTest 执行 Handler 后", tr)
				}()
			}
			return handler(ctx, req)
		}
	}
}

func ParseJwtWithClaims(tokenString string, options ...jwt.ParserOption) (*service.JWTClaims, error) {
	if tokenString == "" {
		return nil, errors.New("鉴权失败")
	}

	println("准备进行 jwt.ParseWithClaims")

	token, err := jwt.ParseWithClaims(tokenString, &service.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("template"), nil
	}, options...)
	if err != nil {
		return nil, err
	}

	println("校验 Claims 对象是否有效，基于 exp（过期时间），nbf（不早于），iat（签发时间）等进行判断（如果有这些声明的话）。")

	// 校验 Claims 对象是否有效，基于 exp（过期时间），nbf（不早于），iat（签发时间）等进行判断（如果有这些声明的话）。
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token.Claims.(*service.JWTClaims), nil
}

// MiddlewareJWTUser 登录后的用户可以访问的鉴权
func MiddlewareJWTUser() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenString := tr.RequestHeader().Get("token")
				println("tr 接收到的东西是：", tokenString)
				claims, err := ParseJwtWithClaims(tokenString)
				if err != nil {
					return -1, err
				}
				if claims.UserRole == 30 {
					return -1, err
				}
			}
			return handler(ctx, req)
		}
	}
}

// MiddlewareJWTAdmin 管理者才能访问的鉴权
func MiddlewareJWTAdmin() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenString := tr.RequestHeader().Get("token")
				println("tr 接收到的东西是：", tokenString)
				claims, err := ParseJwtWithClaims(tokenString)
				if err != nil {
					return -1, err
				}
				if claims.UserRole != 20 {
					return -1, err
				}
			}
			return handler(ctx, req)
		}
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, user *service.UserService, post *service.PostService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(), // 获取堆栈信息（便于排错）

			validate.Validator(), // 执行参数校验规则

			selector.Server( // 对非登录注册操作进行鉴权
				MiddlewareJWTUser()).
				Prefix("/http.post.Post/").
				Path("/http.user.User/List").
				Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	opts = append(opts, http.ResponseEncoder(responseEncoder)) // 自定义响应格式
	opts = append(opts, http.ErrorEncoder(errorEncoder))

	srv := http.NewServer(opts...)
	v1.RegisterUserHTTPServer(srv, user) // 注册 user 的服务
	v2.RegisterPostHTTPServer(srv, post) // 注册 post 的服务

	return srv
}
