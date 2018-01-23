package main

import (
	_ "iOS-OTA/routers"
	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"time"
	"iOS-OTA/models"
	"iOS-OTA/models/ipa"
	"log"
	"github.com/googollee/go-socket.io"
	"net/http"
	"sync"
	"runtime"
	"errors"
	"fmt"
	"os"
	"net"
	"iOS-OTA/common"
	"flag"
)

var (
	connectClient map[string]socketio.Socket = make(map[string]socketio.Socket)
	mux sync.Mutex
)

func main() {


	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
	} else {

		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					fmt.Println("ip: ", ipnet.IP.String())
				}
			}
		}
	}


	if runtime.GOOS == "linux" {
		common.Cfg.TmpPath = flag.String("tmp","/OTA/tmp", "tmp file path")
		common.Cfg.DownloadPath = flag.String("download", "/OTA/download", "download file path")
	} else {
		common.Cfg.TmpPath = flag.String("tmp","tmp", "tmp file path")
		common.Cfg.DownloadPath = flag.String("download", "download", "download file path")
	}

	flag.Parse()

	if runtime.GOOS == "linux" {

		beego.LoadAppConfig("ini", "/OTA/conf/app.conf")
		beego.SetViewsPath("/OTA/views")
		beego.BConfig.Listen.HTTPSCertFile = "/OTA/conf/server.cer"
		beego.BConfig.Listen.HTTPSKeyFile = "/OTA/conf/server.key"
		beego.SetStaticPath("/download", *common.Cfg.DownloadPath)
		beego.SetStaticPath("/static", "/OTA/static")
		beego.BConfig.Listen.HTTPSPort = 8080

	} else {

		beego.BConfig.Listen.HTTPSCertFile = "conf/server.cer"
		beego.BConfig.Listen.HTTPSKeyFile = "conf/server.key"
		beego.SetStaticPath("/download", *common.Cfg.DownloadPath)
		beego.BConfig.Listen.HTTPSPort = 443

	}

	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.EnableHTTP = false

	retryCount := 1
	for  {
		if err := registerDataBase(); err != nil {
			log.Println("MySQL error: ", err)
			retryCount++
			if retryCount >= 10 {
				log.Fatalln("MySQL 连接失败....")
				os.Exit(0)
			}
			time.Sleep(time.Second * 5)
		} else {
			break
		}
	}

	// 启动 socket io
	socketHandler, err := createSocketHandler()
	if err != nil {
		log.Fatalln(err)
	}

	beego.BConfig.MaxMemory = 32 << 20
	beego.Handler("/socket.io/", socketHandler)
	beego.AddFuncMap("sameoutput",SameOutput)
	beego.Run()
}

func SameOutput(in string) (out string) {

	out = fmt.Sprint(in)
	log.Println("in: ", in)
	return
}

func registerDataBase() error {

	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		return errors.New("orm.RegisterDriver error: " + err.Error())
	}
	var source string
	if runtime.GOOS == "linux" {
		source = "root:123456789@tcp(mysql:3306)/ios_ota_db?charset=utf8&loc=Asia%2FShanghai"
	} else {
		source = "root:123456789@tcp(127.0.0.1:9999)/ios_ota_db?charset=utf8&loc=Asia%2FShanghai"
	}
	log.Println("source: ", source)
	maxIdle := 30
	maxConn := 30
	err = orm.RegisterDataBase("default",
		"mysql",
		source,
		maxIdle,
		maxConn)

	if err != nil {
		return errors.New("orm.RegisterDataBase error: " + err.Error())
	}

	orm.RegisterModel(new(models.NoticeObject),new(ipa.IpaInfo))
	orm.DefaultTimeLoc = time.UTC
	orm.Debug = true
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		return errors.New("orm.RunSyncdb error: " + err.Error())
	}
	return nil
}

func createSocketHandler() (http.Handler, error) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		return nil, err
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")

		// 接受到一个客户端连接 将连接的客户端保存起来 便于后续 通知
		id := so.Id()
		mux.Lock()
		connectClient[id] = so
		mux.Unlock()

		so.On("disconnection", func() {
			log.Println("on disconnect")
			mux.Lock()
			delete(connectClient, id)
			mux.Unlock()
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})
	return server, nil
}