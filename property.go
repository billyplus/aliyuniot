package aliyuniot

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type PropertyMsg struct {
	ID      int         `json:"id"`
	Version string      `json:"version"`
	Params  interface{} `json:"params"`
	Method  string      `json:"method"`
}

type Params map[string]interface{}

// NewPropertyMessage create property message for aliyun
// params maybe struct or map[string]interface{}
func newPropertyMessage(params interface{}) *PropertyMsg {
	return &PropertyMsg{
		ID:      getGUID(),
		Version: "1.0",
		Method:  PostProperty,
		Params:  params,
	}
}

func (dev *device) SendProperty(params interface{}) error {
	msg := newPropertyMessage(params)
	jsonstr, err := json2string(msg)
	if err != nil {
		return err
	}
	return dev.Publish(dev.postPropsTopic, 0, false, jsonstr)
}

func (dev *device) SubscribePropertyMessage(callback mqtt.MessageHandler) error {
	// dev.client.Subscribe("property/set", 0, nil)
	topic := fmt.Sprintf(ServiceTopic, dev.productKey, dev.name, "property/set")
	if err := dev.Subscribe(topic, 0, callback); err != nil {
		Error.Log("method", "subscribe", "topic", topic, "err", err.Error())
		return err
	}
	return nil
}

func json2string(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
