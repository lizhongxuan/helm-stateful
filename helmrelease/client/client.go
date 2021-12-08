package client

import (
	"context"
	"errors"
	"fmt"
	log "github.com/golang/glog"
	"helm-statuful/helmrelease/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

func hostURI(group, version, resource, ns, name string) (string, error) {
	groupVersion := ""
	if group != "" {
		groupVersion = fmt.Sprintf("/apis/%s/%s", group, version)
	} else {
		groupVersion = fmt.Sprintf("/api/%s", version)
	}

	if ns != "" {
		ns = fmt.Sprintf("/namespaces/%s", ns)
	}

	if resource == "" {
		return "", errors.New("uri resource not exist")
	}

	if name != "" {
		name = fmt.Sprintf("/%s", name)
	}

	return fmt.Sprintf("%s%s/%s%s", groupVersion, ns, resource, name), nil
}

func GetRESTClient(ctx context.Context) (*rest.RESTClient, error) {
	cfg,err := InclusterRestConfig()
	if err!=nil {
		return nil, err
	}
	return buildRestClient(ctx,cfg,v1.GroupVersion.Group,v1.GroupVersion.Version,v1.Codecs.WithoutConversion())
}

func InclusterRestConfig() (*rest.Config, error) {
	inclusterCfg, err := rest.InClusterConfig()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return inclusterCfg, nil
}

func buildRestClient(ctx context.Context,cfg *rest.Config, group, version string, negoSeri runtime.NegotiatedSerializer) (*rest.RESTClient, error) {
	// 创建restclient
	cfg.GroupVersion = &schema.GroupVersion{
		Group:   group,
		Version: version,
	}
	cfg.APIPath = "/api"
	cfg.ContentType = runtime.ContentTypeJSON
	cfg.NegotiatedSerializer = negoSeri

	restClient, err := rest.RESTClientFor(cfg)
	if err != nil {
		log.Error( err)
		return nil, err
	}
	return restClient, nil
}
