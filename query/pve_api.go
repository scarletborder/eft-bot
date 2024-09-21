package query

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func QueryPVEFlea(bsgId string) (string, error) {
	jsonData := fmt.Sprintf(
		`{"query":"\n     {\n         items(ids: \"%s\", lang: zh, gameMode: pve) {\n             name\n             lastLowPrice\n             avg24hPrice\n             low24hPrice\n             high24hPrice\n         }\n     }\n     "}`,
		bsgId)

	// 创建一个 HTTP 客户端
	client := &http.Client{
		Timeout: 100 * time.Second,
	}

	// 构造 HTTP POST 请求
	req, err := http.NewRequest("POST", "https://api.tarkov.dev/graphql", bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return "", fmt.Errorf("Error creating request: %v", err)
	}

	// 设置请求头，表明这是一个 JSON 请求
	req.Header.Set("Content-Type", "application/json")

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// 打印响应状态
	// fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
