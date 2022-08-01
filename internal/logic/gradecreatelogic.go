package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GradeCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGradeCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GradeCreateLogic {
	return &GradeCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GradeCreateLogic) GradeCreate(req *types.STGradeCreateReq) (rsp *types.STGradeCreateRsp, err error) {

	rsp = &types.STGradeCreateRsp{}
	if req.UserId == 0 {
		return rsp, errors.New("参数错误")
	}

	grade := &model.Grade{
		UserId:  req.UserId,
		Chinese: req.Chinese,
		Math:    req.Math,
		English: req.English,
		Year:    req.Year,
		Term:    req.Term,
	}
	if err = l.svcCtx.DB.Create(grade).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	return rsp, nil
}
