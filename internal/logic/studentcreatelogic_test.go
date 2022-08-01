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

func TestStudentCreateLogic_StudentCreate(t *testing.T) {

	var configFile = flag.String("f", "../../etc/greet-api.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)

	type args struct {
		req *types.STStudentCreateReq
	}
	tests := []struct {
		name    string
		l       *StudentCreateLogic
		args    args
		wantRsp *types.STStudentCreateRsp
		wantErr bool
	}{
		// {"t", &StudentCreateLogic{svcCtx: svcCtx}, args{&types.STStudentCreateReq{Name: "赵六3", Age: 21, Gender: 1}}, nil, false},
		// {"t", &StudentCreateLogic{svcCtx: svcCtx}, args{&types.STStudentCreateReq{Name: "赵六4", Age: 21, Gender: 1}}, nil, false},
		// {"t", &StudentCreateLogic{svcCtx: svcCtx}, args{&types.STStudentCreateReq{Name: "赵六5", Age: 21, Gender: 1}}, nil, false},
		// {"t", &StudentCreateLogic{svcCtx: svcCtx}, args{&types.STStudentCreateReq{Name: "赵六6", Age: 21, Gender: 1}}, nil, false},
		// {"t", &StudentCreateLogic{svcCtx: svcCtx}, args{&types.STStudentCreateReq{Name: "赵六7", Age: 21, Gender: 1}}, nil, false},
		{"t", &StudentCreateLogic{svcCtx: svcCtx}, args{&types.STStudentCreateReq{Name: "赵六8", Age: 21, Gender: 1}}, &types.STStudentCreateRsp{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.l.StudentCreate(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("StudentCreateLogic.StudentCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("StudentCreateLogic.StudentCreate() = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
