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

func TestStudentUpdateLogic_StudentUpdate(t *testing.T) {

	var configFile = flag.String("f", "../../etc/greet-api.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)

	type args struct {
		req *types.STStudentUpdateReq
	}
	tests := []struct {
		name    string
		l       *StudentUpdateLogic
		args    args
		wantRsp *types.STStudentUpdateRsp
		wantErr bool
	}{
		{
			"t",
			&StudentUpdateLogic{
				svcCtx: svcCtx,
			},
			args{
				&types.STStudentUpdateReq{
					Id:     2,
					Name:   "赵六1",
					Age:    21,
					Gender: 2,
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.l.StudentUpdate(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("StudentUpdateLogic.StudentUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("StudentUpdateLogic.StudentUpdate() = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
