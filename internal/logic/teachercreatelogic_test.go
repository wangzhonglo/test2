package logic

import (
	"flag"
	"greet/internal/config"
	"greet/internal/svc"
	"greet/internal/types"
	"reflect"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
)

func TestTeacherCreateLogic_TeacherCreate(t *testing.T) {
	var configFile = flag.String("f", "../../etc/greet-api.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	type args struct {
		req *types.STTeacherCreateReq
	}
	tests := []struct {
		name    string
		l       *TeacherCreateLogic
		args    args
		wantRsp *types.STTeacherCreateRsp
		wantErr bool
	}{
		{"t", &TeacherCreateLogic{svcCtx: svcCtx}, args{&types.STTeacherCreateReq{Name: "李红", TeacherId: "19880101", Step: "Math", Course: "Math"}}, &types.STTeacherCreateRsp{}, false},
		{"t", &TeacherCreateLogic{svcCtx: svcCtx}, args{&types.STTeacherCreateReq{Name: "王兰", TeacherId: "19850201", Step: "Chinese", Course: "Chinese"}}, &types.STTeacherCreateRsp{}, false},
		{"t", &TeacherCreateLogic{svcCtx: svcCtx}, args{&types.STTeacherCreateReq{Name: "赵丽艳", TeacherId: "19780301", Step: "English", Course: "English"}}, &types.STTeacherCreateRsp{}, false},
		{"t", &TeacherCreateLogic{svcCtx: svcCtx}, args{&types.STTeacherCreateReq{Name: "黄亮", TeacherId: "19800202", Step: "Chinese", Course: "Chinese"}}, &types.STTeacherCreateRsp{}, false},
		{"t", &TeacherCreateLogic{svcCtx: svcCtx}, args{&types.STTeacherCreateReq{Name: "赵煜辉", TeacherId: "19770102", Step: "Math", Course: "Math"}}, &types.STTeacherCreateRsp{}, false},
		{"t", &TeacherCreateLogic{svcCtx: svcCtx}, args{&types.STTeacherCreateReq{Name: "王美美", TeacherId: "19850302", Step: "English", Course: "English"}}, &types.STTeacherCreateRsp{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.l.TeacherCreate(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TeacherCreateLogic.TeacherCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("TeacherCreateLogic.TeacherCreate() = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
