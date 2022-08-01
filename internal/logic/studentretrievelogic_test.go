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

func TestStudentRetrieveLogic_StudentRetrieve(t *testing.T) {

	var configFile = flag.String("f", "../../etc/greet-api.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)

	type args struct {
		req *types.STStudentRetrieveyReq
	}
	tests := []struct {
		name    string
		l       *StudentRetrieveLogic
		args    args
		wantRsp *types.STStudentRetrieveRsp
		wantErr bool
	}{
		{
			"t",
			&StudentRetrieveLogic{
				svcCtx: svcCtx,
			},
			args{
				&types.STStudentRetrieveyReq{
					Id: 2,
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.l.StudentRetrieve(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("StudentRetrieveLogic.StudentRetrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("StudentRetrieveLogic.StudentRetrieve() = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
