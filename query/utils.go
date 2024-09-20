package query

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/tidwall/gjson"
)

func QueryByLang(s string, lang string) (string, error) {
	// TODO: 实现查询

	// 网络请求
	resp, err := queryItemAPI(s, lang)
	if err != nil {
		return "", fmt.Errorf("无法请求Market API %v", err)
	}

	// decode
	s, err = Decode(resp)
	if err != nil {
		return "", err
	}

	s, err = Decrypt(s)
	if err != nil {
		return "", err
	}

	return s, nil
}

func QueryByEng(s string) (string, error) {
	return QueryByLang(s, "en")
}

func QueryByCn(s string) (string, error) {
	encodeQuery := url.PathEscape(s)

	return QueryByLang(encodeQuery, s)
}

func QueryEuro() (int, error) {
	euros, err := QueryByCn("euro")
	if err != nil {
		return 0, err
	}
	firstElement := gjson.Get(euros, "0")
	if !firstElement.Exists() {
		return 0, fmt.Errorf("no expected field exists")
	}
	firstElement = firstElement.Get("price")
	if !firstElement.Exists() {
		return 0, fmt.Errorf("no expected field exists")
	}

	return int(firstElement.Int()), nil
}

func QueryDollar() (int, error) {
	dollars, err := QueryByCn("dollar")
	if err != nil {
		return 0, err
	}
	firstElement := gjson.Get(dollars, "0")
	if !firstElement.Exists() {
		return 0, fmt.Errorf("no expected field exists")
	}
	firstElement = firstElement.Get("price")
	if !firstElement.Exists() {
		return 0, fmt.Errorf("no expected field exists")
	}
	return int(firstElement.Int()), nil
}

func QueryExchange() (string, error) {
	var ret string
	// TODO: error handle
	c, err := QueryEuro()
	ret += "euro:" + strconv.Itoa(c) + "\t"
	if err != nil {
		return "", err
	}
	c, err = QueryDollar()
	if err != nil {
		return "", err
	}
	ret += "dollar:" + strconv.Itoa(c)
	return ret, nil
}

// 将json文本转换成便于观察的表格形式
func TarkovTable(json string) string {
	return json
}
