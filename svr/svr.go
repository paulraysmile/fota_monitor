package svr

import "github.com/paulraysmile/fota_monitor/conf"

func StartSvr() {
	for k, _ := range conf.C.Svrs {
		ProcChs[k] = make(chan *Msg, 10)
	}

	// start http svr
	go StartHttpSvr()

	// start cron svr
	go StartCronSvr()

	select {}
}