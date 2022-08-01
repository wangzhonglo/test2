package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeacherDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeacherDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeacherDeleteLogic {
	return &TeacherDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeacherDeleteLogic) TeacherDelete(req *types.STTeacherDeleteReq) (rsp *types.STTeacherDeleteRsp, err error) {
	rsp = &types.STTeacherDeleteRsp{}

	if req.Id == 0 {
		return rsp, errors.New("参数错误")
	}

	query := l.svcCtx.DB.Model(&model.Teacher{})
	query = query.Where("id=?", req.Id)
	if err = query.Delete(&model.Teacher{}).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	return
}
