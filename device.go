package aliyuniot

import (
	// "crypto/tls"
	"errors"
	"fmt"
	// "net"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	// "github.com/jeffallen/mqtt"
	"time"
)

type Device interface {
	Connect() error
	// SendProperty send property message to aliyun
	// params maybe struct or map[string]interface{}
	SendProperty(params interface{}) error
	SubscribePropertyMessage(callback mqtt.MessageHandler) error
	Publish(topic string, qos byte, retained bool, message string) error
	Subscribe(topic string, qos byte, callback mqtt.MessageHandler) error
}

type device struct {
	regionID            string
	signAlgorithm       string
	productKey          string
	name                string
	secret              string
	brokerURL           string
	clientID            string
	brokerProtocol      string
	postPropsTopic      string
	postPropsReplyTopic string
	postEventTopic      string
	postEventReplyTopic string
	serviceTopic        string
	brokerPort          int
	secureMode          int
	timestamp           int64
	tls                 bool
	client              mqtt.Client
}

// NewDevice create new aliyun iot device from device config
func NewDevice(conf DeviceConfig) Device {
	if conf.ProductKey == "" {
		panic(errors.New("productKey should not be empty"))
	}
	if conf.DeviceName == "" {
		panic(errors.New("deviceName should not be empty"))
	}
	if conf.DeviceSecret == "" {
		panic(errors.New("deviceSecret should not be empty"))
	}
	dev := &device{
		productKey: conf.ProductKey,
		name:       conf.DeviceName,
		secret:     conf.DeviceSecret,
	}

	dev.timestamp = int64(time.Now().UnixNano() / 1000000)

	if conf.Websocket {
		if conf.TLS {
			dev.brokerProtocol = ProtocolWSS
		} else {
			dev.brokerProtocol = ProtocolWS
		}
		dev.brokerPort = 443
	} else {
		if conf.TLS {
			dev.brokerProtocol = ProtocolTLS
		} else {
			dev.brokerProtocol = ProtocolTCP
		}
		dev.brokerPort = 1883
	}
	// client ID
	if conf.ClientID != "" {
		dev.clientID = conf.ClientID
	} else {
		dev.clientID = dev.productKey + "&" + dev.name + "_aliyun-iot-golang"
	}
	//默认区域
	if conf.RegionID != "" {
		dev.regionID = conf.RegionID
	} else {
		dev.regionID = DefaultRegionID
	}
	if conf.SignAlgorithm == "" {
		dev.signAlgorithm = HMACSHA1
	} else {
		dev.signAlgorithm = conf.SignAlgorithm
	}
	// 默认broker URL
	if conf.BrokerURL != "" {
		dev.brokerURL = conf.BrokerURL
	} else {
		dev.brokerURL = fmt.Sprintf(BrokerURL, dev.brokerProtocol, dev.productKey, dev.regionID, dev.brokerPort)
	}

	dev.secureMode = 3

	// topic
	dev.postPropsTopic = fmt.Sprintf(PropertyPostTopic, dev.productKey, dev.name)
	dev.postPropsReplyTopic = fmt.Sprintf(PropertyPostReplyTopic, dev.productKey, dev.name)
	// dev.postEventTopic = fmt.Sprintf(EventPostTopic, dev.productKey, dev.name)
	// fmt.Println(dev.postEventTopic)
	// dev.postEventReplyTopic = fmt.Sprintf(EventPostReplyTopic, dev.productKey, dev.name)
	// dev.serviceTopic = fmt.Sprintf(ServiceTopic, dev.productKey, dev.name)
	return dev
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}

func (dev *device) Connect() error {
	id := fmt.Sprintf("%s|securemode=%d,signmethod=hmac%s,timestamp=%d|", dev.clientID, dev.secureMode, dev.signAlgorithm, dev.timestamp)
	user := dev.name + "&" + dev.productKey
	pwd := fmt.Sprintf("clientId%sdeviceName%sproductKey%stimestamp%d", dev.clientID, dev.name, dev.productKey, dev.timestamp)
	passwd := hmacSign(dev.signAlgorithm, dev.secret, pwd)

	opt := mqtt.NewClientOptions()
	opt.AddBroker(dev.brokerURL)
	opt.SetClientID(id)
	opt.SetUsername(user)
	opt.SetPassword(passwd)
	opt.SetKeepAlive(60 * time.Second)
	opt.SetCleanSession(true)
	// conn, err := net.Dial("tcp", "a1NYBcAltPZ.iot-as-mqtt.cn-shanghai.aliyuncs.com:1883")
	// if err != nil {
	// 	fmt.Println("dial err:", err.Error())
	// 	return err
	// }
	// clientconn := mqtt.NewClientConn(conn)
	// clientconn.ClientId = "faskdfjjjdsiawerwer|securemode=3,signmethod=hmacsha1,timestamp=10|"
	// user = "dwl0o6uvt6pkgUQXbvUW&a1NYBcAltPZ"
	// passwd = "87e4dcead2f992cf4f56c815dc1fa5f61a090af3"
	// if err = clientconn.Connect(user, passwd); err != nil {
	// 	fmt.Println("connect err:", err.Error())
	// 	return err
	// }

	// tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	// opt.SetTLSConfig(tlsConfig)
	opt.OnConnect = func(c mqtt.Client) {
		fmt.Println("connected. ")
	}

	client := mqtt.NewClient(opt)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	dev.client = client

	return nil
}

func (dev *device) Publish(topic string, qos byte, retained bool, message string) error {
	token := dev.client.Publish(dev.postPropsTopic, qos, retained, message)
	token.Wait()
	return token.Error()
}

func (dev *device) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) error {
	token := dev.client.Subscribe(topic, qos, callback)
	token.Wait()
	return token.Error()
}
