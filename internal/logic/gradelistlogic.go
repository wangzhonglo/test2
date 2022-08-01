package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GradeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGradeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GradeListLogic {
	return &GradeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GradeListLogic) GradeList(req *types.STGradeListReq) (rsp *types.STGradeListRsp, err error) {
	rsp = &types.STGradeListRsp{}

	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 20
	}

	query := l.svcCtx.DB.Model(&model.Grade{})
	if req.Id > 0 {
		query = query.Where("id=?", req.Id)
	}
	if req.UserId != 0 {
		query = query.Where("user=?", req.UserId)
	}
	if req.Chinese > 0 {
		query = query.Where("chinese=?", req.Chinese)
	}
	if req.Math > 0 {
		query = query.Where("math=?", req.Math)
	}
	if req.English > 0 {
		query = query.Where("English=?", req.English)
	}
	total := int64(0)
	if err = query.Count(&total).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}
	rsp.Total = uint64(total)

	query = query.Limit(int(req.Limit)).Offset((int(req.Page-1) * int(req.Limit)))

	data := []*model.Grade{}
	if err = query.Find(&data).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	for _, grade := range data {
		rsp.Items = append(rsp.Items, types.Grade{
			Id:        grade.Id,
			UserId:    grade.UserId,
			Chinese:   grade.Chinese,
			Math:      grade.Math,
			English:   grade.English,
			CreatedAt: grade.CreatedAt.Unix(),
			UpdatedAt: grade.UpdatedAt.Unix(),
		})
	}

	return rsp, nil

}
