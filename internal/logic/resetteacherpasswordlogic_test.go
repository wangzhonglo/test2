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

func TestResetTeacherPasswordLogic_ResetTeacherPassword(t *testing.T) {

	var configFile = flag.String("f", "../../etc/greet-api.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)

	type args struct {
		req *types.TResetPassworReq
	}
	tests := []struct {
		name    string
		l       *ResetTeacherPasswordLogic
		args    args
		wantRsp *types.TResetPassworRsp
		wantErr bool
	}{
		{
			"t",
			&ResetTeacherPasswordLogic{
				svcCtx: svcCtx,
			},
			args{
				&types.TResetPassworReq{
					TeacherId:  "19780301",
					Password:   "123456",
					RePassword: "123456",
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.l.ResetTeacherPassword(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResetTeacherPasswordLogic.ResetTeacherPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("ResetTeacherPasswordLogic.ResetTeacherPassword() = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
