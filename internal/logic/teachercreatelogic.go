package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeacherCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeacherCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeacherCreateLogic {
	return &TeacherCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeacherCreateLogic) TeacherCreate(req *types.STTeacherCreateReq) (rsp *types.STTeacherCreateRsp, err error) {
	rsp = &types.STTeacherCreateRsp{}
	if req.Name == "" || req.TeacherId == "" {
		return rsp, errors.New("参数错误")
	}

	teacher := &model.Teacher{
		Name:      req.Name,
		Course:    req.Course,
		Step:      req.Step,
		TeacherId: req.TeacherId,
	}
	if err = l.svcCtx.DB.Create(teacher).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	return rsp, nil
}
