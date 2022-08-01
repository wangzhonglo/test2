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

func TestStudentDeleteLogic_StudentDelete(t *testing.T) {
	var configFile = flag.String("f", "../../etc/greet-api.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)

	type args struct {
		req *types.STStudentDeleteReq
	}
	tests := []struct {
		name    string
		l       *StudentDeleteLogic
		args    args
		wantRsp *types.STStudentDeleteRsp
		wantErr bool
	}{
		{
			"t",
			&StudentDeleteLogic{
				svcCtx: svcCtx,
			},
			args{
				&types.STStudentDeleteReq{
					Id: 1,
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.l.StudentDelete(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("StudentDeleteLogic.StudentDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("StudentDeleteLogic.StudentDelete() = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
