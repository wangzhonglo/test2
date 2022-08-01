package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StudentRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudentRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StudentRetrieveLogic {
	return &StudentRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StudentRetrieveLogic) StudentRetrieve(req *types.STStudentRetrieveyReq) (rsp *types.STStudentRetrieveRsp, err error) {
	rsp = &types.STStudentRetrieveRsp{}

	if req.Id == 0 {
		return rsp, errors.New("参数错误")
	}

	student := &model.Student{}
	query := l.svcCtx.DB.Model(&model.Student{})
	query = query.Where("id=?", req.Id)
	if err = query.First(&student).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	rsp = &types.STStudentRetrieveRsp{
		Id:        student.Id,
		Name:      student.Name,
		Age:       student.Age,
		Gender:    student.Gender,
		CreatedAt: student.CreatedAt.Unix(),
		UpdatedAt: student.UpdatedAt.Unix(),
	}
	return
}
