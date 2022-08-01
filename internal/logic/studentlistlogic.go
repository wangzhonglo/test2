package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StudentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StudentListLogic {
	return &StudentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StudentListLogic) StudentList(req *types.STStudentListReq) (rsp *types.STStudentListRsp, err error) {
	rsp = &types.STStudentListRsp{}

	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 20
	}

	query := l.svcCtx.DB.Model(&model.Student{})
	if req.Id > 0 {
		query = query.Where("id=?", req.Id)
	}
	if req.Name != "" {
		query = query.Where("name=?", req.Name)
	}
	if req.Age > 0 {
		query = query.Where("age=?", req.Age)
	}
	if req.Gender > 0 {
		query = query.Where("gender=?", req.Gender)
	}

	total := int64(0)
	if err = query.Count(&total).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}
	rsp.Total = uint64(total)

	query = query.Limit(int(req.Limit)).Offset((int(req.Page-1) * int(req.Limit)))

	data := []*model.Student{}
	if err = query.Find(&data).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	for _, student := range data {
		rsp.Items = append(rsp.Items, types.Student{
			Id:        student.Id,
			Name:      student.Name,
			Age:       student.Age,
			Gender:    student.Gender,
			CreatedAt: student.CreatedAt.Unix(),
			UpdatedAt: student.UpdatedAt.Unix(),
		})
	}

	return rsp, nil
}
