package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScoreUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScoreUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScoreUpdateLogic {
	return &ScoreUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScoreUpdateLogic) ScoreUpdate(req *types.STScoreUpdateReq) (rsp *types.STScoreUpdateRsp, err error) {
	rsp = &types.STScoreUpdateRsp{}
	if req.Id == 0 {
		return rsp, errors.New("参数错误")
	}

	score := &model.Score{
		Chinese:    req.Chinese,
		CTeacherId: req.CTeacherId,
		Math:       req.Math,
		MTeacherId: req.MTeacherId,
		English:    req.English,
		ETeacherId: req.ETeacherId,
		StuId:      req.StuId,
	}
	query := l.svcCtx.DB.Model(&model.Score{})
	query = query.Where("id=?", req.Id)
	if err = query.Updates(&score).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	return
}
