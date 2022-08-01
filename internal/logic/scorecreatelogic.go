package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScoreCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScoreCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScoreCreateLogic {
	return &ScoreCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScoreCreateLogic) ScoreCreate(req *types.STScoreCreateReq) (rsp *types.STScoreCreateRsp, err error) {
	rsp = &types.STScoreCreateRsp{}
	if req.StuId == "" {
		return rsp, errors.New("参数错误")
	}

	score := &model.Score{
		StuId:      req.StuId,
		Chinese:    req.Chinese,
		Math:       req.Math,
		English:    req.English,
		CTeacherId: req.CTeacherId,
		MTeacherId: req.MTeacherId,
		ETeacherId: req.ETeacherId,
		Year:       req.Year,
		Term:       req.Term,
	}
	if err = l.svcCtx.DB.Create(score).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	return rsp, nil
}
