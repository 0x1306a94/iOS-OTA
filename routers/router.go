package routers

import (
	"iOS-OTA/controllers"
	"github.com/astaxie/beego"
	"iOS-OTA/controllers/manifest"
	"iOS-OTA/controllers/upload"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/manifest", &manifest.ManifestController{})
    beego.Router("/upload", &upload.UploadController{})
}
