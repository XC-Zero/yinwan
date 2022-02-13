package main

import (
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/pkg/utils/email"
	"log"
)

func main() {
	config.InitConfiguration()
	//go config.ViperMonitor()

	err := email.SendEmail("WTF", "  <div style=\"color: aqua;font-size: 35px \">FUCK YOU!</div>\n<ul>\n    <li>a</li>\n    <li>b</li>\n</ul>", "645171033@qq.com", "1285180855@qq.com")
	if err != nil {
		log.Println(err)
	}
}
