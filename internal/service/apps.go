package service

import (
	"context"
	"wuzigoweb/internal/biz"
	"wuzigoweb/internal/data/model"

	pb "wuzigoweb/api/http/app"
)

type AppService struct {
	uc *biz.AppUsecase
}

func NewAppService(uc *biz.AppUsecase) *AppService {
	return &AppService{uc: uc}
}

func (s *AppService) CreatApp(ctx context.Context, req *pb.CreatAppRequest) (*pb.CreatAppReply, error) {
	err := s.uc.CreatApp(ctx, &model.App{
		AppName:         req.AppName,
		AppDesc:         &req.AppDesc,
		AppIcon:         &req.AppIcon,
		AppType:         req.AppType,
		ScoringStrategy: req.ScoringStrategy,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreatAppReply{}, nil
}

func (s *AppService) ModifyApp(ctx context.Context, req *pb.ModifyAppRequest) (*pb.ModifyAppReply, error) {
	err := s.uc.ModifyApp(ctx, req.AppId, &model.App{
		AppName:         req.AppModInfo.AppName,
		AppDesc:         &req.AppModInfo.AppDesc,
		AppIcon:         &req.AppModInfo.AppIcon,
		AppType:         req.AppModInfo.AppType,
		ScoringStrategy: req.AppModInfo.ScoringStrategy,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ModifyAppReply{}, nil
}

func (s *AppService) ListAppPage(ctx context.Context, req *pb.ListAppPageRequest) (*pb.ListAppPageReply, error) {
	app, total, err := s.uc.ListAppPage(ctx, req.Current, req.PageSize)
	if err != nil {
		return nil, err
	}
	appList := make([]*pb.AppInfo, req.PageSize)
	for i := range appList {
		if i > len(app)-1 {
			break
		}
		appList[i] = &pb.AppInfo{
			AppId:           app[i].ID,
			AppName:         app[i].AppName,
			AppDesc:         *app[i].AppDesc,
			AppIcon:         *app[i].AppIcon,
			AppType:         app[i].AppType,
			ScoringStrategy: app[i].ScoringStrategy,
			UserId:          app[i].UserID,
		}
	}
	return &pb.ListAppPageReply{
		AppList: appList,
		Total:   int32(total),
	}, nil
}

func (s *AppService) GetAppById(ctx context.Context, req *pb.GetAppByIdRequest) (*pb.GetAppByIdReply, error) {
	app, user, err := s.uc.GetAppById(ctx, req.AppId)
	if err != nil {
		return nil, err
	}
	apps := &pb.AppInfo{
		AppId:           app.ID,
		AppName:         app.AppName,
		AppDesc:         *app.AppDesc,
		AppIcon:         *app.AppIcon,
		AppType:         app.AppType,
		ScoringStrategy: app.ScoringStrategy,
		UserId:          app.UserID,
	}
	users := &pb.UserInfo{
		Id:        user.ID,
		Account:   user.Account,
		Password:  user.Password,
		Username:  user.Username,
		AvatarUrl: user.AvatarURL,
		Gender:    user.Gender,
		UserInfo:  user.UserInfo,
		UserRole:  user.UserRole,
	}
	return &pb.GetAppByIdReply{App: apps, User: users}, nil
}
