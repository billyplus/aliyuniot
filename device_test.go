package aliyuniot

import (
	"fmt"
	"testing"
)

func xTestConnect(t *testing.T) {
	cnf := DeviceConfig{}
	cnf.ProductKey = "a1NYBcAltPZ"
	cnf.DeviceSecret = "Fuyn3n8ysC6YzhGuKT96il7uqbEcg0RS"
	cnf.DeviceName = "dwl0o6uvt6pkgUQXbvUW"
	cnf.RegionID = "cn-shanghai"
	// cnf.ClientID = "faskdfjjjdsiawerwer"
	dev := NewDevice(cnf)
	if err := dev.Connect(); err != nil {
		fmt.Println("error on connecting:", err.Error())
		t.Fail()
	}
	// params := Params{
	// 	"IndoorTemperature": 30.0,
	// 	"Data":              "Hello World",
	// }
	// fmt.Println("connected.")

	// if err := dev.SendPropertyMessage(params); err != nil {
	// 	t.Error(err.Error())
	// }
}
