package common

import (
	"iOS-OTA/models/ipa"
	"archive/zip"
	"errors"
	"os"
	"fmt"
	"io"
	"io/ioutil"
	"howett.net/plist"
	"strings"
)
func UnpackIpa(path, saveIconPath string) (*ipa.IpaPlistInfo, error) {

	if path == "" || !PathExist(path) {
		return nil, errors.New("path: " + path + " file not exist")
	}

	zipFile, err := zip.OpenReader(path)
	defer zipFile.Close()
	if err != nil {
		return nil, err
	}

	var (
		info ipa.IpaPlistInfo
		iconSaveFlag = false
	)
	for _, k := range zipFile.Reader.File {
		if strings.HasSuffix(k.Name, "app/Info.plist") {
			r, err := k.Open()
			if err != nil {
				fmt.Println(err)
				continue
			}
			data, err := ioutil.ReadAll(r)
			r.Close()
			if err != nil {
				fmt.Println(err)
				continue
			}
			_, err = plist.Unmarshal(data, &info)
		}

		// 提取AppIcon
		if strings.Contains(k.Name, "app/AppIcon60x60@3x.png")  {

			r, err := k.Open()
			if err != nil {
				fmt.Println(err)
				continue
			}

			newFile, err := os.Create(saveIconPath + "icon.png")
			if err != nil {
				fmt.Println(err)
				r.Close()
				continue
			}

			io.Copy(newFile, r)
			newFile.Close()
			r.Close()
			iconSaveFlag = true
		}
	}

	if info.Identifier != "" && iconSaveFlag {
		return &info, nil
	}
	return nil, nil
}

/*
	解析ipa 中的info.plist
*/
func UnmarshalPlist(path string) (*ipa.IpaPlistInfo, error)  {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	var info ipa.IpaPlistInfo
	_, err = plist.Unmarshal(data, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil

}