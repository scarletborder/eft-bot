package query_test

import (
	"eftbot/query"
	"testing"
)

func TestDecrypt(t *testing.T) {
	var encode string = `JTVCJNjMkgTdCJTIydWlkJTIyJTNBJTIyZGYxNjU4YWEtZDgyZC00ZDJmLTk3MzctMTk3MjAwODZlZmFmJTIyJTJDJTIyYXZnRGF5UHJpY2UlMjIlM0ExNjQlMkMlMjJhdmdXZWVrUHJpY2UlMjIlM0ExNjQlMkMlMjJjaGFuZ2UyNCUyMiUzQTAlMkMlMjJjaGFuZ2U3ZCUyMiUzQTAlMkMlMjJwcmljZSUyMiUzQTE2NCUyQyUyMnNpemUlMjIlM0ExJTJDJTIydGFncyUyMiUzQSU1QiUyMkN1cnJlbmN5JTIyJTVEJTJDJTIydXBkYXRlZCUyMiUzQSUyMjIwMjQtMDktMTlUMTAlM0EyMyUzQTQwLjY4MVolMjIlMkMlMjJzaG9ydE5hbWUlMjIlM0ElMjJFVVIlMjIlMkMlMjJic2dJZCUyMiUzQSUyMjU2OTY2ODc3NGJkYzJkYTIyOThiNDU2OCUyMiUyQyUyMnByaWNlVXBkYXRlZCUyMiUzQSUyMjIwMjAtMDYtMjBUMDIlM0ExOSUzQTQwLjQ3OVolMjIlMkMlMjJiYXNlUHJpY2UlMjIlM0ExNDIlMkMlMjJncmlkJTIyJTNBJTIyMXgxJTIyJTJDJTIybmFtZSUyMiUzQSUyMkV1cm9zJTIyJTJDJTIydHJhZGVyTmFtZSUyMiUzQSUyMkZlbmNlJTIyJTJDJTIydHJhZGVyUHJpY2UlMjIlM0E5MSUyQyUyMnRyYWRlclByaWNlQ3VyJTIyJTNBJTIyJUUyJTgyJUJEJTIyJTJDJTIyY2FuU2VsbE9uRmxlYSUyMiUzQWZhbHNlJTJDJTIyaXNGdW5jdGlvbmFsJTIyJTNBdHJ1ZSUyQyUyMnByaWNlUGVyU2xvdCUyMiUzQTE2NCUyQyUyMmF2Z0RheVByaWNlUGVyU2xvdCUyMiUzQTE2NCUyQyUyMmF2Z1dlZWtQcmljZVBlclNsb3QlMjIlM0ExNjQlMkMlMjJ1cGRhdGVkTG9uZ1RpbWVBZ28lMjIlM0F0cnVlJTJDJTIyaGF2ZU1hcmtldERhdGElMjIlM0FmYWxzZSUyQyUyMnRyYWRlclByaWNlUnViJTIyJTNBOTElMkMlMjJidXlQcmljZXMlMjIlM0ElNUIlN0IlMjJ0eXBlJTIyJTNBJTIydHJhZGVyJTIyJTJDJTIyY3VyJTIyJTNBJTIyJUUyJTgyJUJEJTIyJTJDJTIybGV2ZWwlMjIlM0ExJTJDJTIybGltaXQlMjIlM0FudWxsJTJDJTIydHJhZGVyJTIyJTNBJTIyU2tpZXIlMjIlMkMlMjJwcmljZSUyMiUzQTE2NCUyQyUyMnByaWNlQ3VyJTIyJTNBMCUyQyUyMnJlcXVpcmUlMjIlM0ElNUIlNUQlN0QlNUQlMkMlMjJzZWxsUHJpY2VzJTIyJTNBJTVCJTdCJTIydHlwZSUyMiUzQSUyMnRyYWRlciUyMiUyQyUyMmN1ciUyMiUzQSUyMiVFMiU4MiVCRCUyMiUyQyUyMmxldmVsJTIyJTNBMSUyQyUyMmxpbWl0JTIyJTNBbnVsbCUyQyUyMnRyYWRlciUyMiUzQSUyMlNraWVyJTIyJTJDJTIycHJpY2UlMjIlM0ExNjQlMkMlMjJwcmljZUN1ciUyMiUzQTAlMkMlMjJyZXF1aXJlJTIyJTNBJTVCJTVEJTdEJTJDJTdCJTIydHlwZSUyMiUzQSUyMnNlbGxUb1RyYWRlciUyMiUyQyUyMnRyYWRlciUyMiUzQSUyMkZlbmNlJTIyJTJDJTIybGV2ZWwlMjIlM0ExJTJDJTIyY3VyJTIyJTNBJTIyJUUyJTgyJUJEJTIyJTJDJTIycHJpY2UlMjIlM0E5MSUyQyUyMnByaWNlQ3VyJTIyJTNBMCU3RCU1RCUyQyUyMmZlZSUyMiUzQTAlMkMlMjJwcm9maXRGbGVhVnNUcmFkZXIlMjIlM0EwJTJDJTIydHJhZGVyQnV5UHJpY2UlMjIlM0ElN0IlMjJ0eXBlJTIyJTNBJTIydHJhZGVyJTIyJTJDJTIyY3VyJTIyJTNBJTIyJUUyJTgyJUJEJTIyJTJDJTIybGV2ZWwlMjIlM0ExJTJDJTIybGltaXQlMjIlM0FudWxsJTJDJTIydHJhZGVyJTIyJTNBJTIyU2tpZXIlMjIlMkMlMjJwcmljZSUyMiUzQTE2NCUyQyUyMnByaWNlQ3VyJTIyJTNBMCUyQyUyMnJlcXVpcmUlMjIlM0ElNUIlNUQlN0QlMkMlMjJwcm9maXRGbGlwRmxlYVRvVHJhZGVyJTIyJTNBMCUyQyUyMnByb2ZpdEZsaXBUcmFkZXJUb0ZsZWElMjIlM0EwJTJDJTIycnVOYW1lJTIyJTNBJTIyJUQwJTk1JUQwJUIyJUQxJTgwJUQwJUJFJTIyJTJDJTIydXJsJTIyJTNBJTIyZXVyb3MlMjIlMkMlMjJ3aWtpSWNvbiUyMiUzQSUyMmh0dHBzJTNBJTJGJTJGY2RuLnRhcmtvdi1tYXJrZXQuYXBwJTJGaW1hZ2VzJTJGaXRlbXMlMkZldXJvc19zbS5wbmclMjIlMkMlMjJ3aWtpSW1nJTIyJTNBJTIyaHR0cHMlM0ElMkYlMkZjZG4udGFya292LW1hcmtldC5hcHAlMkZpbWFnZXMlMkZpdGVtcyUyRmV1cm9zX2xnLnBuZyUyMiUyQyUyMndpa2lVcmwlMjIlM0ElMjJodHRwcyUzQSUyRiUyRmVzY2FwZWZyb210YXJrb3YuZmFuZG9tLmNvbSUyRndpa2klMkZFdXJvcyUyMiUyQyUyMnJ1U2hvcnROYW1lJTIyJTNBJTIyJUQwJTk1JUQwJUIyJUQxJTgwJUQwJUJFJTIyJTJDJTIyZXNOYW1lJTIyJTNBJTIyRXVyb3MlMjIlMkMlMjJlc1Nob3J0TmFtZSUyMiUzQSUyMkVVUiUyMiUyQyUyMmZyTmFtZSUyMiUzQSUyMkV1cm9zJTIyJTJDJTIyZnJTaG9ydE5hbWUlMjIlM0ElMjJFVVIlMjIlMkMlMjJkZU5hbWUlMjIlM0ElMjJFdXJvJTIyJTJDJTIyZGVTaG9ydE5hbWUlMjIlM0ElMjJFVVIlMjIlMkMlMjJjek5hbWUlMjIlM0ElMjJFdXJhJTIyJTJDJTIyY3pTaG9ydE5hbWUlMjIlM0ElMjJFVVIlMjIlMkMlMjJodU5hbWUlMjIlM0ElMjJFdXIlQzMlQjMlMjIlMkMlMjJodVNob3J0TmFtZSUyMiUzQSUyMkVVUiUyMiUyQyUyMnRyTmFtZSUyMiUzQSUyMkV1cm8lMjIlMkMlMjJ0clNob3J0TmFtZSUyMiUzQSUyMkV1cm8lMjIlMkMlMjJjbk5hbWUlMjIlM0ElMjIlRTYlQUMlQTclRTUlODUlODMlMjIlMkMlMjJjblNob3J0TmFtZSUyMiUzQSUyMiVFNiVBQyVBNyVFNSU4NSU4MyUyMiUyQyUyMmpwTmFtZSUyMiUzQSUyMiVFMyU4MyVBNiVFMyU4MyVCQyVFMyU4MyVBRCVFNyVCNCU5OSVFNSVCOSVBMyUyMiUyQyUyMmpwU2hvcnROYW1lJTIyJTNBJTIyRVVSJTIyJTJDJTIya3JOYW1lJTIyJTNBJTIyJUVDJTlDJUEwJUVCJUExJTlDJTIyJTJDJTIya3JTaG9ydE5hbWUlMjIlM0ElMjIlRUMlOUMlQTAlRUIlQTElOUMlMjIlMkMlMjJzZWFyY2glMjIlM0ElMjJldXIlMjBldXJvcyUyMCUyMCVEMCVCNSVEMCVCMiVEMSU4MCVEMCVCRSUyMCVEMCVCNSVEMCVCMiVEMSU4MCVEMCVCRSUyMGV1ciUyMGV1cm8lMjBldXIlMjBldXJvcyUyMGV1ciUyMGV1cm9zJTIwJUU2JUFDJUE3JUU1JTg1JTgzJTIwJUU2JUFDJUE3JUU1JTg1JTgzJTIwZXVyJTIwZXVyYSUyMGV1ciUyMGV1ciVDMyVCMyUyMGV1cm8lMjBldXJvJTIwZXVyJTIwJUUzJTgzJUE2JUUzJTgzJUJDJUUzJTgzJUFEJUU3JUI0JTk5JUU1JUI5JUEzJTIwJUVDJTlDJUEwJUVCJUExJTlDJTIwJUVDJTlDJUEwJUVCJUExJTlDJTIyJTdEJTVE`
	var expected string = `[{"uid":"df1658aa-d82d-4d2f-9737-19720086efaf","avgDayPrice":164,"avgWeekPrice":164,"change24":0,"change7d":0,"price":164,"size":1,"tags":["Currency"],"updated":"2024-09-19T10:23:40.681Z","shortName":"EUR","bsgId":"569668774bdc2da2298b4568","priceUpdated":"2020-06-20T02:19:40.479Z","basePrice":142,"grid":"1x1","name":"Euros","traderName":"Fence","traderPrice":91,"traderPriceCur":"₽","canSellOnFlea":false,"isFunctional":true,"pricePerSlot":164,"avgDayPricePerSlot":164,"avgWeekPricePerSlot":164,"updatedLongTimeAgo":true,"haveMarketData":false,"traderPriceRub":91,"buyPrices":[{"type":"trader","cur":"₽","level":1,"limit":null,"trader":"Skier","price":164,"priceCur":0,"require":[]}],"sellPrices":[{"type":"trader","cur":"₽","level":1,"limit":null,"trader":"Skier","price":164,"priceCur":0,"require":[]},{"type":"sellToTrader","trader":"Fence","level":1,"cur":"₽","price":91,"priceCur":0}],"fee":0,"profitFleaVsTrader":0,"traderBuyPrice":{"type":"trader","cur":"₽","level":1,"limit":null,"trader":"Skier","price":164,"priceCur":0,"require":[]},"profitFlipFleaToTrader":0,"profitFlipTraderToFlea":0,"ruName":"Евро","url":"euros","wikiIcon":"https://cdn.tarkov-market.app/images/items/euros_sm.png","wikiImg":"https://cdn.tarkov-market.app/images/items/euros_lg.png","wikiUrl":"https://escapefromtarkov.fandom.com/wiki/Euros","ruShortName":"Евро","esName":"Euros","esShortName":"EUR","frName":"Euros","frShortName":"EUR","deName":"Euro","deShortName":"EUR","czName":"Eura","czShortName":"EUR","huName":"Euró","huShortName":"EUR","trName":"Euro","trShortName":"Euro","cnName":"欧元","cnShortName":"欧元","jpName":"ユーロ紙幣","jpShortName":"EUR","krName":"유로","krShortName":"유로","search":"eur euros  евро евро eur euro eur euros eur euros 欧元 欧元 eur eura eur euró euro euro eur ユーロ紙幣 유로 유로"}]`

	res, err := query.Decrypt(encode)

	if err != nil {
		t.Errorf("error happen in decrypt %v", err)
	}

	if res != expected {
		t.Errorf("result is not equal to expected")
	}

	t.Logf("res is %v", res)
}
