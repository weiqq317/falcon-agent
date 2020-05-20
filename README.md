# falcon-agent

This is an agent port of open-falcon, just like the origin agent module, but can running on both windows and linux.

## Features

- besic data collection(cpu, mem, disk and etc.)
- process cpu, mem, nums collection
- snmp device custom collection
- http api to push
- single execute file, can working with windows service or linux systemd
- plugin execute support, but not support git or http sync

## Installation

it is a golang classic project

```shell
# set GOPATH and GOROOT
go get github.com/geekerlw/falcon-agent
cd $GOPATH/src/github.com/geekerlw/falcon-agent
go build -o falcon-agent.exe 	# for windows
go build -o falcon-agent 	# for linux
```

## Support metrics

### common

|  Counters   | Type  | Tag |     Notes      |
| :---------: | :---: | :-: | :------------: |
| agent.alive | GAUGE |  /  | agent is alive |

### cpu

|     Counters     | Type  | Tag |       Notes        |
| :--------------: | :---: | :-: | :----------------: |
|     cpu.user     | GAUGE |  /  |   cpu user time    |
|    cpu.system    | GAUGE |  /  |  cpu system time   |
|     cpu.idle     | GAUGE |  /  |   cpu idle time    |
|     cpu.nice     | GAUGE |  /  |   cpu nice time    |
|    cpu.iowait    | GAUGE |  /  |  cpu iowait time   |
|     cpu.irq      | GAUGE |  /  |    cpu irq time    |
|   cpu.softirq    | GAUGE |  /  |  cpu softirq time  |
|    cpu.steal     | GAUGE |  /  |   cpu steal time   |
|  cpu.guestnice   | GAUGE |  /  | cpu gusesnice time |
|    cpu.stolen    | GAUGE |  /  |  cpu stolen time   |
| cpu.used.percent | GAUGE |  /  |  cpu used percent  |

## memory

|       Counters        | Type  | Tag |             Notes              |
| :-------------------: | :---: | :-: | :----------------------------: |
|    mem.swap.total     | GAUGE |  /  |       total swap memory        |
|     mem.swap.used     | GAUGE |  /  |        used swap memory        |
|     mem.swap.free     | GAUGE |  /  |        free swap memory        |
| mem.swap.used.percent | GAUGE |  /  |    swap memory used percent    |
|       mem.total       | GAUGE |  /  |      total virtual memory      |
|     mem.available     | GAUGE |  /  | total available virtual memory |
|       mem.used        | GAUGE |  /  |          used memory           |
|       mem.free        | GAUGE |  /  |          free memory           |
|   mem.used.percent    | GAUGE |  /  |      memory used percent       |

### Disk

|     Counters      | Type  |     Tag     |    Notes     |
| :---------------: | :---: | :---------: | :----------: |
|    disk.total     | GAUGE | diskpath=%s |    total     |
|     disk.free     | GAUGE | diskpath=%s |     free     |
|     disk.used     | GAUGE | diskpath=%s |     used     |
| disk.used.percent | GAUGE | diskpath=%s | used percent |

### Network

|      Counters       |  Type   | Tag |         Notes          |
| :-----------------: | :-----: | :-: | :--------------------: |
|  net.if.bytes.send  | COUNTER |  /  | sum of all information |
|  net.if.bytes.recv  | COUNTER |  /  | sum of all information |
| net.if.packets.send | COUNTER |  /  | sum of all information |
| net.if.packets.recv | COUNTER |  /  | sum of all information |
|    net.if.err.in    | COUNTER |  /  | sum of all information |
|   net.if.err.out    | COUNTER |  /  | sum of all information |
|   net.if.drop.in    | COUNTER |  /  | sum of all information |
|   net.if.drop.out   | COUNTER |  /  | sum of all information |
|   net.if.fifo.in    | COUNTER |  /  | sum of all information |
|   net.if.fifo.out   | COUNTER |  /  | sum of all information |

### Process

