package aliyuniot

var (
	guid = 0
)

func getGUID() int {
	guid++
	return guid
}
