package query

import (
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/tidwall/gjson"
)

// Extract vanilla response json text, and return the items field
func Decode(resp string) (string, error) {
	value := gjson.Get(resp, "items")
	if !value.Exists() {
		return "", fmt.Errorf("can not parse field `items` in raw text `%s`", resp)
	}
	return value.String(), nil
}

// 解析被加密的EFT Market 信息
func Decrypt(inputString string) (string, error) {
	// 拼接第1到5个字符和第11个字符及之后的字符
	i := inputString[:5] + inputString[10:]

	// 使用 base64 解码
	decoded, err := base64.StdEncoding.DecodeString(i)
	if err != nil {
		fmt.Println("[-] ERROR When Parse Base64")
		return "", err
	}

	// 解码URL编码
	decodedStr, err := url.QueryUnescape(string(decoded))
	if err != nil {
		fmt.Println("[-] ERROR When Decode URL")
		return "", err
	}

	return decodedStr, nil
}
