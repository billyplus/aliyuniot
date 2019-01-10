package main

import (
	"encoding/json"
	"flag"
	"fmt"
	iot "github.com/billyplus/aliyuniot"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kit/kit/log"
	"github.com/oklog/run"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func onReceiveMessage(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)

	file := flag.String("conf", "./sample/auth.json", "auth info for aliyun")
	flag.Parse()
	data, err := ioutil.ReadFile(*file)
	if err != nil {
		logger.Log("err", err.Error())
		return
	}
	var cnf iot.DeviceConfig
	if err := json.Unmarshal(data, &cnf); err != nil {
		logger.Log("err", err.Error())
		return
	}

	dev := iot.NewDevice(cnf)
	if err := dev.Connect(); err != nil {
		logger.Log("error on connecting:", err.Error())
		return
	}
	logger.Log("msg", "connected.")
	if err := dev.SubscribePropertyMessage(onReceiveMessage); err != nil {
		logger.Log("error on subscribe property message:", err.Error())
		return
	}
	logger.Log("msg", "subscrib property message.")
	cancelChan := make(chan bool, 1)
	g := &run.Group{}
	{
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		g.Add(func() error {
			select {
			case sig := <-c:
				logger.Log("msg", "stopping", "signal", sig)
				return fmt.Errorf("receive signal: %v", sig)
			case <-cancelChan:
				return nil
			}
		}, func(err error) {
			close(cancelChan)
		})
	}
	{
		g.Add(func() error {
			i := 0
		Loop:
			for {
				select {
				case <-cancelChan:
					return nil
				default:
					i++
					if i > 10 {
						break Loop
					}
					params := &iot.Params{
						"IndoorTemperature": 30.0,
					}

					if err := dev.SendProperty(params); err != nil {
						logger.Log("err", err.Error())
					}
					logger.Log("msg", "send property successfully.")
					time.Sleep(10 * time.Second)
				}
			}
			return nil
		}, func(err error) {})
	}
	logger.Log("err", g.Run())
}
