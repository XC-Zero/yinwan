package logger

import (
	"fmt"
	"log"
	"time"
)

//todo 将所有的 log 替换为 logger

//goland:noinspection GoSnakeCaseUsage
const SPLIT_LINE = "\n----------------------------------------------------\n"

//goland:noinspection GoSnakeCaseUsage,GoNameStartsWithPackageName
const LOGGER_TOPIC = "System-Log"

type LogType string

const (
	FATAL  LogType = "fatal"
	ERROR  LogType = "error"
	INFO   LogType = "info"
	WARING LogType = "waring"
)

type LogHeader string

//goland:noinspection GoSnakeCaseUsage,GoSnakeCaseUsage,GoSnakeCaseUsage,GoSnakeCaseUsage
const (
	FATAL_HEADER  LogHeader = "$#8000fa$(致命问题)"
	ERROR_HEADER  LogHeader = "$#e50c3e$(错误)"
	WARING_HEADER LogHeader = "$#fabe07$(警告)"
	INFO_HEADER   LogHeader = "$#04ea5f$(信息)"
)

// Logger 打印日志到系统日志文件中，应符合日志打印的基本原则
type Logger struct {
	message   string
	stackInfo []string
	logType   LogType
	logHeader LogHeader
	time      time.Time
}

// Fatal 慎用，会导致程序退出！
func Fatal(err error, mes string) {

	sprintf := fmt.Sprintf("[PANIC] 程序退出！ %s %s \n fatal error is %+v %s", SPLIT_LINE, mes, err, SPLIT_LINE)

	log.Fatalln(sprintf)
}

// Error 打印错误信息
func Error(err error, mes string) {
	sprintf := fmt.Sprintf("[ERROR] %s %s \n error is %+v %s", SPLIT_LINE, mes, err, SPLIT_LINE)
	//if OperateLogSocketClient != nil {
	//	OperateLogSocketClient.WriteMessage(websocket.TextMessage, []byte(mes))
	//}
	//if SystemLogSocketClient != nil {
	//	SystemLogSocketClient.WriteMessage(websocket.TextMessage, []byte(sprintf))
	//}
	log.Println(sprintf)
	//Logger{
	//	message:   sprintf,
	//	stackInfo: err.StackTraces,
	//	logType:   ERROR,
	//	logHeader: ERROR_HEADER,
	//	time:      time.Now(),
	//}.sendLogger()
}

// Waring 打印警告信息
func Waring(err error, mes string) {
	sprintf := fmt.Sprintf("[WARING] %s %s \n error is %+v %s", SPLIT_LINE, mes, err, SPLIT_LINE)
	log.Println(sprintf)
	//if SystemLogSocketClient != nil {
	//	SystemLogSocketClient.WriteMessage(websocket.TextMessage, []byte(sprintf))
	//}
	//Logger{
	//	message:   sprintf,
	//	stackInfo: err.StackTraces,
	//	logType:   WARING,
	//	logHeader: WARING_HEADER,
	//	time:      time.Now(),
	//}.sendLogger()
}

// Info 打印普通信息
func Info(mes string) {
	log.Println("[INFO]  " + mes)
	//if SystemLogSocketClient != nil {
	//	SystemLogSocketClient.WriteMessage(websocket.TextMessage, []byte(mes))
	//}
	//Logger{
	//	message:   mes,
	//	logType:   INFO,
	//	logHeader: INFO_HEADER,
	//	time:      time.Now(),
	//}.sendLogger()
}

//func (l Logger) sendLogger() {
//	err := client.PushInterfaceToKafka(LOGGER_TOPIC, []interface{}{l})
//	if err != nil {
//		log.Println(err)
//	}
//}
