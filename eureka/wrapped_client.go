package eureka

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var client *Client
var instanceRegistered *InstanceInfo

func Start(appName string, port int, machines []string) {

	instanceId := fmt.Sprintf("%s:%s:%d", GetLocalAddress(), appName, port)
	instanceRegistered = NewInstanceInfo(instanceId, GetLocalAddress(), appName, GetLocalAddress(), appName, port, 30, false) //Create a new instance to register
	instanceRegistered.Metadata = &MetaData{
		Map: make(map[string]string),
	}

	client = NewClient(machines)
	client.RegisterInstance(appName, instanceRegistered) // Register new instance in your eureka(s)

	go sendHeartbeat(appName, instanceId)

	registerSignal()
}

func sendHeartbeat(appName string, instanceId string) {
	func() {
		for {
			err := client.SendHeartbeat(appName, instanceId)
			if err != nil {
				logger.Warning("Send HeartBeat failed for appname=%s,instanceid=%s,err=%v,doing register again", appName, instanceId, err)
				client.RegisterInstance(appName, instanceRegistered)
			}
			time.Sleep(time.Second * 30)
			logger.Info("send heartbeat for appname=%s,instanceid=%s", appName, instanceId)
		}
	}()
}

func unregisterInstance() {
	if client != nil {
		client.UnregisterInstance(instanceRegistered.App, instanceRegistered.InstanceId)
		client = nil
	}
}

func registerSignal() {
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signals

		logger.Info("got signal=%s, unregister eureka ", sig.String())

		unregisterInstance()

		os.Exit(0)
	}()

}
