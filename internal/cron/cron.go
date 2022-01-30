package cron

import (
	cr "github.com/XC-Zero/yinwan/pkg/utils/cron"
	"github.com/XC-Zero/yinwan/pkg/utils/currency_rate"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/fwhezfwhez/errorx"
	"github.com/robfig/cron/v3"
)

func Starter() {
	c := cron.New()
	harvestCurrencyRate(c)

	c.Start()
}

// 定时任务：获取货币汇率
//  每天跑一次
func harvestCurrencyRate(c *cron.Cron) {
	_, err := c.AddFunc(cr.DAILY, func() {
		currency_rate.GetCurrencyList()
	})
	if err != nil {
		logger.Error(errorx.MustWrap(err), "初始化加载定时任务 ----- 货币汇率  失败！  ")
	}
	logger.Info("初始化加载定时任务 ----- 货币汇率  成功！  ")

}
