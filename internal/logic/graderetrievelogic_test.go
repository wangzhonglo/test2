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

func TestGradeRetrieveLogic_GradeRetrieve(t *testing.T) {
	var configFile = flag.String("f", "../../etc/greet-api.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	type args struct {
		req *types.STGradeRetrieveyReq
	}
	tests := []struct {
		name    string
		l       *GradeRetrieveLogic
		args    args
		wantRsp *types.STGradeRetrieveRsp
		wantErr bool
	}{
		{
			"t",
			&GradeRetrieveLogic{
				svcCtx: svcCtx,
			},
			args{
				&types.STGradeRetrieveyReq{
					Id: 2,
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.l.GradeRetrieve(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GradeRetrieveLogic.GradeRetrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("GradeRetrieveLogic.GradeRetrieve() = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
