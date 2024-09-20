package query_test

import (
	"eftbot/query"
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	s, _ := query.QueryByCn("pack of milk")

	fmt.Print(s)
}

func TestQueryDetail(t *testing.T) {
	s, _ := query.QueryWeather()
	t.Log(s)
	s = query.QueryClock()
	t.Log(s)

}
