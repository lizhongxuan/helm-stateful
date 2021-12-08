package helm
//
//import (
//	"context"
//	"encoding/json"
//	log "github.com/golang/glog"
//	"helm.sh/helm/v3/pkg/action"
//	"helm.sh/helm/v3/pkg/chart"
//	"helm.sh/helm/v3/pkg/release"
//	"helm.sh/helm/v3/pkg/repo"
//	"io/ioutil"
//)
//
//type Chart struct{}
//
//var (
//	ChartService Chart
//)
//
//// TODO 有个接口可以查询集群上的各个节点的资源/状态/标签,预计部署的后可能出现的情况: 如部署到哪些标签节点上/因集群驱逐或节点出问题或资源不足而部署失败
//
//// Install 安装chart
//func (*Chart) DryInstall(ctx context.Context, ns string,) (*release.Release,error) {
//
//	actionCfg, errRes := getActionConfig(ctx, zid, ns)
//	if errRes != nil {
//		return errRes
//	}
//
//
//}
//
//func (*Chart) IndexFile(ctx context.Context) (*repo.IndexFile, error) {
//	indexPath, errRes := getIndexFilePath(ctx)
//	if errRes != nil {
//		return nil, errRes
//	}
//	indexFile, err := repo.LoadIndexFile(indexPath)
//	if err != nil {
//		log.Error(ctx, err)
//		return nil, err
//	}
//	return indexFile, nil
//}
//
//func (*Chart) ReadmeFile(ctx context.Context, chartName, version string) (string, error) {
//	readmePath, errRes := getReadmeFilePath(ctx, chartName, version)
//	if errRes != nil {
//		return "", errRes
//	}
//	readmeBytes, err := ioutil.ReadFile(readmePath)
//	if err != nil {
//		log.Error( err, readmePath)
//		return "", err
//	}
//	return string(readmeBytes), nil
//}
//
//// ArgsFile 返回chart的参数,只有index标志了isOtherArg参数为true,前端才会请求该接口
//func (*Chart) ArgsFile(ctx context.Context, chartName, version string) (map[string]interface{}, error) {
//	argPath, errRes := getArgsFilePath(ctx, chartName, version)
//	if errRes != nil {
//		return nil, errRes
//	}
//	argBytes, err := ioutil.ReadFile(argPath)
//	if err != nil {
//		log.Error(ctx, err, argPath)
//		return nil, err
//	}
//	argsJson := make(map[string]interface{})
//	if err := json.Unmarshal(argBytes, &argsJson); err != nil {
//		log.Error(ctx, err, argPath)
//		return nil, err
//	}
//	return argsJson, nil
//}
