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

func TestResetPasswordLogic_ResetPassword(t *testing.T) {

	var configFile = flag.String("f", "../../etc/greet-api.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)

	type args struct {
		req *types.STResetPassworReq
	}
	tests := []struct {
		name    string
		l       *ResetPasswordLogic
		args    args
		wantRsp *types.STResetPassworRsp
		wantErr bool
	}{
		{
			"t",
			&ResetPasswordLogic{
				svcCtx: svcCtx,
			},
			args{
				&types.STResetPassworReq{
					StuId:      "20201108",
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
			gotRsp, err := tt.l.ResetPassword(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResetPasswordLogic.ResetPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("ResetPasswordLogic.ResetPassword() = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