|     Counters     | Type  |            Tag            |       Notes        |
| :--------------: | :---: | :-----------------------: | :----------------: |
|     proc.num     | GAUGE | name=name,cmdline=cmdline |   process number   |
| proc.cpu.percent | GAUGE | name=name,cmdline=cmdline |  process cpu use   |
| proc.mem.percent | GAUGE | name=name,cmdline=cmdline | process memory use |

### Snmp

| Counters | Type  |         Tag          |           Notes            |
| :------: | :---: | :------------------: | :------------------------: |
| snmp.get | GAUGE | addr=address,oid=oid | get oid value from address |

## Configuration

- **heartbeat**: heartbeat server rpc address
- **transfer**: transfer rpc address
- **collector**: metric configs
- **ignore**: the metrics should be ignored

Refer to `cfg.example.json`, modify the file name to `cfg.json` :

```config
{
    "debug": true,
    "hostname": "",
    "ip": "",
    "plugin": {
        "enabled": false,
        "dir": "./plugin",
        "git": "https://github.com/open-falcon/plugin.git",
        "logs": "./logs"
    },
    "heartbeat": {
        "enabled": true,
        "addr": "127.0.0.1:6030",
        "interval": 60,
        "timeout": 1000
    },
    "transfer": {
        "enabled": true,
        "addrs": [
            "127.0.0.1:8433",
            "127.0.0.1:8433"
        ],
        "interval": 60,
        "timeout": 1000
    },
    "http": {
        "enabled": true,
        "listen": ":1988",
        "backdoor": false
    },
    "collector": {
        "ifacePrefix": [], // deprecated
        "mountPoint": [], // deprecated
    },
    "default_tags": {
    },
    "ignore": {
    }
}

```

## License

This software is licensed under the Apache License. See the LICENSE file in the top distribution directory for the full license text.
using System;
using System.ComponentModel.DataAnnotations;
using System.Net.Http;
using System.Net.Http.Headers;
using Newtonsoft.Json;
namespace dotnet
{
class Program
{
　　 public class Metric
{
public string endpoint { get; set; }
public string metric { get; set; }
public int timestamp { get; set; }
public int step { get; set; }
public float value { get; set; }
public string counterType { get; set; }
public string tags { get; set; }
public string addtag { get; set; }

        }

        public static string HttpClientPost(string url,string requestJson)
        {
            try
            {
                string result = string.Empty;
                Uri postUrl = new Uri(url);

                using (HttpContent httpContent = new StringContent(requestJson))
                {
                    httpContent.Headers.ContentType = new MediaTypeHeaderValue("application/json");

                    using (var httpClient = new HttpClient())
                    {
                        httpClient.Timeout = new TimeSpan(0, 0, 60);
                        result = httpClient.PostAsync(url, httpContent).Result.Content.ReadAsStringAsync().Result;

                    }

                }
                return result;
            }
            catch (Exception e)
            {
                throw e;
            }
        }
        public static int ConvertDateTimeInt(System.DateTime time)
        {
            System.DateTime startTime = TimeZone.CurrentTimeZone.ToLocalTime(new System.DateTime(1970, 1, 1));
            return (int)(time - startTime).TotalSeconds;
        }

        static void Main(string[] args)
        {
            //"endpoint": "test-endpoint",
            //"metric": "test-metric",
            //"timestamp": ts,
            //"step": 60,
            //"value": 1,
            //"counterType": "GAUGE",
            //"tags": "location=beijing,service=falcon",
            Metric[] me = new Metric[5];

            for (int i=0; i<5; i++)
            {
                me[i] = new Metric();
                me[i].endpoint = "c# Tester";
                me[i].metric = "vscode.count";
                me[i].timestamp = ConvertDateTimeInt(DateTime.Now);
                me[i].step = 60;
                me[i].value =  0.12345f;
                me[i].counterType = "GAUGE";
                me[i].tags = "vscode=ThinkPadsX230";
                me[i].addtag = "test";

            }



            string json = JsonConvert.SerializeObject(me);


           json = HttpClientPost("http://127.0.0.1:1988/v1/push", json);
            Console.WriteLine(json);


            //Console.WriteLine("Hello World!");
           // Console.ReadKey();
        }
    }

}
