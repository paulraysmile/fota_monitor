// 进程监控服务
// author：leizhen
// date：2018/12/12
package main

import (
	"fmt"
	"time"

	"github.com/paulraysmile/clog/api"
	"github.com/paulraysmile/fota_monitor/comm"
	"github.com/paulraysmile/fota_monitor/conf"
	"github.com/paulraysmile/fota_monitor/svr"
	"github.com/paulraysmile/namecli/api"
	"github.com/paulraysmile/utils"
)

func request(command string, service string) {
	url := fmt.Sprintf("http://%s:%d", utils.LocalIp, conf.C.Port)
	params := map[string]string{
		"command": command,
		"service": service,
	}
	gpp := &utils.GPP{
		Uri:     url,
		Timeout: time.Second * 8,
		Params:  params,
	}
	body, err := utils.Get(gpp)
	if err != nil {
		fmt.Printf("Error: [fota_monitor maybe down!] %v, %s\n", err, body)
		return
	}

	fmt.Printf(string(body))
	return
}

func init() {
	clog.AddrFunc = func() (string, error) {
		return api.Name(conf.C.Clog.Addr)
	}
	clog.Init(conf.C.Clog.Name, "", conf.C.Clog.level, conf.C.Clog.Mode)
}

func main() {
	switch {
	case conf.Start != "":
		request(comm.START, conf.Start)
	case conf.Stop != "":
		request(comm.STOP, conf.Stop)
	case conf.Restart != "":
		request(comm.RESTART, conf.Restart)
	case conf.Status != "":
		request(comm.STATUS, conf.Status)
	default:
		clog.Info("main() StartSvr")
		svr.StartSvr()
	}
}
