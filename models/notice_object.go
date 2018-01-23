package models

import (
	_ "github.com/astaxie/beego/orm"
	"time"
)

type NoticeObject struct {
	Id 				int64 `json:"id,omitempty" orm:"column(id)"`
	Name 			string `json:"name" orm:"column(name);size(50)"`
	Platform 		string `json:"platform" orm:"column(platform);size(50)"`
	Version 		string `json:"version" orm:"column(version);size(50)"`
	Branch  		string `json:"branch" orm:"column(branch);size(50)"`
	Author			string `json:"author" orm:"column(author);size(50)"`
	AuthorMail		string `json:"author_mail" orm:"column(author_mail);size(100)"`
	Committer		string `json:"committer" orm:"column(committer);size(50)"`
	CommitterMail	string `json:"committer_mail" orm:"column(committer_mail);size(100)"`
	CommitId		string `json:"commit_id" orm:"column(commit_id);size(255)"`
	CommitContent 	string `json:"commit_content" orm:"column(commit_content);size(10000)"`
	Code			string `json:"code" orm:"column(code);size(500)"`
	Created 		time.Time `json:"created,omitempty" orm:"column(created);auto_now_add;type(date)"`
}

func (u *NoticeObject) TableName() string {
	return "package_notice_tab"
}