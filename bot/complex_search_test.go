package bot_test

import (
	"eftbot/bot"
	"eftbot/query"
	"log"
	"testing"
)

func TestComplexSearch(t *testing.T) {
	res, _ := query.QueryByCn("milk")
	res = bot.ParseComplexItems(res)
	log.Print(res)
}
