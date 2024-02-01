package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	//"github.com/astaxie/beego"

	_ "github.com/beego/beego/v2/server/web/session/memcache"
	_ "github.com/beego/beego/v2/server/web/session/mysql"
	_ "github.com/beego/beego/v2/server/web/session/redis"
	"github.com/kardianos/service"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mindoc-org/mindoc/commands"
	"github.com/mindoc-org/mindoc/commands/daemon"
	_ "github.com/mindoc-org/mindoc/routers"
)

func isViaDaemonUnix() bool {
	parentPid := os.Getppid()

	cmdLineBytes, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", parentPid))
	if err != nil {
		return false
	}

	cmdLine := string(cmdLineBytes)
	executable := strings.Split(cmdLine, " ")[0]
	fmt.Printf("Parent executable: %s\n", executable)
	filename := filepath.Base(executable)
	return strings.Contains(filename, "mindoc-daemon")
}

func main() {

	if len(os.Args) >= 3 && os.Args[1] == "service" {
		if os.Args[2] == "install" {
			daemon.Install()
		} else if os.Args[2] == "remove" {
			daemon.Uninstall()
		} else if os.Args[2] == "restart" {
			daemon.Restart()
		}
	}
	commands.RegisterCommand()

	d := daemon.NewDaemon()

	if runtime.GOOS != "windows" && !isViaDaemonUnix() {
		s, err := service.New(d, d.Config())

		if err != nil {
			fmt.Println("Create service error => ", err)
			os.Exit(1)
		}

		if err := s.Run(); err != nil {
			log.Fatal("启动程序失败 ->", err)
		}
	} else {
		d.Run()
	}
	// 设置vue静态访问映射路径
	/*vueStaticDir := beego.AppConfig.String("vueStaticDir")
	beego.SetStaticPath("/vue-admin", vueStaticDir+"index.html")
	beego.SetStaticPath("/vue-admin/css", vueStaticDir+"css")
	beego.SetStaticPath("/vue-admin/fonts", vueStaticDir+"fonts")
	beego.SetStaticPath("/vue-admin/img", vueStaticDir+"img")
	beego.SetStaticPath("/vue-admin/js", vueStaticDir+"js")*/

	// beego.Run()
}
