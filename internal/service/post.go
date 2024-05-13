package service

import (
	"context"
	"wuzigoweb/internal/biz"
	"wuzigoweb/internal/data/model"

	pb "wuzigoweb/api/http/post"
)

type PostService struct {
	uc *biz.PostUsecase
}

func NewPostService(uc *biz.PostUsecase) *PostService {
	return &PostService{uc: uc}
}

func (s *PostService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostReply, error) {
	// 参数校验 validate 基本上都做了

	// 参数转换
	id, err := s.uc.CreateAPost(ctx, &model.Post{
		Title:   req.Title,
		Content: req.Text,
		Tags:    req.Tags,
		UserID:  req.UserId,
	})

	if err != nil {
		return nil, err
	}
	return &pb.CreatePostReply{Id: id}, nil
}
