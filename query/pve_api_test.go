package query_test

import (
	"eftbot/query"
	"testing"
)

func TestPVEAPI(t *testing.T) {
	bsgID := "587e02ff24597743df3deaeb"
	query.QueryPVEFlea(bsgID)

}
