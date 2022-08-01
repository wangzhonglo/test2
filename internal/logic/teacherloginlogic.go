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
)

type TeacherLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeacherLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeacherLoginLogic {
	return &TeacherLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeacherLoginLogic) TeacherLogin(req *types.TLoginReq) (rsp *types.TLoginRsp, err error) {

	rsp = &types.TLoginRsp{}

	{
		if req.TeacherId == "" || req.Password == "" {
			return rsp, errors.New("参数错误")
		}
	}
	teacher := &model.Teacher{}
	{
		query := l.svcCtx.DB.Model(&model.Teacher{})
		query = query.Where("teacher_id=?", req.TeacherId)
		if err = query.First(&teacher).Error; err != nil {
			return rsp, errors.New("数据库错误")
		}
		if teacher.Id == 0 {
			return rsp, errors.New("用户名或密码错误")
		}

	}

	{
		if err = bcrypt.CompareHashAndPassword([]byte(teacher.Password), []byte(req.Password)); err != nil {
			return rsp, errors.New("用户名或密码错误")
		}
	}
	{
		iat := time.Now().Unix()

		claims := make(jwt.MapClaims)
		claims["exp"] = iat + l.svcCtx.Config.Auth.AccessExpire
		claims["iat"] = iat
		claims["userId"] = teacher.TeacherId
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims = claims

		if rsp.Token, err = token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret)); err != nil {
			return rsp, errors.New("系统错误")
		}

	}
	return
}
