package client

import (
	"context"
	log "github.com/golang/glog"
	v1 "helm-statuful/helmrelease/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

// PS.
// HelmRelease operator In https://github.com/lizhongxuan/kubehelm-operator

func List(ctx context.Context,  ns string, opts metav1.ListOptions) (*v1.HelmReleaseList, error) {
	cli,err:= GetRESTClient(ctx)
	if err !=nil {
		return nil, err
	}

	uri, err := hostURI(v1.GroupVersion.Group, v1.GroupVersion.Version, v1.Resource, ns, "")
	if err != nil {
		return nil, err
	}
	result := v1.HelmReleaseList{}

	if err = cli.
		Get().
		RequestURI(uri).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result);err != nil {
		return nil, err
	}

	return &result, nil
}

func Get(ctx context.Context,  ns, name string, opts metav1.GetOptions) (*v1.HelmRelease, error) {
	cli,err:= GetRESTClient(ctx)
	if err !=nil {
		return nil, err
	}


	uri, err := hostURI(v1.GroupVersion.Group, v1.GroupVersion.Version, v1.Resource, ns, name)
	if err != nil {
		return nil, err
	}

	result := v1.HelmRelease{}

	if err = cli.
		Get().
		RequestURI(uri).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result);err != nil {
		return nil, err
	}

	return &result, nil
}

func Create(ctx context.Context, hr *v1.HelmRelease) (*v1.HelmRelease, error) {
	cli,err:= GetRESTClient(ctx)
	if err !=nil {
		return nil, err
	}

	uri, err := hostURI(v1.GroupVersion.Group, v1.GroupVersion.Version, v1.Resource, hr.Namespace, "")
	if err != nil {
		return nil, err
	}

	result := &v1.HelmRelease{}
	if err = cli.
		Post().
		RequestURI(uri).
		Body(hr).
		Do(ctx).
		Into(result);err != nil {
		return nil, err
	}
	return result, nil
}

func Delete(ctx context.Context,  ns, name string) (*v1.HelmRelease, error) {
	cli,err:= GetRESTClient(ctx)
	if err !=nil {
		return nil, err
	}

	uri, err := hostURI(v1.GroupVersion.Group, v1.GroupVersion.Version, v1.Resource, ns, name)
	if err != nil {
		return nil, err
	}

	result := v1.HelmRelease{}
	if err = cli.
		Delete().
		RequestURI(uri).
		Do(ctx).
		Error();err != nil {
		return nil, err
	}
	log.Info(name, "@ HelmRelease delete success.")
	return &result, nil
}

func Update(ctx context.Context, hr *v1.HelmRelease) (*v1.HelmRelease, error) {
	cli,err:= GetRESTClient(ctx)
	if err !=nil {
		return nil, err
	}

	uri, err := hostURI(v1.GroupVersion.Group, v1.GroupVersion.Version, v1.Resource, hr.Namespace, hr.Name)
	if err != nil {
		return nil, err
	}

	result := v1.HelmRelease{}
	err = cli.
		Put().
		RequestURI(uri).
		Body(hr).
		Do(ctx).
		Into(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}