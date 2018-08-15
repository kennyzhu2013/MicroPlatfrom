package conf


//for log conf..
//conf written in go...
var (
	// log conf
	LoggerFlag = "os"
	LogOutPutPath = "/dev/stdout"

	//database conf
	OrmMaxIdleConn = 30
	OrmMaxOpen = 50
)
