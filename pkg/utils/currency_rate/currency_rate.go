package currency_rate

/*  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *
	货币汇率

	货币汇率使用的是第三方的免费API

    https://www.nowapi.com/api/finance.rate

	所以有次数限制，目前限制为 每小时 50次调用
	且免费用户不支持一对多，多对多查询，只能委屈的间隔三分钟一个一个查


	QAQ
*  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  *  */
import (
	"fmt"
	"github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/pkg/utils/logger"
	"github.com/dollarkillerx/urllib"
	"github.com/fwhezfwhez/errorx"
	"time"
)

const MAX_RETRY = 3

var Interval = time.Minute * 3
var currencyList currencyListResponse

// Retry 重试次数
var retry = MAX_RETRY

// CurrencyNameMap 货币中文名和缩写对照
var CurrencyNameMap = make(map[string]string)

// CurrencyNameReverseMap 缩写到货币中文名对照
var CurrencyNameReverseMap = make(map[string]string)

// CurrencyRateMap 各货币与人民币汇率表
var CurrencyRateMap = make(map[string]CurrencyRate)

// 下面俩个都是临时结构体，没什么意义，用来解析接收api返回的参数而已
type currencyListResponse struct {
	Success string `json:"success"`
	Result  struct {
		TotLine    string `json:"totLine"`
		LastUpdate string `json:"lastUpdate"`
		Lists      []struct {
			CurNo   string `json:"curNo"`
			CurNm   string `json:"curNm"`
			CurNmEn string `json:"curNmEn"`
		} `json:"lists"`
	} `json:"result"`
}

type currencyRate struct {
	Success string `json:"success"`
	Result  struct {
		Status string `json:"status"`
		Scur   string `json:"scur"`
		Tcur   string `json:"tcur"`
		Ratenm string `json:"ratenm"`
		Rate   string `json:"rate"`
		Update string `json:"update"`
	} `json:"result"`
}

type CurrencyRate struct {
	Ratenm string `json:"ratenm"`
	Rate   string `json:"rate"`
	Update string `json:"update"`
}

// GetCurrencyList 获取货币列表
// 下面是以Get请求为例的测试链接
// GET https://sapi.k780.com/?app=finance.rate_curlist&curType=rateRealtime&appkey=64165&sign=39fa1b9f58cb3fcbaea198b869d9c243
func GetCurrencyList() {
	cfg := config.CONFIG.ApiConfig.CurrencyRateConfig
	listUrl := fmt.Sprintf("%s&appkey=%s&sign=%s", cfg.ListURL, cfg.AppKey, cfg.Sign)
	err := urllib.Get(listUrl).FromJson(&currencyList)
	if err != nil {
		retry--
		logger.Error(errorx.MustWrap(err), fmt.Sprintf("调用接口用于获取货币列表失败,正在重试 %d/%d 次 ！ ", MAX_RETRY-retry, MAX_RETRY))
		if retry < 1 {
			return
		} else {
			GetCurrencyList()
		}
	}
	for _, currency := range currencyList.Result.Lists {
		CurrencyNameMap[currency.CurNm] = currency.CurNo
		CurrencyNameReverseMap[currency.CurNo] = currency.CurNm
		rate := GetCurrencyRate(currency.CurNo, "CNY").Result
		CurrencyRateMap[currency.CurNo] = CurrencyRate{
			Ratenm: rate.Ratenm,
			Rate:   rate.Rate,
			Update: rate.Update,
		}
		logger.Info(fmt.Sprintf("更新 %s（%s） 到 CNY 的汇率成功！", currency.CurNo, currency.CurNm))
		time.Sleep(Interval)
	}
}

// GetCurrencyRate 获取某一货币的汇率
// 下面是以Get请求为例的测试链接
// GET http://api.k780.com?app=finance.rate&scur=AED&tcur=CNY&appkey=64165&sign=39fa1b9f58cb3fcbaea198b869d9c243
func GetCurrencyRate(sourceCurrency, targetCurrency string) currencyRate {
	cfg, currencyRate := config.CONFIG.ApiConfig.CurrencyRateConfig, currencyRate{}
	err := urllib.Get(fmt.Sprintf("%s&scur=%s&tcur=%s&appkey=%s&sign=%s",
		cfg.RateURL, sourceCurrency, targetCurrency, cfg.AppKey, cfg.Sign,
	)).FromJson(&currencyRate)
	if err != nil {
		logger.Error(errorx.MustWrap(err), fmt.Sprintf("获取 %s 货币（%s）到 %s 货币（%s）的汇率失败！ ",
			sourceCurrency, CurrencyNameReverseMap[sourceCurrency], targetCurrency, CurrencyNameReverseMap[targetCurrency]))
	}
	return currencyRate
}
