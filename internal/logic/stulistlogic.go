package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StuListLogic {
	return &StuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StuListLogic) StuList(req *types.STStuListReq) (rsp *types.STStuListRsp, err error) {
	rsp = &types.STStuListRsp{}

	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 20
	}

	query := l.svcCtx.DB.Model(&model.Stu{})
	if req.Id > 0 {
		query = query.Where("id=?", req.Id)
	}
	if req.Name != "" {
		query = query.Where("name=?", req.Name)
	}
	if req.Age > 0 {
		query = query.Where("age=?", req.Age)
	}
	if req.Sex > 0 {
		query = query.Where("sex=?", req.Sex)
	}
	if req.StuId != "" {
		query = query.Where("stuId=?", req.StuId)
	}
	total := int64(0)
	if err = query.Count(&total).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}
	rsp.Total = uint64(total)

	query = query.Limit(int(req.Limit)).Offset((int(req.Page-1) * int(req.Limit)))

	data := []*model.Stu{}
	if err = query.Find(&data).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	for _, stu := range data {
		rsp.Items = append(rsp.Items, types.Stu{
			Id:        stu.Id,
			Name:      stu.Name,
			Age:       stu.Age,
			Sex:       stu.Sex,
			StuId:     stu.StuId,
			Step:      stu.Step,
			Grade:     req.Grade,
			CreatedAt: stu.CreatedAt.Unix(),
			UpdatedAt: stu.UpdatedAt.Unix(),
		})
	}

	return rsp, nil
}
