package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScoreRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScoreRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScoreRetrieveLogic {
	return &ScoreRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScoreRetrieveLogic) ScoreRetrieve(req *types.STScoreRetrieveReq) (rsp *types.STScoreRetrieveRsp, err error) {
	rsp = &types.STScoreRetrieveRsp{}

	if req.Id == 0 {
		return rsp, errors.New("参数错误")
	}

	score := &model.Score{}
	query := l.svcCtx.DB.Model(&model.Score{})
	query = query.Where("id=?", req.Id)
	if err = query.First(&score).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}

	rsp = &types.STScoreRetrieveRsp{
		Id:         score.Id,
		StuId:      score.StuId,
		CTeacherId: score.CTeacherId,
		Chinese:    score.Chinese,
		Math:       score.Math,
		MTeacherId: score.MTeacherId,
		English:    score.English,
		Year:       score.Year,
		Term:       score.Term,
		CreatedAt:  score.CreatedAt.Unix(),
		UpdatedAt:  score.UpdatedAt.Unix(),
	}

	return
}
