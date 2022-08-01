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

func TestGradeDeleteLogic_GradeDelete(t *testing.T) {
	var configFile = flag.String("f", "../../etc/greet-api.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	type args struct {
		req *types.STGradeDeleteReq
	}
	tests := []struct {
		name    string
		l       *GradeDeleteLogic
		args    args
		wantRsp *types.STGradeDeleteRsp
		wantErr bool
	}{
		{
			"t",
			&GradeDeleteLogic{
				svcCtx: svcCtx,
			},
			args{
				&types.STGradeDeleteReq{
					Id: 1,
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.l.GradeDelete(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GradeDeleteLogic.GradeDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("GradeDeleteLogic.GradeDelete() = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
