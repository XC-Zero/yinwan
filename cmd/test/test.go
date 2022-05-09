package main

import (
	"fmt"
	"github.com/XC-Zero/yinwan/internal/controller/services_controller/system"
	"log"
	"math/rand"
	"time"
)

func main() {
	//config.InitConfiguration()
	//// 开协程监听配置文件修改，实现热加载
	//go config.ViperMonitor()
	//client.InitSystemStorage(config.CONFIG.StorageConfig)
	//
	//file, err := os.Open("D:\\Administrator\\Documents\\6.jpg")
	//if err != nil {
	//	panic(err)
	//}
	//stat, err := file.Stat()
	//if err != nil {
	//	panic(err)
	//}
	//objectName := "PIC_" + strconv.FormatInt(time.Now().Unix(), 10)
	//
	//info, err := client.MinioClient.PutObject(
	//	context.Background(),
	//	"images", objectName, file, stat.Size(), minio.PutObjectOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//log.Printf("%+v", info)

	a := "wtf"
	var name = fmt.Sprintf("%s-%s-%d", system.Translate(a), time.Now().Format("2006-01-02"), rand.Int63n(1000000))
	log.Println("Create bookname storage name is ", name)
}
