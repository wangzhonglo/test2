package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StuDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StuDeleteLogic {
	return &StuDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StuDeleteLogic) StuDelete(req *types.STStuDeleteReq) (rsp *types.STStuDeleteRsp, err error) {
	rsp = &types.STStuDeleteRsp{}

	if req.Id == 0 {
		return rsp, errors.New("参数错误")
	}

	query := l.svcCtx.DB.Model(&model.Stu{})
	query = query.Where("id=?", req.Id)
	if err = query.Delete(&model.Stu{}).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	return
}
