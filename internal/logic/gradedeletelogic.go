package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GradeDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGradeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GradeDeleteLogic {
	return &GradeDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GradeDeleteLogic) GradeDelete(req *types.STGradeDeleteReq) (rsp *types.STGradeDeleteRsp, err error) {

	if req.Id == 0 {
		return rsp, errors.New("参数错误")
	}

	query := l.svcCtx.DB.Model(&model.Grade{})
	query = query.Where("id=?", req.Id)
	if err = query.Delete(&model.Grade{}).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}
	return
}
