package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StudentCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudentCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StudentCreateLogic {
	return &StudentCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StudentCreateLogic) StudentCreate(req *types.STStudentCreateReq) (rsp *types.STStudentCreateRsp, err error) {
	rsp = &types.STStudentCreateRsp{}
	if req.Name == "" || req.Age == 0 || req.Gender == 0 {
		return rsp, errors.New("参数错误")
	}

	student := &model.Student{
		Name:   req.Name,
		Age:    req.Age,
		Gender: req.Gender,
	}
	if err = l.svcCtx.DB.Create(student).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	return rsp, nil
}
