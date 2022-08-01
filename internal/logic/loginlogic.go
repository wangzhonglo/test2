package logic

import (
	"context"
	"errors"
	"time"

	"greet/internal/model"
	"greet/internal/svc"
	"greet/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.STLoginReq) (rsp *types.STLoginRsp, err error) {
	rsp = &types.STLoginRsp{}

	{ // step 1: valid params
		if req.StuId == "" || req.Password == "" {
			return rsp, errors.New("参数错误")
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
			return rsp, errors.New("用户名或密码错误")
		}
	}

	{ //step 3: compare password
		if err = bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(req.Password)); err != nil {
			return rsp, errors.New("用户名或密码错误")
		}
	}

	{ //step 4: create jwt token
		iat := time.Now().Unix()

		claims := make(jwt.MapClaims)
		claims["exp"] = iat + l.svcCtx.Config.Auth.AccessExpire
		claims["iat"] = iat
		claims["userId"] = student.StuId

		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims = claims

		if rsp.Token, err = token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret)); err != nil {
			return rsp, errors.New("系统错误")
		}
	}

	return rsp, nil
}
