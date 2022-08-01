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

func TestStudentListLogic_StudentList(t *testing.T) {

	var configFile = flag.String("f", "../../etc/greet-api.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)

	type args struct {
		req *types.STStudentListReq
	}
	tests := []struct {
		name    string
		l       *StudentListLogic
		args    args
		wantRsp *types.STStudentListRsp
		wantErr bool
	}{
		{
			"t",
			&StudentListLogic{
				svcCtx: svcCtx,
			},
			args{
				&types.STStudentListReq{
					Id:     2,
					Name:   "赵六",
					Age:    20,
					Gender: 1,
				},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.l.StudentList(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("StudentListLogic.StudentList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("StudentListLogic.StudentList() = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
