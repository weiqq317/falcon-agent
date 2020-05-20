// Copyright 2017 Xiaomi, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"os"

	"falcon-agent/cron"
	"falcon-agent/funcs"
	"falcon-agent/g"
	"falcon-agent/http"
)

func main() {

	cfg := flag.String("c", "cfg.json", "配置文件")
	version := flag.Bool("v", false, "查看版本号")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	if g.Config().Debug {
		g.InitLog("debug")
	} else {
		g.InitLog("info")
	}

	g.InitRootDir()    //获取当前目前 给Root
	g.InitLocalIp()    //获取本地ip 给 LocalIp
	g.InitRpcClients() //初始化心跳连接 HbsClient

	funcs.BuildMappers() //定时采集上报数据

	go cron.InitDataHistory()

	cron.ReportAgentStatus()
	cron.SyncMinePlugins()
	cron.SyncBuiltinMetrics()
	cron.SyncTrustableIps()
	cron.Collect()

	go http.Start()

	//g.HideConsole()

	select {}

}
