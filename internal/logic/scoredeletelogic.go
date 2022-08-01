package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScoreDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScoreDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScoreDeleteLogic {
	return &ScoreDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScoreDeleteLogic) ScoreDelete(req *types.STScoreDeleteReq) (rsp *types.STScoreDeleteRsp, err error) {
	rsp = &types.STScoreDeleteRsp{}
	if req.Id == 0 {
		return rsp, errors.New("参数错误")
	}

	query := l.svcCtx.DB.Model(&model.Score{})
	query = query.Where("id=?", req.Id)
	if err = query.Delete(&model.Score{}).Error; err != nil {
		return rsp, errors.New("数据库错误")
	}
	return
}
