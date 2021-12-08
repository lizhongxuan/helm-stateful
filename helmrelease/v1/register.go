package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

const (
	Kind       = "HelmRelease"
	APIVersion = "xuan.com/v1"
	Resource   = "helmreleases"
)

var (
	Scheme         = runtime.NewScheme()
	Codecs         = serializer.NewCodecFactory(Scheme)
	ParameterCodec = runtime.NewParameterCodec(Scheme)

)

func init() {
	utilruntime.Must(SchemeBuilder.AddToScheme(Scheme))
}