package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StuCreatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStuCreatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StuCreatLogic {
	return &StuCreatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StuCreatLogic) StuCreat(req *types.STStuCreateReq) (rsp *types.STStuCreateRsp, err error) {
	rsp = &types.STStuCreateRsp{}
	if req.Name == "" || req.Age == 0 || req.Sex == 0 {
		return rsp, errors.New("参数错误")
	}

	stu := &model.Stu{
		Name:  req.Name,
		Age:   req.Age,
		Sex:   req.Sex,
		Grade: req.Grade,
		Step:  req.Step,
		StuId: req.StuId,
	}
	if err = l.svcCtx.DB.Create(stu).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	return rsp, nil
}
