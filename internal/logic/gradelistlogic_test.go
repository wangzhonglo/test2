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

func TestGradeListLogic_GradeList(t *testing.T) {
	var configFile = flag.String("f", "../../etc/greet-api.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	type args struct {
		req *types.STGradeListReq
	}
	tests := []struct {
		name    string
		l       *GradeListLogic
		args    args
		wantRsp *types.STGradeListRsp
		wantErr bool
	}{
		{
			"t",
			&GradeListLogic{
				svcCtx: svcCtx,
			},
			args{
				&types.STGradeListReq{
					Id: 2,
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.l.GradeList(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GradeListLogic.GradeList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("GradeListLogic.GradeList() = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
