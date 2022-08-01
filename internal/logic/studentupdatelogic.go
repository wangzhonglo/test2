package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StudentUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StudentUpdateLogic {
	return &StudentUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StudentUpdateLogic) StudentUpdate(req *types.STStudentUpdateReq) (rsp *types.STStudentUpdateRsp, err error) {
	rsp = &types.STStudentUpdateRsp{}

	if req.Id == 0 {
		return rsp, errors.New("参数错误")
	}

	student := &model.Student{
		Name:   req.Name,
		Age:    req.Age,
		Gender: req.Gender,
	}
	query := l.svcCtx.DB.Model(&model.Student{})
	query = query.Where("id=?", req.Id)
	if err = query.Updates(&student).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	return
}
