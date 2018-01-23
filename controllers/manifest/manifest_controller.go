package manifest

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"iOS-OTA/models/ipa"
)

type ManifestController struct {
	beego.Controller
}

func (c *ManifestController) Get()  {

	o := orm.NewOrm()
	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		c.Abort("502")
		return
	}

	qb.Select("*").
		From("ipa_info").
		OrderBy("id").
		Desc().
		Limit(1)

	sql := qb.String()

	fmt.Println("sql: ", sql)
	obj := new(ipa.IpaInfo)
	o.Raw(sql).QueryRow(&obj)

	if obj.Name == "" {
		c.Abort("404")
		return
	}
	fmt.Println("tmpObj: ", obj)

	content := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>items</key>
	<array>
		<dict>
			<key>assets</key>
			<array>
				<dict>
					<key>kind</key>
					<string>software-package</string>
					<key>url</key>
					<string>https://test.me%s</string>
				</dict>
				<dict>
					<key>kind</key>
					<string>display-image</string>
					<key>url</key>
					<string>https://test.me%s</string>
				</dict>
				<dict>
					<key>kind</key>
					<string>full-size-image</string>
					<key>url</key>
					<string>https://test.me%s</string>
				</dict>
			</array>
			<key>metadata</key>
			<dict>
				<key>bundle-identifier</key>
				<string>%s</string>
				<key>bundle-version</key>
				<string>%s</string>
				<key>kind</key>
				<string>software</string>
				<key>title</key>
				<string>%s</string>
			</dict>
		</dict>
	</array>
</dict>
</plist>
`
	c.Ctx.Request.Header.Set("Content-Type", "text/xml")
	c.Ctx.WriteString(
		fmt.Sprintf(content,obj.Path,
		obj.IconPath,
		obj.IconPath,
		obj.Identifier,
		obj.Version,
		obj.Name))

}