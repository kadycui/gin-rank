package conf

var (
	STATIC_URL = "/fan/static/" // 静态资源目录
	MYSQL_URL  = struct {
		UserName string
		Password string
		Ip       string
		Port     string
		DataBase string
	}{UserName: "root", Password: "123456", Ip: "127.0.0.1", Port: "3306", DataBase: "test"}
	REDIS_URL = struct {
		Ip       string
		Port     string
		Password string
	}{Ip: "10.16.168.61", Port: "6379", Password: "123456"}

	LOG_CONF = struct {
		LogFilePath string
		LogFileName string
	}{LogFileName: "web.logs", LogFilePath: "logs/"}
)

func init() {

}
