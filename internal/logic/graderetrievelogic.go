package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GradeRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGradeRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GradeRetrieveLogic {
	return &GradeRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GradeRetrieveLogic) GradeRetrieve(req *types.STGradeRetrieveyReq) (rsp *types.STGradeRetrieveRsp, err error) {
	rsp = &types.STGradeRetrieveRsp{}

	if req.Id == 0 {
		return rsp, errors.New("参数错误")
	}

	grade := &model.Grade{}
	query := l.svcCtx.DB.Model(&model.Grade{})
	query = query.Where("id=?", req.Id)
	if err = query.First(&grade).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	rsp = &types.STGradeRetrieveRsp{
		Id:        grade.Id,
		UserId:    grade.UserId,
		Chinese:   grade.Chinese,
		Math:      grade.Math,
		English:   grade.English,
		CreatedAt: grade.CreatedAt.Unix(),
		UpdatedAt: grade.UpdatedAt.Unix(),
	}

	return
}
