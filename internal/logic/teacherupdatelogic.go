package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeacherUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeacherUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeacherUpdateLogic {
	return &TeacherUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeacherUpdateLogic) TeacherUpdate(req *types.STTeacherUpdateReq) (rsp *types.STTeacherUpdateRsp, err error) {
	rsp = &types.STTeacherUpdateRsp{}

	if req.Id == 0 {
		return rsp, errors.New("参数错误")
	}

	teacher := &model.Teacher{
		Name:      req.Name,
		TeacherId: req.TeacherId,
		Step:      req.Step,
		Course:    req.Course,
	}
	query := l.svcCtx.DB.Model(&model.Teacher{})
	query = query.Where("id=?", req.Id)
	if err = query.Updates(&teacher).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	return
}
