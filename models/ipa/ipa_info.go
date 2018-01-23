package ipa

import (
	_ "github.com/astaxie/beego/orm"
	"time"
)
type IpaInfo struct {
	Id 			int64 `json:"id,omitempty" orm:"column(id)"`
	Name 		string `json:"name" orm:"column(name);size(50)"`
	Identifier  string `json:"identifier" orm:"column(identifier);size(100)""`
	Version 	string `json:"version" orm:"column(version);size(50)"`
	BuildId		int64 `json:"build_id" orm:"column(build_id)"`
	Path 		string `json:"path" orm:"column(path);size(5000)"`
	IconPath	string `json:"icon_path" orm:"column(icon_path);size(5000)"`
	CreatedAt	time.Time `json:"created_at,omitempty" orm:"column(created_at);auto_now_add;type(date)"`
}