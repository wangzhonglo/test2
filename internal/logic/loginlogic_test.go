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

func TestLoginLogic_Login(t *testing.T) {

	var configFile = flag.String("f", "../../etc/greet-api.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)

	type args struct {
		req *types.STLoginReq
	}
	tests := []struct {
		name    string
		l       *LoginLogic
		args    args
		wantRsp *types.STLoginRsp
		wantErr bool
	}{
		{
			"t",
			&LoginLogic{
				svcCtx: svcCtx,
			},
			args{
				&types.STLoginReq{
					StuId:    "20201108",
					Password: "123456",
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.l.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginLogic.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("LoginLogic.Login() = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
