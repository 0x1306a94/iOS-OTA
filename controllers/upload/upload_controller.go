package upload

import (
	"github.com/astaxie/beego"
	"fmt"
	"iOS-OTA/common"
	"time"
	"os"
	"strings"
	"github.com/astaxie/beego/orm"
	"iOS-OTA/models/ipa"
)

type UploadController struct {
	beego.Controller
}

func (u *UploadController) Post()  {

	fmt.Println("接受到文件上传请求......")
	var (
		ipaPath string
		iconPath string
		folderPath string
		lastBuildId int64 = 1
		version string = u.GetString("version", "")
		fileName string
	)

	if version == "" {
		u.Data["json"] = &map[string]interface{}{"state":0, "msg":"缺少 version 参数"}
		u.ServeJSON()
		return
	}

	f, fh, err := u.GetFile("uploadFile")
	defer f.Close()
	if err != nil {
		fmt.Println("get file error: ", err)
		u.Data["json"] = &map[string]interface{}{"state":0, "msg":"获取文件错误"}
		u.ServeJSON()
		return
	}

	fileName = strings.Replace(fh.Filename, ".ipa", "", -1)

	o := orm.NewOrm()

	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		u.Abort("502")
		return
	}

	qb.Select("name, build_id").
		From("ipa_info").
		Where("name = ?").
		And("version = ?").
		OrderBy("build_id").
		Desc().
		Limit(1)

	sql := qb.String()

	fmt.Println("sql: ", sql)
	tmpObj := new(ipa.IpaInfo)
	o.Raw(sql, fileName, version).
		QueryRow(&tmpObj)

	fmt.Println("tmpObj: ", tmpObj)

	if tmpObj.Name != "" {
		lastBuildId = tmpObj.BuildId + 1
		folderPath = "ipa/" +fileName + "/" + version + "-" + fmt.Sprint(lastBuildId)
	} else {
		lastBuildId = 1
		folderPath = "ipa/" + fileName + "/" + version + "-" + fmt.Sprint(lastBuildId)
	}


	basePath := *common.Cfg.DownloadPath + "/" + folderPath + "/"
	savePath := basePath + fh.Filename

	ipaPath = "/download/" + folderPath + "/" + fh.Filename
	iconPath = "/download/" + folderPath + "/icon.png"

	fmt.Println("folderPath: ", folderPath)
	fmt.Println("basePath: ", basePath)
	fmt.Println("savePath: ", savePath)

	err = os.MkdirAll(basePath, os.ModePerm)
	saveFlag := false

	defer func(b *bool) {
		if !*b {
			fmt.Println("文件保存失败, 清理目录....")
			os.RemoveAll(basePath)
		}
	}(&saveFlag)

	if err != nil {
		saveFlag = false
		u.Abort("502")
		return
	}
	err = u.SaveToFile("uploadFile", savePath)
	if err != nil {
		saveFlag = false
		u.Data["json"] = &map[string]interface{}{"state":0, "msg":"保存文件错误"}
		u.ServeJSON()
		return
	}


	info, err := common.UnpackIpa(savePath, basePath)
	if err != nil {
		saveFlag = false
		u.Data["json"] = &map[string]interface{}{"state":0, "msg":"保存文件错误"}
		u.ServeJSON()
		return
	}

	ipaInfo := new(ipa.IpaInfo)
	ipaInfo.Name = info.Name
	ipaInfo.Identifier = info.Identifier
	ipaInfo.Version = info.Version
	ipaInfo.BuildId = lastBuildId
	ipaInfo.Path = ipaPath
	ipaInfo.IconPath = iconPath
	ipaInfo.CreatedAt = time.Now()
	id, err := o.Insert(ipaInfo)
	if err != nil {
		saveFlag = false
		u.Data["json"] = &map[string]interface{}{"state":0, "msg":"保存文件信息错误"}
		u.ServeJSON()
		return
	}
	fmt.Println("保存数据库成功 id: ", id)

	saveFlag = true
	fmt.Println("info: ", info)
	u.Data["json"] = &map[string]interface{}{
		"state":1,
		"result" : ipaInfo,
	}
	u.ServeJSON()
}