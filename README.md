Go Eureka Client
================

Based on code from https://github.com/bryanstephens/go-eureka-client .

## Getting started

```go
import (
	"github.com/nxsre/go-eureka-client/eureka"
)
func main() {

	client := eureka.NewClient([]string{
		"http://127.0.0.1:8761/eureka", //From a spring boot based eureka server
		// add others servers here
	})
	appName:="myapp"
	port:=8080
	instanceId := fmt.Sprintf("%s:%s:%d", getIp(), appName, port)
    instanceRegistered = eureka.NewInstanceInfo(instanceId, getIp(), appName, getIp(), appName, port, 30, false) //Create a new instance to register
    instanceRegistered.Metadata = &eureka.MetaData{
        Map: make(map[string]string),
    }
    client.RegisterInstance(appName, instanceRegistered) // Register new instance in your eureka(s)

	applications, _ := client.GetApplications() // Retrieves all applications from eureka server(s)
	client.GetApplication(instance.App) // retrieve the application "test"
	client.GetInstance(instance.App, instance.HostName) // retrieve the instance from "test.com" inside "test"" app
	client.SendHeartbeat(instance.App, instance.HostName) // say to eureka that your app is alive (here you must send heartbeat before 30 sec)
}
```

**Note:**
- `appId` here is the name of the app
- `instanceId` is the hostname of the app
- When calling `RegisterInstance` the `appId` is needed but not used by eureka, this is not the appId but a whatever value

All these strange behaviour come from Eureka.

## Create Client from a config file

You can create from a json file with this form (here we called it `config.json`):

```json
{
  "config": {
    "certFile": "",
    "keyFile": "",
    "caCertFiles": null,
    "timeout": 1000000000,
    "consistency": ""
  },
  "cluster": {
    "leader": "http://127.0.0.1:8761/eureka",
    "machines": [
      "http://127.0.0.1:8761/eureka"
    ]
  }
}
```

And to load it:

```go
client := NewClientFromFile("config.json")
```


## update 2017-12-05
Add Support for springcloud eureka