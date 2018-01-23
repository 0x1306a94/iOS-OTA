package common

import (
	"os"
)

var (
	LayoutSections = make(map[string]string)
)

func init()  {
	LayoutSections["Header"] = "common/common_header.tpl"
	LayoutSections["Footer"] = "common/common_footer.tpl"
	LayoutSections["Script"] = ""
	LayoutSections["Css"] = ""
}



func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}