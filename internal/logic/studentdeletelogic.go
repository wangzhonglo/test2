package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StudentDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StudentDeleteLogic {
	return &StudentDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StudentDeleteLogic) StudentDelete(req *types.STStudentDeleteReq) (rsp *types.STStudentDeleteRsp, err error) {
	rsp = &types.STStudentDeleteRsp{}

	if req.Id == 0 {
		return rsp, errors.New("参数错误")
	}

	query := l.svcCtx.DB.Model(&model.Student{})
	query = query.Where("id=?", req.Id)
	if err = query.Delete(&model.Student{}).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	return
}
