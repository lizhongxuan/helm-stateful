package helm

import (
	"fmt"
	log "github.com/golang/glog"
	"helm-statuful/helmrelease/client"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)


func Build(ns string,chrt *chart.Chart, vals map[string]interface{})(string,error){
	acfg,err := getActionConfig(ns,"kubehelm")
	if err != nil {
		log.Error(err)
		return "",err
	}

	install:=action.NewInstall(acfg)
	install.DryRun = true
	rls,err:=install.Run(chrt, vals)
	if err != nil {
		log.Error(err)
		return "",err
	}
	return rls.Manifest,nil
}


//获取helm3配置
func getActionConfig(ns,name string) (*action.Configuration, error) {
	cfg, err := client.InclusterRestConfig()
	if err != nil {
		log.Error( err)
		return nil, err
	}

	config := &clientcmdapi.Config{
		Kind:           "Config",
		APIVersion:     "v1",
		CurrentContext: name,
		Preferences:    *clientcmdapi.NewPreferences(),
		Contexts: map[string]*clientcmdapi.Context{
			name: {
				Cluster:  name,
				AuthInfo: name,
			},
		},
		Clusters: map[string]*clientcmdapi.Cluster{
			name: {
				Server:                cfg.Host,
				InsecureSkipTLSVerify: true,
			},
		},
		AuthInfos: map[string]*clientcmdapi.AuthInfo{
			name:{
				Token:     cfg.BearerToken,
				TokenFile: cfg.BearerTokenFile,
			},
		},
	}

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(NewGetter(clientcmd.NewDefaultClientConfig(*config, nil)), ns, "", logf); err != nil {
		log.Error(err, ns)
		return nil, err
	}
	return actionConfig, nil
}

func logf(format string, v ...interface{}) {
		if len(v) > 0 {
			log.Info(fmt.Sprintf(format, v...))
		} else {
			log.Info(format)
		}
}
