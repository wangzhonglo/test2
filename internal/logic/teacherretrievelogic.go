package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeacherRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeacherRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeacherRetrieveLogic {
	return &TeacherRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeacherRetrieveLogic) TeacherRetrieve(req *types.STTeacherRetrieveReq) (rsp *types.STTeacherRetrieveRsp, err error) {
	rsp = &types.STTeacherRetrieveRsp{}

	if req.Id == 0 {
		return rsp, errors.New("参数错误")
	}

	teacher := &model.Teacher{}
	query := l.svcCtx.DB.Model(&model.Teacher{})
	query = query.Where("id=?", req.Id)
	if err = query.First(&teacher).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	rsp = &types.STTeacherRetrieveRsp{
		Id:        teacher.Id,
		Name:      teacher.Name,
		TeacherId: teacher.TeacherId,
		Step:      teacher.Step,
		Course:    teacher.Course,
		CreatedAt: teacher.CreatedAt.Unix(),
		UpdatedAt: teacher.UpdatedAt.Unix(),
	}
	return
}
