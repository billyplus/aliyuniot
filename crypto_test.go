package aliyuniot

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHMACSign(t *testing.T) {
	cid := "a3NYBbAltPz&fcl0o7uvt7okgUQXbuUW_aliyun-iot-device-sdk-js|securemode=3,signmethod=hmacsha1,timestamp=1547087731906|"
	name := "fcl0o7uvt7okgUQXbuUW"
	key := "a3NYBbAltPz"
	pwd := fmt.Sprintf("clientId%sdeviceName%sproductKey%stimestamp%d", cid, name, key, 1547087731906)
	secret := "a3NYBbAltPz"
	wanted := "cafe2bbc26a0c6752c649476ac7e48b753c09573"
	result := hmacSign("sha1", secret, pwd)
	assert.Equal(t, wanted, result, "sha1测试")

	wanted = "f7281708dabcb67b8c43d64e391278b6bcad991eb209ec32197e8795f5a84311"
	result = hmacSign("sha256", secret, pwd)
	assert.Equal(t, wanted, result, "sha256测试")
}
