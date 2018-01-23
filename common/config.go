package common

type ConfigInfo struct {
	TmpPath			*string
	DownloadPath 	*string
}

var Cfg = new(ConfigInfo)