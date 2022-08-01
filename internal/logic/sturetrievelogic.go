package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StuRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStuRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StuRetrieveLogic {
	return &StuRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StuRetrieveLogic) StuRetrieve(req *types.STStuRetrieveReq) (rsp *types.STStuRetrieveRsp, err error) {
	rsp = &types.STStuRetrieveRsp{}

	if req.Id == 0 {
		return rsp, errors.New("参数错误")
	}

	stu := &model.Stu{}
	query := l.svcCtx.DB.Model(&model.Stu{})
	query = query.Where("id=?", req.Id)
	if err = query.First(&stu).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	rsp = &types.STStuRetrieveRsp{
		Id:        stu.Id,
		Name:      stu.Name,
		Age:       stu.Age,
		Sex:       stu.Sex,
		StuId:     stu.StuId,
		Step:      stu.Step,
		Grade:     stu.Grade,
		CreatedAt: stu.CreatedAt.Unix(),
		UpdatedAt: stu.UpdatedAt.Unix(),
	}
	return
}
