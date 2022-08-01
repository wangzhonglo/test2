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

func TestGradeUpdateLogic_GradeUpdate(t *testing.T) {
	var configFile = flag.String("f", "../../etc/greet-api.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	type args struct {
		req *types.STGradeUpdateReq
	}
	tests := []struct {
		name    string
		l       *GradeUpdateLogic
		args    args
		wantRsp *types.STGradeUpdateRsp
		wantErr bool
	}{
		{
			"t",
			&GradeUpdateLogic{
				svcCtx: svcCtx,
			},
			args{
				&types.STGradeUpdateReq{
					Id:      2,
					UserId:  2,
					Chinese: 99,
					Math:    99,
					English: 99,
					Year:    2020,
					Term:    2,
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.l.GradeUpdate(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GradeUpdateLogic.GradeUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("GradeUpdateLogic.GradeUpdate() = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
