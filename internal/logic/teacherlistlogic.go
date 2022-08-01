package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeacherListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeacherListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeacherListLogic {
	return &TeacherListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeacherListLogic) TeacherList(req *types.STTeacherListReq) (rsp *types.STTeacherListRsp, err error) {
	rsp = &types.STTeacherListRsp{}

	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 20
	}

	query := l.svcCtx.DB.Model(&model.Teacher{})
	if req.Id > 0 {
		query = query.Where("id=?", req.Id)
	}
	if req.Name != "" {
		query = query.Where("name=?", req.Name)
	}
	if req.Course != "" {
		query = query.Where("course=?", req.Course)
	}
	if req.TeacherId != "" {
		query = query.Where("teacherId=?", req.TeacherId)
	}

	total := int64(0)
	if err = query.Count(&total).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}
	rsp.Total = uint64(total)

	query = query.Limit(int(req.Limit)).Offset((int(req.Page-1) * int(req.Limit)))

	data := []*model.Teacher{}
	if err = query.Find(&data).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	for _, teacher := range data {
		rsp.Items = append(rsp.Items, types.Teacher{
			Id:        teacher.Id,
			Name:      teacher.Name,
			Step:      teacher.Step,
			TeacherId: teacher.TeacherId,
			CreatedAt: teacher.CreatedAt.Unix(),
			UpdatedAt: teacher.UpdatedAt.Unix(),
		})
	}

	return rsp, nil
}
