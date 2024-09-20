package query

import (
	"fmt"
	"io"
	"log"
	"time"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

func queryItemAPI(s string, lang string) (string, error) {
	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(profiles.Chrome_120),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar), // create cookieJar instance and pass it as argument
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		log.Println(err)
		return "", err
	}

	q_uri := fmt.Sprintf("https://tarkov-market.com/api/be/items?lang=%s&search=%s&tag=&sort=change24&sort_direction=desc&trader=&skip=0&limit=20", lang, s)

	req, err := http.NewRequest(http.MethodGet, q_uri, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}

	req.Header = http.Header{
		"authority":       {"tarkov-market.com"},
		"accept":          {"*/*"},
		"accept-language": {"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6,zh-TW;q=0.5,zh-HK;q=0.4"},
		"user-agent":      {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"},
		http.HeaderOrderKey: {
			"accept",
			"accept-language",
			"user-agent",
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	readBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(readBytes), nil
}

func QueryWeather() (string, error) {
	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(profiles.Chrome_120),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar), // create cookieJar instance and pass it as argument
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		log.Println(err)
		return "", err
	}

	q_uri := "https://tarkov-market.com/api/be/weather"

	req, err := http.NewRequest(http.MethodGet, q_uri, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}

	req.Header = http.Header{
		"authority":       {"tarkov-market.com"},
		"accept":          {"*/*"},
		"accept-language": {"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6,zh-TW;q=0.5,zh-HK;q=0.4"},
		"user-agent":      {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"},
		http.HeaderOrderKey: {
			"accept",
			"accept-language",
			"user-agent",
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	readBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(readBytes), nil
}

const tarkovRatio = 7

// hrs 将小时数转换为毫秒数
func hrs(num int64) int64 {
	return num * 60 * 60 * 1000
}

// realTimeToTarkovTime 将现实时间转换为塔科夫时间
func realTimeToTarkovTime(t time.Time, left bool) time.Time {
	// 一天的毫秒数
	oneDay := hrs(24)
	// UTC 偏移量，3小时
	utcOffset := hrs(3)
	// 获取UTC时间戳（以毫秒为单位）
	utcTimestamp := t.UTC().UnixMilli()

	// 根据 left 参数决定时间偏移量，如果 left 为 true，使用白天时间，否则加上 12 小时
	offset := utcOffset
	if !left {
		offset += hrs(12)
	}

	// 计算塔科夫时间，取模 oneDay 以确保时间在 24 小时范围内
	tarkovTimestamp := (offset + utcTimestamp*tarkovRatio) % oneDay
	// 返回塔科夫时间，使用1970年1月1日UTC时间作为基准
	tarkovTime := time.UnixMilli(tarkovTimestamp).UTC()

	return tarkovTime
}

// formatTime 格式化时间为 HH:mm:ss
func formatTime(t time.Time) string {
	return t.Format("15:04:05")
}

func QueryClock() string {
	// 获取当前时间
	now := time.Now()

	// 转换为塔科夫白天时间
	leftTime := realTimeToTarkovTime(now, true)
	// 转换为塔科夫夜晚时间
	rightTime := realTimeToTarkovTime(now, false)

	// 输出格式化后的塔科夫时间
	leftTimeString := formatTime(leftTime)
	rightTimeString := formatTime(rightTime)

	// 显示塔科夫时间
	return fmt.Sprintf("Tarkov Time: %s / %s\n", leftTimeString, rightTimeString)
}
