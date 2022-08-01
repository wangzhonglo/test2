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

func TestScoreCreateLogic_ScoreCreate(t *testing.T) {
	var configFile = flag.String("f", "../../etc/greet-api.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	type args struct {
		req *types.STScoreCreateReq
	}
	tests := []struct {
		name    string
		l       *ScoreCreateLogic
		args    args
		wantRsp *types.STScoreCreateRsp
		wantErr bool
	}{
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201101", Year: 2020, Term: 1, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201101", Year: 2020, Term: 2, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19880101", English: 95, ETeacherId: "19780301"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201101", Year: 2021, Term: 1, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201101", Year: 2021, Term: 2, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19880101", English: 95, ETeacherId: "19780301"}}, &types.STScoreCreateRsp{}, false},

		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201102", Year: 2020, Term: 1, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201102", Year: 2020, Term: 2, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201102", Year: 2021, Term: 1, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19880101", English: 95, ETeacherId: "19780301"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201102", Year: 2021, Term: 2, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},

		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201103", Year: 2020, Term: 1, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201103", Year: 2020, Term: 2, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19880101", English: 95, ETeacherId: "19780301"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201103", Year: 2021, Term: 1, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201103", Year: 2021, Term: 2, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},

		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201104", Year: 2020, Term: 1, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19880101", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201104", Year: 2020, Term: 2, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201104", Year: 2021, Term: 1, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19880101", English: 95, ETeacherId: "19780301"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201104", Year: 2021, Term: 2, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},

		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201105", Year: 2020, Term: 1, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201105", Year: 2020, Term: 2, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19880101", English: 95, ETeacherId: "19780301"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201105", Year: 2021, Term: 1, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20201105", Year: 2021, Term: 2, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},

		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191101", Year: 2019, Term: 1, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191101", Year: 2019, Term: 2, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191101", Year: 2020, Term: 1, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19880101", English: 95, ETeacherId: "19780301"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191101", Year: 2020, Term: 2, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191101", Year: 2021, Term: 1, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191101", Year: 2021, Term: 2, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},

		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191102", Year: 2019, Term: 1, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191102", Year: 2019, Term: 2, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19780301"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191102", Year: 2020, Term: 1, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19880101", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191102", Year: 2020, Term: 2, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19780301"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191102", Year: 2021, Term: 1, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191102", Year: 2021, Term: 2, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19880101", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},

		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191103", Year: 2019, Term: 1, Chinese: 95, CTeacherId: "19800203", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191103", Year: 2019, Term: 2, Chinese: 95, CTeacherId: "19800203", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191103", Year: 2020, Term: 1, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19880101", English: 95, ETeacherId: "19780301"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191103", Year: 2020, Term: 2, Chinese: 95, CTeacherId: "19800203", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191103", Year: 2021, Term: 1, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19880101", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20191103", Year: 2021, Term: 2, Chinese: 95, CTeacherId: "19800203", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19780301"}}, &types.STScoreCreateRsp{}, false},

		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20211101", Year: 2021, Term: 1, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20211101", Year: 2021, Term: 2, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19880101", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20211102", Year: 2021, Term: 1, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19780301"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20211102", Year: 2021, Term: 2, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19880101", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20211103", Year: 2021, Term: 1, Chinese: 95, CTeacherId: "19800202", Math: 95, MTeacherId: "19770102", English: 95, ETeacherId: "19780301"}}, &types.STScoreCreateRsp{}, false},
		{"t", &ScoreCreateLogic{svcCtx: svcCtx}, args{&types.STScoreCreateReq{StuId: "20211103", Year: 2021, Term: 2, Chinese: 95, CTeacherId: "19850201", Math: 95, MTeacherId: "19880101", English: 95, ETeacherId: "19850302"}}, &types.STScoreCreateRsp{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRsp, err := tt.l.ScoreCreate(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScoreCreateLogic.ScoreCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
				t.Errorf("ScoreCreateLogic.ScoreCreate() = %v, want %v", gotRsp, tt.wantRsp)
			}
		})
	}
}
