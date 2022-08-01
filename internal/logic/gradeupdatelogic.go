package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GradeUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGradeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GradeUpdateLogic {
	return &GradeUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GradeUpdateLogic) GradeUpdate(req *types.STGradeUpdateReq) (rsp *types.STGradeUpdateRsp, err error) {
	rsp = &types.STGradeUpdateRsp{}
	if req.Id == 0 {
		return rsp, errors.New("参数错误")
	}

	grade := &model.Grade{
		Chinese: req.Chinese,
		Math:    req.Math,
		English: req.English,
		UserId:  req.UserId,
	}
	query := l.svcCtx.DB.Model(&model.Grade{})
	query = query.Where("id=?", req.Id)
	if err = query.Updates(&grade).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	return
}
