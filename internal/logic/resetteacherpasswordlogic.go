package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ResetTeacherPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetTeacherPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetTeacherPasswordLogic {
	return &ResetTeacherPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetTeacherPasswordLogic) ResetTeacherPassword(req *types.TResetPassworReq) (rsp *types.TResetPassworRsp, err error) {
	rsp = &types.TResetPassworRsp{}

	{ // step 1: valid params
		if req.TeacherId == "" || req.Password == "" || req.RePassword == "" {
			return rsp, errors.New("参数错误")
		}

		if req.Password != req.RePassword {
			return rsp, errors.New("密码错误")
		}
	}

	teacher := &model.Teacher{}
	{ //step 2: find student record
		query := l.svcCtx.DB.Model(&model.Teacher{})
		query = query.Where("teacher_id=?", req.TeacherId)
		if err = query.First(&teacher).Error; err != nil && err != gorm.ErrRecordNotFound {
			return rsp, errors.New("数据库错误")
		}
		if teacher.Id == 0 {
			return rsp, errors.New("差无此学生")
		}
	}

	password := []byte{}
	{ //step 3: create password
		if password, err = bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost); err != nil {
			return rsp, errors.New("系统错误")
		}
	}

	{ //step 4: update student record and save password
		upd := &model.Teacher{
			Password: string(password),
		}
		query := l.svcCtx.DB.Model(&model.Teacher{})
		query = query.Where("id=?", teacher.Id)
		if err = query.Updates(&upd).Error; err != nil {
			return rsp, errors.New("数据库错误")
		}
	}

	return rsp, nil
}
