package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StuUpdateLogic {
	return &StuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StuUpdateLogic) StuUpdate(req *types.STStuUpdateReq) (rsp *types.STStuUpdateRsp, err error) {
	rsp = &types.STStuUpdateRsp{}

	if req.Id == 0 {
		return rsp, errors.New("参数错误")
	}

	stu := &model.Stu{
		Name:  req.Name,
		Age:   req.Age,
		Sex:   req.Sex,
		Grade: req.Grade,
		StuId: req.StuId,
		Step:  req.Step,
	}
	query := l.svcCtx.DB.Model(&model.Stu{})
	query = query.Where("id=?", req.Id)
	if err = query.Updates(&stu).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	return
}
