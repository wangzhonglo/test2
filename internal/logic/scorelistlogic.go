package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScoreListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScoreListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScoreListLogic {
	return &ScoreListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScoreListLogic) ScoreList(req *types.STScoreListReq) (rsp *types.STScoreListRsp, err error) {
	rsp = &types.STScoreListRsp{}

	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 20
	}

	query := l.svcCtx.DB.Model(&model.Score{})
	if req.Id > 0 {
		query = query.Where("id=?", req.Id)
	}
	if req.StuId != "" {
		query = query.Where("user=?", req.StuId)
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

	data := []*model.Score{}
	if err = query.Find(&data).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	for _, grade := range data {
		rsp.Items = append(rsp.Items, types.Score{
			Id:        grade.Id,
			StuId:     grade.StuId,
			Chinese:   grade.Chinese,
			Math:      grade.Math,
			English:   grade.English,
			CreatedAt: grade.CreatedAt.Unix(),
			UpdatedAt: grade.UpdatedAt.Unix(),
		})
	}

	return rsp, nil
}
