# fota_monitor
---
中文
===

# [fota_monitor](http://github.com/paulraysmile/fota_monitor)
## 功能
* 用于进程监控，管理

## 安装
> go get -u github.com/paulraysmile/fota_monitor

## 实现
* 被监控进程启动后，按每300ms执行一次状态检测（通过发signal0信号检测），每个被监控进程在一个独立的协程里被监测。
* fota_monitor启动后会监听一个http端口用于接收管理命令（start|stop|restart|status|...）

## 使用方法
* 配置文件：[conf.json](http://github.com/paulraysmile/fota_monitor/tree/master/conf/conf.json) (json格式，支持注释)，可以通过传入自定义的env及conf参数来重定义配置文件里的参数，如：./fota_monitor -env dev -conf='port=8080;clog.mode=1'，多个参数用`;`分隔
```
{
  "dev": {
    "port": 29118, // configure admin port
    "rootpath": "/home/leizhen/Programming",
    "environ": "ulimit -n 65536",
    "svrs": {
      // demo
      "demo": "wsp/demo/demo"
    },
    "clog": {
      "name": "fota_monitor",
      "addr": "clog.srv.ns",
      "mode": 1, // 0: none, 1: localfile, 2: clog (数字代表bit位)
      "level": 15 // 0: none, 1: debug, 2: warn 4: error 8: info (数字代表bit位)
    }
  },
  "test": {
    "port": 29118, // configure admin port
    "rootpath": "/home/leizhen/Programming",
    "environ": "ulimit -n 65536",
    "svrs": {
      // demo
      "demo": "wsp/demo/demo"
    },
    "clog": {
      "name": "fota_monitor",
      "addr": "clog.srv.ns",
      "mode": 1, // 0: none, 1: localfile, 2: clog (数字代表bit位)
      "level": 15 // 0: none, 1: debug, 2: warn 4: error 8: info (数字代表bit位)
    }
  },
  "prod": {
    "port": 29118, // configure admin port
    "rootpath": "/home/leizhen/Programming",
    "environ": "ulimit -n 65536",
    "svrs": {
      // demo
      "demo": "wsp/demo/demo"
    },
    "clog": {
      "name": "fota_monitor",
      "addr": "clog.srv.ns",
      "mode": 2, // 0: none, 1: localfile, 2: clog (数字代表bit位)
      "level": 14 // 0: none, 1: debug, 2: warn 4: error 8: info (数字代表bit位)
    }
  }
}
```
* 运行方法：fota_monitor.sh [start|stop|restart|status|check]
* 进程管理：fota_monitor -[h|env|status|start|stop|restart] [dev|test|prod|all|["svrname"]]

## 注意
* fota_monitor的运行日志通过clog上报，也可记录在本地fota_monitor.log日志文件里，注意：此fota_monitor.log日志文件不会被切分，所以尽量保持较少的日志输出，建议通过clog方式上报日志
* fota_monitor启动监控进程后，被监控进程控制台日志fota_monitor.log会输出到相应进程目录，最多保存30天，历史日志以fota_monitor.{day}.log方式备份
* 当fota_monitor启动时，会根据conf.json配置启动所有被监控进程，当被监控进程已经启动过，并且符合配置要求时，fota_monitor会自动将其加入监控列表
* fota_monitor会定期检查进程运行状态，如果进程异常退出，fota_monitor会反复重试拉起，并且记录日志
* 当被监控进程为多进程运行模式，fota_monitor只监控管理父进程(子进程应实现检测父进程运行状态，并随父进程退出而退出）
* 被监控进程以nohup方式启动，所以你的程序就不要自己设定daemon运行了
* 每分钟通过ps方式检测一次进程状态，如果出现任何异常，比如有多份进程启动等，记日志
* 由于fota_monitor会同时启动内部httpserver（绑内网ip），所以也支持远程管理，比如在浏览器里输入：http://xxx.xxx.xxx.xxx:29118/?command=status&service=all

## demo
```
$ fota_monitor -env dev -status all

*****STATUS OK SERVICE LIST*****
demo PID:13539

*****STATUS FAIL SERVICE LIST*****

$ fota_monitor -env dev -restart demo
SUCCESS
```
