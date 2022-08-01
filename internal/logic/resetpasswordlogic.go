package logic

import (
	"context"
	"errors"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"golang.org/x/crypto/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ResetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordLogic) ResetPassword(req *types.STResetPassworReq) (rsp *types.STResetPassworRsp, err error) {
	rsp = &types.STResetPassworRsp{}

	{ // step 1: valid params
		if req.StuId == "" || req.Password == "" || req.RePassword == "" {
			return rsp, errors.New("参数错误")
		}

		if req.Password != req.RePassword {
			return rsp, errors.New("密码错误")
		}
	}

	student := &model.Stu{}
	{ //step 2: find student record
		query := l.svcCtx.DB.Model(&model.Stu{})
		query = query.Where("stu_id=?", req.StuId)
		if err = query.First(&student).Error; err != nil && err != gorm.ErrRecordNotFound {
			return rsp, errors.New("数据库错误")
		}
		if student.Id == 0 {
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
		upd := &model.Stu{
			Password: string(password),
		}
		query := l.svcCtx.DB.Model(&model.Stu{})
		query = query.Where("id=?", student.Id)
		if err = query.Updates(&upd).Error; err != nil {
			return rsp, errors.New("数据库错误")
		}
	}

	return rsp, nil
}
