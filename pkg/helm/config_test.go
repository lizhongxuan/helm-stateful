package helm

import (
	"errors"
	log "github.com/golang/glog"
	"helm-statuful/util"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"reflect"
	"testing"
)

func TestBuild(t *testing.T) {
	chrt,err := getChart("/Users/zhongxuan/Desktop/kdcloud.io/kme/runtime/charts/local-path-provisioner@0.0.20")
	if err !=nil {
		t.Error(err)
		return
	}

	type args struct {
		ns   string
		chrt *chart.Chart
		vals map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{ "xxx1",args{
			"lzx",chrt,nil,
		},"",true},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Build(tt.args.ns, tt.args.chrt, tt.args.vals)
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() got = %v, want %v", got, tt.want)
			}
		})
	}
}

var FileNotExists = errors.New("file is not exists")

// GetChart 获取chart
func getChart(path string) (*chart.Chart, error) {
	if !util.FileExists(path) {
		log.Error(FileNotExists)
		return nil, FileNotExists
	}

	chartData, err := loader.Load(path)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return chartData, nil
}

