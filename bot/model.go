package bot

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type Item struct {
	BsgId string
	ItemName
	ItemPrice
}

func (i Item) ToString() string {
	name := fmt.Sprintf("Name: %s(%s) \n 名称: %s(%s)\n", i.EnName, i.EnShortName, i.CnName, i.CnShortName) + "\n"
	var price string
	if i.TraderBuyPrice != nil {
		price += "Purchase\n" + i.TraderBuyPrice.ToString() + "\n"
	}
	if i.SellToFlea != nil {
		price += "Sell to Flea\n" + i.SellToFlea.ToString() + "\n"
	}
	if i.SellToTrader != nil {
		price += "Sell to Trader\n" + i.SellToTrader.ToString() + "\n"
	}
	return name + price
}

type ItemName struct {
	EnName      string
	EnShortName string
	CnName      string
	CnShortName string
}

type ItemPrice struct {
	*TraderBuyPrice
	*SellToFlea
	*SellToTrader
}

type TraderBuyPrice struct {
	Trader  string
	Cur     string
	Level   int
	Price   int
	Limit   int
	Require string
}

func (t *TraderBuyPrice) ToString() string {
	var buffer bytes.Buffer

	data := [][]string{
		{"商人", t.Trader},
		{"Price", fmt.Sprintf("%d%s", t.Price, t.Cur)},
		{"Level", strconv.Itoa(t.Level)},
		{"Limit", strconv.Itoa(t.Limit)},
		{"Require", t.Require},
	}

	// 创建一个 tablewriter 对象，使用 buffer 作为输出目标
	table := tablewriter.NewWriter(&buffer)
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	return buffer.String()
}

type SellToTrader struct {
	Trader string
	Cur    string // 币种
	Price  int    // 币种的数值
}

func (t *SellToTrader) ToString() string {
	var buffer bytes.Buffer

	data := [][]string{
		{"商人", t.Trader},
		{"Price", fmt.Sprintf("%d%s", t.Price, t.Cur)},
	}

	// 创建一个 tablewriter 对象，使用 buffer 作为输出目标
	table := tablewriter.NewWriter(&buffer)
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	return buffer.String()
}

type SellToFlea struct {
	LastLowPrice int
	Avg24hPrice  int
	Low24hPrice  int
	High24hPrice int
}

func (t *SellToFlea) ToString() string {
	if t.Avg24hPrice == 0 {
		return "暂无该物品出售"
	}
	var buffer bytes.Buffer

	data := [][]string{
		{"LastLowPrice", strconv.Itoa(t.LastLowPrice)},
		{"Avg24hPrice", strconv.Itoa(t.Avg24hPrice)},
		{"Low24hPrice", strconv.Itoa(t.Low24hPrice)},
		{"High24hPrice", strconv.Itoa(t.High24hPrice)},
	}

	// 创建一个 tablewriter 对象，使用 buffer 作为输出目标
	table := tablewriter.NewWriter(&buffer)
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	return buffer.String()
}
