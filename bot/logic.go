package bot

// logic,
// All functions will return text which will be directly sent to qq IM

import (
	"bytes"
	"eftbot/query"
	"fmt"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/tidwall/gjson"
	"github.com/zeromicro/go-zero/core/mr"
)

func HelpFunc() string {
	// 创建一个 bytes.Buffer 来存储表格字符串
	var buffer bytes.Buffer

	data := [][]string{
		{"帮助", "查看当前可用的指令"},
		{"help", "Display current supported command"},
		{"汇率", "展示游戏中美元,欧元汇率"},
		{"exchange", "Display the US dollar and euro exchange rates in the game"},
		{"x87 {num}", "Calculate item's price in euro through 口关's theory"},
		{"Other Input...", "Display other item's flea market information"},
	}

	// 创建一个 tablewriter 对象，使用 buffer 作为输出目标
	table := tablewriter.NewWriter(&buffer)
	table.SetHeader([]string{"Command", "Meaning"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render() // 渲染表格到 buffer

	// 返回表格的字符串表示
	return buffer.String()
}

func ExchangeFunc() string {
	euro_c, err := query.QueryEuro()
	if err != nil {
		return fmt.Sprintf("查询失败 Query Fail %s", err.Error())
	}

	dollar_c, err := query.QueryDollar()
	if err != nil {
		return fmt.Sprintf("查询失败 Query Fail %s", err.Error())
	}

	// 创建一个 bytes.Buffer 来存储表格字符串
	var buffer bytes.Buffer

	data := [][]string{
		{"Dollar", strconv.Itoa(dollar_c), fmt.Sprintf("%.3f", float32(dollar_c)*0.87)},
		{"Euro", strconv.Itoa(euro_c), fmt.Sprintf("%.3f", float32(euro_c)*0.87)},
	}

	// 创建一个 tablewriter 对象，使用 buffer 作为输出目标
	table := tablewriter.NewWriter(&buffer)
	table.SetHeader([]string{"Type", "Ruble", "x87"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render() // 渲染表格到 buffer

	// 返回表格的字符串表示
	return buffer.String()
}

// 分析json文本形式的items
// 返回ASCII表格形式字符串
func ParseComplexItems(json string) string {
	item_list_res := gjson.Get(json, "@this")
	// 检查结果是否是一个数组
	if item_list_res.IsArray() {
		item_list := item_list_res.Array()
		result, err := mr.MapReduce(func(source chan<- *Item) {
			for _, res_item := range item_list {
				bsg_Id := res_item.Get("bsgId")
				if !bsg_Id.Exists() {
					continue
				}
				var temp_item Item
				temp_item.BsgId = bsg_Id.String()
				temp_item.ItemName.EnName = res_item.Get("name").String()
				temp_item.ItemName.EnShortName = res_item.Get("shortName").String()
				temp_item.ItemName.CnName = res_item.Get("cnName").String()
				temp_item.ItemName.CnShortName = res_item.Get("cnShortName").String()

				traderBuyPrice := res_item.Get("traderBuyPrice")

				if traderBuyPrice.Exists() {
					cur := traderBuyPrice.Get("cur").Raw
					cur = strings.Trim(cur, "\"")
					var price int
					if cur != "₽" {
						price = int(traderBuyPrice.Get("priceCur").Int())
					} else {
						price = int(traderBuyPrice.Get("price").Int())
					}

					temp_item.ItemPrice.TraderBuyPrice = &TraderBuyPrice{
						Trader:  traderBuyPrice.Get("trader").Raw,
						Cur:     cur,
						Level:   int(traderBuyPrice.Get("level").Int()),
						Price:   price,
						Limit:   int(traderBuyPrice.Get("limit").Int()),
						Require: traderBuyPrice.Get("require").Raw,
					}
				}

				sellPrices := res_item.Get("sellPrices")
				if sellPrices.IsArray() {
					sellPricesList := sellPrices.Array()
					for _, priceItem := range sellPricesList {
						if pr := priceItem.Get("type"); pr.Exists() && pr.String() == "sellToTrader" {
							cur := priceItem.Get("cur").Raw // TODO: can not successfully get roble sign
							cur = strings.Trim(cur, "\"")
							var price int
							if cur != "₽" {
								price = int(priceItem.Get("priceCur").Int())
							} else {
								price = int(priceItem.Get("price").Int())
							}
							temp_item.ItemPrice.SellToTrader = &SellToTrader{
								Trader: priceItem.Get("trader").Raw,
								Cur:    cur,
								Price:  price,
							}
						}
					}
				}
				source <- &temp_item
			}
		}, func(item *Item, writer mr.Writer[*Item], cancel func(error)) {
			defer writer.Write(item)
			res, err := query.QueryPVEFlea(item.BsgId)
			if err != nil {
				return
			}
			res_json := gjson.Get(res, "data")
			if !res_json.Exists() {
				return
			}
			res_json = res_json.Get("items")
			if !res_json.Exists() || !res_json.IsArray() {
				return
			}
			price_json := res_json.Array()[0]
			item.ItemPrice.SellToFlea = &SellToFlea{
				LastLowPrice: int(price_json.Get("lastLowPrice").Int()),
				Avg24hPrice:  int(price_json.Get("avg24hPrice").Int()),
				Low24hPrice:  int(price_json.Get("low24hPrice").Int()),
				High24hPrice: int(price_json.Get("high24hPrice").Int()),
			}
		}, func(pipe <-chan *Item, writer mr.Writer[string], cancel func(error)) {
			var data string
			for item := range pipe {
				data += item.ToString()
			}
			writer.Write(data)
		})

		if err != nil {
			// TODO: handle error
			_ = err
		}

		// TODO: result转换成table
		return result
	} else {
		return fmt.Sprintf("无法识别的json array %s", json)
	}
}
