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

func TestStuCreatLogic_StuCreat(t *testing.T) {

	var configFile = flag.String("f", "../../etc/greet-api.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	type args struct {
		req *types.STStuCreateReq
	}
	tests := []struct {
		name    string
		l       *StuCreatLogic
		args    args
		wantRsp *types.STStuCreateRsp
		wantErr bool
	}{
		{"t", &StuCreatLogic{svcCtx: svcCtx}, args{&types.STStuCreateReq{Name: "张三", Age: 20, Sex: 1, Step: "Math", StuId: "20201101", Grade: 2020}}, &types.STStuCreateRsp{}, false},
		{"t", &StuCreatLogic{svcCtx: svcCtx}, args{&types.STStuCreateReq{Name: "李四", Age: 20, Sex: 1, Step: "Chinese", StuId: "20201102", Grade: 2020}}, &types.STStuCreateRsp{}, false},
		{"t", &StuCreatLogic{svcCtx: svcCtx}, args{&types.STStuCreateReq{Name: "王五", Age: 21, Sex: 2, Step: "Math", StuId: "20191101", Grade: 2019}}, &types.STStuCreateRsp{}, false},
		{"t", &StuCreatLogic{svcCtx: svcCtx}, args{&types.STStuCreateReq{Name: "赵六", Age: 21, Sex: 2, Step: "Chinese", StuId: "20191102", Grade: 2019}}, &types.STStuCreateRsp{}, false},
		{"t", &StuCreatLogic{svcCtx: svcCtx}, args{&types.STStuCreateReq{Name: "张红", Age: 19, Sex: 2, Step: "Math", StuId: "20211101", Grade: 2021}}, &types.STStuCreateRsp{}, false},
		{"t", &StuCreatLogic{svcCtx: svcCtx}, args{&types.STStuCreateReq{Name: "李明", Age: 20, Sex: 1, Step: "English", StuId: "20201103", Grade: 2020}}, &types.STStuCreateRsp{}, false},
		{"t", &StuCreatLogic{svcCtx: svcCtx}, args{&types.STStuCreateReq{Name: "赵凯", Age: 20, Sex: 1, Step: "Math", StuId: "20201104", Grade: 2020}}, &types.STStuCreateRsp{}, false},
		{"t", &StuCreatLogic{svcCtx: svcCtx}, args{&types.STStuCreateReq{Name: "陈芳", Age: 20, Sex: 2, Step: "English", StuId: "20201105", Grade: 2020}}, &types.STStuCreateRsp{}, false},
		{"t", &StuCreatLogic{svcCtx: svcCtx}, args{&types.STStuCreateReq{Name: "刘强", Age: 19, Sex: 1, Step: "Chinese", StuId: "20211102", Grade: 2021}}, &types.STStuCreateRsp{}, false},
		{"t", &StuCreatLogic{svcCtx: svcCtx}, args{&types.STStuCreateReq{Name: "张铭", Age: 19, Sex: 1, Step: "English", StuId: "20201103", Grade: 2021}}, &types.STStuCreateRsp{}, false},
		{"t", &StuCreatLogic{svcCtx: svcCtx}, args{&types.STStuCreateReq{Name: "吴轩", Age: 21, Sex: 2, Step: "Math", StuId: "20191103", Grade: 2019}}, &types.STStuCreateRsp{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.l.StuCreat(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("StuCreatLogic.StuCreat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("StuCreatLogic.StuCreat() = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
