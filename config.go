package aliyuniot

const (
	DefaultRegionID = "cn-shanghai"
	BrokerURL       = "%s%s.iot-as-mqtt.%s.aliyuncs.com:%d"
)

var (
	tlsPrefix = []string{"tls://", "mqtts://", "wss://"}
)

type DeviceConfig struct {
	RegionID      string
	SignAlgorithm string
	ProductKey    string
	DeviceName    string
	DeviceSecret  string
	BrokerURL     string
	ClientID      string
	TLS           bool
	Websocket     bool
}

const (
	DefaultProtocol = ProtocolWS
	ProtocolWS      = "ws://"
	ProtocolTCP     = "tcp://"
	ProtocolTLS     = "tls://"
	ProtocolWSS     = "wss://"
	// ProtocolMQTT    = "mqtt://"
	// ProtocolMQTTS   = "mqtts://"
)

var (
	PostProperty             = "thing.event.property.post"
	PostEvent                = "thing.event.%s.post"
	ThingServiceMethodPrefix = "thing.service."
)

const (
	ServiceTopic           = "/sys/%s/%s/thing/service/%s"
	PropertyPostTopic      = "/sys/%s/%s/thing/event/property/post"
	PropertyPostReplyTopic = "/sys/%s/%s/thing/event/property/post_reply"
	EventPostTopic         = "/sys/%s/%s/thing/event/%s/post"
	EventPostReplyTopic    = "/sys/%s/%s/thing/event/%s/post_reply"
)
