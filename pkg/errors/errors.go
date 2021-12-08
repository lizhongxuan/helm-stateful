package errors

import "errors"

var (
	HELM_INDEX_FILE_LOAD_ERROR       = errors.New("index file load error")
	HELM_CHART_NOT_FOUNT_ERROR             = errors.New("chart not found")
	HELM_CHART_INSTALL_ERROR         = errors.New( "chart install error")
	HELM_CHART_CANNOT_REPEAT_INSTALL_ERROR = errors.New("chart cannot repeat install")
	HELM_CHART_PATH_IS_NULL_ERROR         = errors.New( "chart path is null")


	HELM_RELEASE_NAME_EXIST_ERROR        = errors.New( "release name already exists")
	HELM_RELEASE_NOT_FOUNT_ERROR         = errors.New("release not found")
	HELM_RELEASE_UNINSTALL_ERROR     = errors.New("uninstall release  error")
	HELM_CHART_LOAD_ERROR            = errors.New( "chart load error")
	HELM_RELEASE_LIST_ERROR          = errors.New("list release error")
	HELM_RELEASE_GET_ERROR           = errors.New( "get release error")
	HELM_RELEASE_ARGS_GET_ERROR      = errors.New("get release args error")
	HELM_RELEASE_UPGRADE_ERROR       = errors.New( "upgrade release error")

	HELM_INDEX_NOT_FOUNT_ERROR             = errors.New("chart index not fount")

	DECODE_ERROR   = errors.New("Data format not correct, decode fail") //解码失败
	ENCODE_ERROR   = errors.New("Data format not correct, encode fail")//编码失败

)

